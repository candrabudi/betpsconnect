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

type CoordinationCity interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultAllCoordinatorCity, error)
	Store(ctx context.Context, newData dto.PayloadStoreCoordinatorCity) error
	Update(ctx context.Context, ID int, updatedData dto.PayloadStoreCoordinatorCity) error
}

type coordinationcity struct {
	MongoConn *mongo.Client
}

func NewCoordinationCityRepository(mongoConn *mongo.Client) CoordinationCity {
	return &coordinationcity{
		MongoConn: mongoConn,
	}
}

func (cc *coordinationcity) GetAll(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultAllCoordinatorCity, error) {

	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["korkab_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":            1,
			"id":             1,
			"korkab_name":    1,
			"korkab_nik":     1,
			"korkab_phone":   1,
			"korkab_age":     1,
			"korkab_address": 1,
			"korkab_city":    1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultAllCoordinatorCity{}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindCoordinatorCity
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultAllCoordinatorCity{}, err
	}

	totalResults, err := cc.GetTotalFilteredCoordinationCount(ctx, filter)
	if err != nil {
		return dto.ResultAllCoordinatorCity{}, err
	}

	result := dto.ResultAllCoordinatorCity{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindCoordinatorCity{}
	}

	return result, nil
}

func (cc *coordinationcity) GetTotalFilteredCoordinationCount(ctx context.Context, filter dto.ResidentFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"korkab_name": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
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

func (cc *coordinationcity) Store(ctx context.Context, newData dto.PayloadStoreCoordinatorCity) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := cc.MongoConn.Database(dbName).Collection("coordination_city")
	trueFilter := bson.M{"kab_nik": newData.KorkabNik}

	var existingResident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&existingResident)
	if err == nil {
		return errors.New("NIK already exists")
	}

	lastData, err := cc.getLastData(ctx, collection)
	if err != nil {
		return err
	}
	newID := lastData.ID + 1

	TrueResident := model.CoordinationCity{
		ID:            newID,
		KorkabName:    newData.KorkabName,
		KorkabNik:     newData.KorkabNik,
		KorkabPhone:   newData.KorkabPhone,
		KorkabAge:     newData.KorkabAge,
		KorkabAddress: newData.KorkabAddress,
		KorkabCity:    newData.KorkabCity,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (cc *coordinationcity) Update(ctx context.Context, ID int, updatedData dto.PayloadStoreCoordinatorCity) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	update := bson.M{
		"$set": bson.M{
			"korkab_name":    updatedData.KorkabName,
			"korkab_nik":     updatedData.KorkabNik,
			"korkab_phone":   updatedData.KorkabPhone,
			"korkab_age":     updatedData.KorkabAge,
			"korkab_address": updatedData.KorkabAddress,
			"korkab_city":    updatedData.KorkabCity,
			"updated_at":     time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cc *coordinationcity) getLastData(ctx context.Context, collection *mongo.Collection) (model.Resident, error) {
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
