package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/pkg/util"
	"context"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Coordination interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultAllCoordinatorCity, error)
}

type coordination struct {
	MongoConn *mongo.Client
}

func NewCoordinationRepository(mongoConn *mongo.Client) Coordination {
	return &coordination{
		MongoConn: mongoConn,
	}
}

func (cc *coordination) GetAll(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultAllCoordinatorCity, error) {

	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["coordination_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":                  1,
			"id":                   1,
			"coordination_name":    1,
			"coordination_nik":     1,
			"coordination_phone":   1,
			"coordination_age":     1,
			"coordination_address": 1,
			"coordination_city":    1,
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

func (cc *coordination) GetTotalFilteredCoordinationCount(ctx context.Context, filter dto.ResidentFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"full_name": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
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
