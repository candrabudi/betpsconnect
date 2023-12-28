package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type District interface {
	GetByCity(ctx context.Context, filter dto.GetByCity) ([]dto.GetByCity, error)
}

type district struct {
	MongoConn *mongo.Client
}

func NewDistrictRepository(mongoConn *mongo.Client) District {
	return &district{
		MongoConn: mongoConn,
	}
}

func (d *district) GetByCity(ctx context.Context, filter dto.GetByCity) ([]dto.GetByCity, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "districts"

	collection := d.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	if filter.NamaKabupaten != "" {
		bsonFilter["nama_kabupaten"] = filter.NamaKabupaten
	}
	if filter.NamaKecamatan != "" {
		bsonFilter["nama_kecamatan"] = filter.NamaKecamatan
	}

	// Lakukan query ke MongoDB menggunakan filter yang telah dibuat sebelumnya
	cursor, err := collection.Find(ctx, bsonFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dataDistrictByCity []dto.GetByCity
	for cursor.Next(ctx) {
		var dresident model.Resident
		if err := cursor.Decode(&dresident); err != nil {
			return nil, err
		}
		districtDTO := dto.GetByCity{
			ID:            dresident.ID,
			NamaKabupaten: dresident.NamaKabupaten,
			NamaKecamatan: dresident.NamaKecamatan,
		}

		dataDistrictByCity = append(dataDistrictByCity, districtDTO)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return dataDistrictByCity, nil
}
