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

type District interface {
	GetByCity(ctx context.Context, filter dto.GetByCity) ([]string, error)
	FindOne(ctx context.Context, filter dto.GetByCity) (dto.GetByCity, error)
	Store(ctx context.Context, data model.District) error
	FindLastOne(ctx context.Context) (dto.GetByCity, error)
}

type district struct {
	MongoConn *mongo.Client
}

func NewDistrictRepository(mongoConn *mongo.Client) District {
	return &district{
		MongoConn: mongoConn,
	}
}
func (d *district) GetByCity(ctx context.Context, filter dto.GetByCity) ([]string, error) {
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

	var kecamatan []string
	for cursor.Next(ctx) {
		var dresident model.Resident
		if err := cursor.Decode(&dresident); err != nil {
			return nil, err
		}
		kecamatan = append(kecamatan, dresident.NamaKecamatan)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Menggunakan sort.Strings untuk mengurutkan kecamatan dari awal ke akhir
	sort.Strings(kecamatan)

	// Jika data kosong, kembalikan slice kosong
	if len(kecamatan) == 0 {
		return []string{}, nil
	}

	return kecamatan, nil
}

func (s *district) FindOne(ctx context.Context, filter dto.GetByCity) (dto.GetByCity, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "districts"

	collection := s.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	if filter.NamaKecamatan != "" {
		bsonFilter["nama_kabupaten"] = filter.NamaKabupaten
	}
	if filter.NamaKecamatan != "" {
		bsonFilter["nama_kecamatan"] = filter.NamaKecamatan
	}
	var ddistrict model.District
	err := collection.FindOne(ctx, bsonFilter).Decode(&ddistrict)
	if err != nil {
		return dto.GetByCity{}, err
	}

	districtDTO := dto.GetByCity{
		ID:            ddistrict.ID,
		NamaKecamatan: ddistrict.NamaKecamatan,
	}

	return districtDTO, nil
}

func (s *district) Store(ctx context.Context, data model.District) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "districts"

	collection := s.MongoConn.Database(dbName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *district) FindLastOne(ctx context.Context) (dto.GetByCity, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "districts"

	collection := s.MongoConn.Database(dbName).Collection(collectionName)
	opts := options.FindOne()
	opts.SetSort(bson.D{{"_id", -1}}) // Mengurutkan berdasarkan field "_id" secara descending (dari besar ke kecil)
	var ddistrict model.District
	err := collection.FindOne(ctx, bson.M{}, opts).Decode(&ddistrict)
	if err != nil {
		return dto.GetByCity{}, err
	}

	districtDTO := dto.GetByCity{
		ID:            ddistrict.ID,
		NamaKecamatan: ddistrict.NamaKecamatan,
	}

	return districtDTO, nil
}
