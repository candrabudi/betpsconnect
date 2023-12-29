package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SubDistrict interface {
	GetByDistrict(ctx context.Context, filter dto.GetByDistrict) ([]string, error)
	FindOne(ctx context.Context, filter dto.GetByDistrict) (dto.GetByDistrict, error)
	Store(ctx context.Context, data model.SubDistrict) error
	FindLastOne(ctx context.Context) (dto.GetByDistrict, error)
}

type subdistrict struct {
	MongoConn *mongo.Client
}

func NewSubDistrictRepository(mongoConn *mongo.Client) SubDistrict {
	return &subdistrict{
		MongoConn: mongoConn,
	}
}

func (s *subdistrict) GetByDistrict(ctx context.Context, filter dto.GetByDistrict) ([]string, error) {
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

	var kelurahan []string
	for cursor.Next(ctx) {
		var dresident model.Resident
		if err := cursor.Decode(&dresident); err != nil {
			return nil, err
		}
		kelurahan = append(kelurahan, dresident.NamaKelurahan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Menggunakan sort.Strings untuk mengurutkan kelurahan dari awal ke akhir
	sort.Strings(kelurahan)

	return kelurahan, nil
}

func (s *subdistrict) FindOne(ctx context.Context, filter dto.GetByDistrict) (dto.GetByDistrict, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "subdistricts"

	collection := s.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	if filter.NamaKecamatan != "" {
		bsonFilter["nama_kabupaten"] = filter.NamaKabupaten
	}
	if filter.NamaKecamatan != "" {
		bsonFilter["nama_kecamatan"] = filter.NamaKecamatan
	}
	if filter.NamaKelurahan != "" {
		bsonFilter["nama_kelurahan"] = filter.NamaKelurahan
	}
	var dresident model.Resident
	err := collection.FindOne(ctx, bsonFilter).Decode(&dresident)
	if err != nil {
		return dto.GetByDistrict{}, err
	}

	districtDTO := dto.GetByDistrict{
		ID:            dresident.ID,
		NamaKecamatan: dresident.NamaKecamatan,
		NamaKelurahan: dresident.NamaKelurahan,
	}

	return districtDTO, nil
}

func (s *subdistrict) Store(ctx context.Context, data model.SubDistrict) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "subdistricts"

	collection := s.MongoConn.Database(dbName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *subdistrict) FindLastOne(ctx context.Context) (dto.GetByDistrict, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "subdistricts"

	collection := s.MongoConn.Database(dbName).Collection(collectionName)
	opts := options.FindOne()
	opts.SetSort(bson.D{{"_id", -1}}) // Mengurutkan berdasarkan field "_id" secara descending (dari besar ke kecil)
	var dresident model.Resident
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&dresident)
	if err != nil {
		return dto.GetByDistrict{}, err
	}

	districtDTO := dto.GetByDistrict{
		ID:            dresident.ID,
		NamaKecamatan: dresident.NamaKecamatan,
		NamaKelurahan: dresident.NamaKelurahan,
		// Tambahkan field lain yang perlu di-decode dari dokumen
	}

	return districtDTO, nil
}
