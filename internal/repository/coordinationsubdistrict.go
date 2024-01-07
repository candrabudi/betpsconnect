package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"errors"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CoordinationSubdistrict interface {
	Store(ctx context.Context, newData dto.PayloadStoreCoordinatorSubdistrict) error
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationSubdistrictFilter) (dto.ResultAllCoordinatorSubdistrict, error)
	Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateCoordinatorSubdistrict) error
	Delete(ctx context.Context, ID int) error
}

type coordinationsubdistrict struct {
	MongoConn *mongo.Client
}

func NewCoordinationSubdistrictRepository(mongoConn *mongo.Client) CoordinationSubdistrict {
	return &coordinationsubdistrict{
		MongoConn: mongoConn,
	}
}

func (cs *coordinationsubdistrict) GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationSubdistrictFilter) (dto.ResultAllCoordinatorSubdistrict, error) {

	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_subdistrict"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["kordes_city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["kordes_district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		matchStage["kordes_kelurahan"] = filter.NamaKelurahan
	}

	if filter.Jaringan != "" {
		matchStage["kordes_network"] = filter.Jaringan
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["kordes_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":                1,
			"id":                 1,
			"kordes_name":        1,
			"kordes_nik":         1,
			"kordes_phone":       1,
			"kordes_age":         1,
			"kordes_gender":      1,
			"kordes_address":     1,
			"kordes_city":        1,
			"kordes_district":    1,
			"kordes_subdistrict": 1,
			"kordes_network":     1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultAllCoordinatorSubdistrict{}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindCoordinatorSubdistrict
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultAllCoordinatorSubdistrict{}, err
	}

	totalResults, err := cs.GetTotalFilteredCoordinationCount(ctx, filter)
	if err != nil {
		return dto.ResultAllCoordinatorSubdistrict{}, err
	}

	result := dto.ResultAllCoordinatorSubdistrict{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindCoordinatorSubdistrict{}
	}

	return result, nil
}

func (cs *coordinationsubdistrict) GetTotalFilteredCoordinationCount(ctx context.Context, filter dto.CoordinationSubdistrictFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_subdistrict"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.NamaKabupaten != "" {
		filterOptions["kordes_city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		filterOptions["kordes_district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		filterOptions["kordes_subdistrict"] = filter.NamaKelurahan
	}

	if filter.Jaringan != "" {
		filterOptions["kordes_network"] = filter.Jaringan
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"kordes_name": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
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

func (cs *coordinationsubdistrict) Store(ctx context.Context, newData dto.PayloadStoreCoordinatorSubdistrict) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := cs.MongoConn.Database(dbName).Collection("coordination_subdistrict")
	trueFilter := bson.M{"kordes_nik": newData.KordesNik}

	var existingResident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&existingResident)
	if err == nil {
		return errors.New("NIK already exists")
	}

	lastData, err := cs.getLastData(ctx, collection)
	if err != nil {
		return err
	}
	newID := lastData.ID + 1

	TrueResident := model.CoordinationSubdistrict{
		ID:                newID,
		KordesName:        newData.KordesName,
		KordesNik:         newData.KordesNik,
		KordesPhone:       newData.KordesPhone,
		KordesAge:         newData.KordesAge,
		KordesGender:      newData.KordesGender,
		KordesAddress:     newData.KordesAddress,
		KordesCity:        newData.KordesCity,
		KordesDistrict:    newData.KordesDistrict,
		KordesSubdistrict: newData.KordesSubdistrict,
		KordesNetwork:     newData.KordesNetwork,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (cs *coordinationsubdistrict) Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateCoordinatorSubdistrict) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_subdistrict"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	update := bson.M{
		"$set": bson.M{
			"kordes_name":        updatedData.KordesName,
			"kordes_nik":         updatedData.KordesNik,
			"kordes_phone":       updatedData.KordesPhone,
			"kordes_age":         updatedData.KordesAge,
			"kordes_gender":      updatedData.KordesGender,
			"kordes_address":     updatedData.KordesAddress,
			"kordes_city":        updatedData.KordesCity,
			"kordes_district":    updatedData.KordesDistrict,
			"kordes_subdistrict": updatedData.KordesSubdistrict,
			"kordes_network":     updatedData.KordesNetwork,
			"updated_at":         time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cs *coordinationsubdistrict) Delete(ctx context.Context, ID int) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_subdistrict"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (cs *coordinationsubdistrict) getLastData(ctx context.Context, collection *mongo.Collection) (model.CoordinationSubdistrict, error) {
	opts := options.FindOne().SetSort(bson.D{{"$natural", -1}})

	var person model.CoordinationSubdistrict
	if err := collection.FindOne(ctx, bson.D{}, opts).Decode(&person); err != nil {
		if err == mongo.ErrNoDocuments {
			return model.CoordinationSubdistrict{}, nil
		}
		return model.CoordinationSubdistrict{}, err
	}

	return person, nil
}
