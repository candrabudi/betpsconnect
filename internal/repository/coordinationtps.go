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

type CoordinationTps interface {
	Store(ctx context.Context, newData dto.PayloadStoreCoordinatorTps) error
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationTpsFilter) (dto.ResultAllCoordinatorTps, error)
	Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateCoordinatorTps) error
}

type coordinationtps struct {
	MongoConn *mongo.Client
}

func NewCoordinationTpsRepository(mongoConn *mongo.Client) CoordinationTps {
	return &coordinationtps{
		MongoConn: mongoConn,
	}
}

func (cs *coordinationtps) GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationTpsFilter) (dto.ResultAllCoordinatorTps, error) {

	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_tps"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["kortps_city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["kortps_district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		matchStage["kortps_kelurahan"] = filter.NamaKelurahan
	}

	if filter.Jaringan != "" {
		matchStage["kortps_network"] = filter.Jaringan
	}

	if filter.Tps != "" {
		matchStage["kortps_tps"] = filter.Tps
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["kortps_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":                1,
			"id":                 1,
			"kortps_name":        1,
			"kortps_nik":         1,
			"kortps_phone":       1,
			"kortps_age":         1,
			"kortps_gender":      1,
			"kortps_address":     1,
			"kortps_city":        1,
			"kortps_district":    1,
			"kortps_subdistrict": 1,
			"kortps_network":     1,
			"kortps_tps":         1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultAllCoordinatorTps{}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindCoordinatorTps
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultAllCoordinatorTps{}, err
	}

	totalResults, err := cs.GetTotalFilteredCoordinationCount(ctx, filter)
	if err != nil {
		return dto.ResultAllCoordinatorTps{}, err
	}

	result := dto.ResultAllCoordinatorTps{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindCoordinatorTps{}
	}

	return result, nil
}

func (cs *coordinationtps) GetTotalFilteredCoordinationCount(ctx context.Context, filter dto.CoordinationTpsFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_tps"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.NamaKabupaten != "" {
		filterOptions["kortps_city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		filterOptions["kortps_district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		filterOptions["kortps_subdistrict"] = filter.NamaKelurahan
	}

	if filter.Jaringan != "" {
		filterOptions["kortps_jaringan"] = filter.Jaringan
	}

	if filter.Tps != "" {
		filterOptions["kortps_tps"] = filter.Tps
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"kortps_name": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
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

func (cs *coordinationtps) Store(ctx context.Context, newData dto.PayloadStoreCoordinatorTps) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := cs.MongoConn.Database(dbName).Collection("coordination_tps")
	trueFilter := bson.M{"kortps_nik": newData.KorTpsNik}

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

	TrueResident := model.CoordinationTps{
		ID:                newID,
		KorTpsName:        newData.KorTpsName,
		KorTpsNik:         newData.KorTpsNik,
		KorTpsPhone:       newData.KorTpsPhone,
		KorTpsAge:         newData.KorTpsAge,
		KorTpsGender:      newData.KorTpsGender,
		KorTpsAddress:     newData.KorTpsAddress,
		KorTpsCity:        newData.KorTpsCity,
		KorTpsDistrict:    newData.KorTpsDistrict,
		KorTpsSubdistrict: newData.KorTpsSubdistrict,
		KorTpsTps:         removeLeadingZeros(newData.KorTpsTps),
		KorTpsNetwork:     newData.KorTpsNetwork,
		CreatedAt:         time.Now(),
		UpdatedAt:         time.Now(),
	}

	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (cs *coordinationtps) Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateCoordinatorTps) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_tps"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	update := bson.M{
		"$set": bson.M{
			"kortps_name":        updatedData.KorTpsName,
			"kortps_nik":         updatedData.KorTpsNik,
			"kortps_phone":       updatedData.KorTpsPhone,
			"kortps_age":         updatedData.KorTpsAge,
			"kortps_gender":      updatedData.KorTpsGender,
			"kortps_address":     updatedData.KorTpsAddress,
			"kortps_city":        updatedData.KorTpsCity,
			"kortps_district":    updatedData.KorTpsDistrict,
			"kortps_subdistrict": updatedData.KorTpsSubdistrict,
			"kortps_tps":         removeLeadingZeros(updatedData.KorTpsTps),
			"kortps_network":     updatedData.KorTpsNetwork,
			"updated_at":         time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cs *coordinationtps) getLastData(ctx context.Context, collection *mongo.Collection) (model.CoordinationTps, error) {
	opts := options.FindOne().SetSort(bson.D{{"$natural", -1}})

	var person model.CoordinationTps
	if err := collection.FindOne(ctx, bson.D{}, opts).Decode(&person); err != nil {
		if err == mongo.ErrNoDocuments {
			return model.CoordinationTps{}, nil
		}
		return model.CoordinationTps{}, err
	}

	return person, nil
}
