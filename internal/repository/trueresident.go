package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrueResident interface {
	Store(ctx context.Context, newData dto.TrueResidentPayload) error
	GetAll(ctx context.Context, limit, offset int64, filter dto.TrueResidentFilter) (dto.ResultAllTrueResident, error)
}

type trueresident struct {
	MongoConn *mongo.Client
}

func NewTrueResidentRepository(mongoConn *mongo.Client) TrueResident {
	return &trueresident{
		MongoConn: mongoConn,
	}
}

func (tr *trueresident) GetAll(ctx context.Context, limit, offset int64, filter dto.TrueResidentFilter) (dto.ResultAllTrueResident, error) {

	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)

	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		matchStage["subdistrict"] = filter.NamaKelurahan
	}

	if filter.TPS != "" {
		matchStage["tps"] = filter.TPS
	}

	if filter.IsManual != "" {
		isManualInt, err := strconv.Atoi(filter.IsManual)
		if err != nil {
			fmt.Println("Failed to convert IsManual to int:", err)
		} else {
			matchStage["is_manual"] = isManualInt
		}
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["$or"] = []bson.M{{"full_name": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":          1,
			"id":           1,
			"full_name":    1,
			"gender":       1,
			"city":         1,
			"district":     1,
			"subdistrict":  1,
			"address":      1,
			"birth_date":   1,
			"age":          1,
			"tps":          1,
			"nik":          1,
			"no_handphone": 1,
			"is_manual":    1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultAllTrueResident{}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindTrueAllResident
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultAllTrueResident{}, err
	}

	totalResults, err := tr.GetTotalFilteredResidentCount(ctx, filter)
	if err != nil {
		return dto.ResultAllTrueResident{}, err
	}

	result := dto.ResultAllTrueResident{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindTrueAllResident{}
	}

	return result, nil
}

func (tr *trueresident) GetTotalFilteredResidentCount(ctx context.Context, filter dto.TrueResidentFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.NamaKabupaten != "" {
		filterOptions["nama_kabupaten"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		filterOptions["nama_kecamatan"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		filterOptions["nama_kelurahan"] = filter.NamaKelurahan
	}

	if filter.TPS != "" {
		filterOptions["tps"] = filter.TPS
	}

	if filter.IsManual != "" {
		isManualInt, err := strconv.Atoi(filter.IsManual)
		if err != nil {
			fmt.Println("Failed to convert IsManual to int:", err)
		} else {
			filterOptions["is_manual"] = isManualInt
		}
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"nama": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}

	countQuery := []bson.M{
		{"$match": filterOptions},
		{"$count": "total"},
	}

	cursor, err := collection.Aggregate(ctx, countQuery)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(ctx)

	var result []bson.M
	if err = cursor.All(ctx, &result); err != nil {
		return 0, err
	}

	if len(result) > 0 {
		total := result[0]["total"].(int32)
		return total, nil
	}

	return 0, nil
}

func (tr *trueresident) Store(ctx context.Context, newData dto.TrueResidentPayload) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := tr.MongoConn.Database(dbName).Collection("true_residents")
	trueFilter := bson.M{"nik": newData.Nik}

	var existingResident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&existingResident)
	if err == nil {
		return errors.New("NIK already exists")
	}

	person, err := tr.getLastPerson(ctx, collection)
	if err != nil {
		return err
	}
	newUserID := person.ID + 1

	TrueResident := model.TrueResident{
		ID:          newUserID,
		FullName:    newData.FullName,
		Nik:         newData.Nik,
		NoHandphone: newData.NoHandphone,
		Age:         newData.Age,
		Gender:      newData.Gender,
		City:        newData.City,
		District:    newData.District,
		SubDistrict: newData.Subdistrict,
		Tps:         newData.TPS,
		IsManual:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (tr *trueresident) getLastPerson(ctx context.Context, collection *mongo.Collection) (model.Resident, error) {
	opts := options.FindOne().SetSort(bson.D{{"$natural", -1}})

	var person model.Resident
	if err := collection.FindOne(ctx, bson.D{}, opts).Decode(&person); err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Resident{}, nil
		}
		return model.Resident{}, err
	}

	return person, nil
}
