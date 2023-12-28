package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type SubDistrict interface {
	GetByDistrict(ctx context.Context, filter dto.GetByDistrict) ([]dto.GetByDistrict, error)
}

type subdistrict struct {
	MongoConn *mongo.Client
}

func NewSubDistrictRepository(mongoConn *mongo.Client) SubDistrict {
	return &subdistrict{
		MongoConn: mongoConn,
	}
}

func (s *subdistrict) GetByDistrict(ctx context.Context, filter dto.GetByDistrict) ([]dto.GetByDistrict, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "subdistricts"

	collection := s.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	if filter.NamaKecamatan != "" {
		bsonFilter["nama_kecamatan"] = filter.NamaKecamatan
	}
	if filter.NamaKelurahan != "" {
		bsonFilter["nama_kelurahan"] = filter.NamaKelurahan
	}
	cursor, err := collection.Find(ctx, bsonFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dataDistrictByCity []dto.GetByDistrict
	for cursor.Next(ctx) {
		var dresident model.Resident
		if err := cursor.Decode(&dresident); err != nil {
			return nil, err
		}
		districtDTO := dto.GetByDistrict{
			ID:            dresident.ID,
			NamaKecamatan: dresident.NamaKecamatan,
			NamaKelurahan: dresident.NamaKelurahan,
		}

		dataDistrictByCity = append(dataDistrictByCity, districtDTO)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return dataDistrictByCity, nil
}
