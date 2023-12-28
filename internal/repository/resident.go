package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resident interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error)
	GetKecamatanByKabupaten(ctx context.Context, kabupatenName string) ([]dto.FindAllResidentGrouped, error)
}

type resident struct {
	MongoConn *mongo.Client
}

func NewResidentRepository(mongoConn *mongo.Client) Resident {
	return &resident{
		MongoConn: mongoConn,
	}
}

func (r *resident) GetAll(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)

	bsonFilter := bson.M{}
	if filter.NamaKabupaten != "" {
		bsonFilter["nama_kabupaten"] = filter.NamaKabupaten
	}
	if filter.NamaKecamatan != "" {
		bsonFilter["nama_kecamatan"] = filter.NamaKecamatan
	}
	if filter.NamaKelurahan != "" {
		bsonFilter["nama_kelurahan"] = filter.NamaKelurahan
	}
	if filter.TPS != "" {
		bsonFilter["tps"] = filter.TPS
	}
	if filter.Nama != "" {
		regex := primitive.Regex{Pattern: filter.Nama, Options: "i"} // "i" adalah opsi untuk pencarian case-insensitive
		bsonFilter["nama"] = regex
	}

	totalResults, err := collection.CountDocuments(ctx, bsonFilter)
	if err != nil {
		return dto.ResultResident{}, err
	}

	findOptions := options.Find().SetLimit(limit).SetSkip(offset)
	cursor, err := collection.Find(ctx, bsonFilter, findOptions)
	if err != nil {
		return dto.ResultResident{}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindAllResident
	for cursor.Next(ctx) {
		var dresident model.Resident
		if err := cursor.Decode(&dresident); err != nil {
			return dto.ResultResident{}, err
		}
		serverDTO := dto.FindAllResident{
			ID:   dresident.ID,
			Nama: dresident.Nama,
		}

		dataAllResident = append(dataAllResident, serverDTO)
	}

	if err := cursor.Err(); err != nil {
		return dto.ResultResident{}, err
	}

	count := len(dataAllResident)

	result := dto.ResultResident{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        count,
		},
	}

	return result, nil
}

func (r *resident) GetKecamatanByKabupaten(ctx context.Context, kabupatenName string) ([]dto.FindAllResidentGrouped, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"nama_kabupaten": kabupatenName},
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"nama_kecamatan": "$nama_kecamatan",
					"nama_kabupaten": "$nama_kabupaten",
				},
				"count": bson.M{"$sum": 1}, // Jika ingin menghitung jumlah data
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dataKecamatanInKabupaten []dto.FindAllResidentGrouped
	for cursor.Next(ctx) {
		var result bson.M
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}

		groupedData := dto.FindAllResidentGrouped{
			NamaKecamatan: result["_id"].(bson.M)["nama_kecamatan"].(string),
			NamaKabupaten: result["_id"].(bson.M)["nama_kabupaten"].(string),
			Count:         result["count"].(int32), // Jika menghitung jumlah data
		}

		dataKecamatanInKabupaten = append(dataKecamatanInKabupaten, groupedData)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return dataKecamatanInKabupaten, nil
}

func (r *resident) checkMongoDBConnection(ctx context.Context) error {
	if r.MongoConn == nil {
		return errors.New("MongoDB connection is not established")
	}
	return nil
}

// Helper function to check if the collection exists
func (r *resident) checkCollectionExists(ctx context.Context, dbName, collectionName string) error {
	collections, err := r.MongoConn.Database(dbName).ListCollectionNames(ctx, bson.M{"name": collectionName})
	if err != nil {
		return err
	}

	if len(collections) == 0 {
		return fmt.Errorf("Collection '%s' does not exist in database '%s'", collectionName, dbName)
	}

	return nil
}
