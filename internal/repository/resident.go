package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resident interface {
	GetResidentTps(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error)
	GetAllResidents(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error)
	GetTpsBySubDistrict(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error)
	Store(ctx context.Context, data model.Resident) error
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

func (r *resident) GetResidentTps(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error) {
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

	// Menggunakan EstimatedDocumentCount untuk menghitung total data
	totalResults, err := collection.EstimatedDocumentCount(ctx)
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
		age := dresident.Usia
		if age == 0 {
			dob, err := time.Parse("2006-01-02", dresident.TanggalLahir)
			if err != nil {
				fmt.Println(err)
			}
			now := time.Now()
			age = now.Year() - dob.Year()
			if now.YearDay() < dob.YearDay() {
				age--
			}
		}

		serverDTO := dto.FindAllResident{
			ID:             dresident.ID,
			Nama:           dresident.Nama,
			Alamat:         dresident.Alamat,
			JenisKelamin:   dresident.JenisKelamin,
			Kawin:          dresident.Kawin,
			NamaKabupaten:  dresident.NamaKabupaten,
			NamaKecamatan:  dresident.NamaKecamatan,
			NamaKelurahan:  dresident.NamaKelurahan,
			Nik:            dresident.Nik,
			Nkk:            dresident.Nkk,
			NoKtp:          dresident.NoKtp,
			Rt:             dresident.Rt,
			Rw:             dresident.Rw,
			Status:         dresident.Status,
			StatusTpsLabel: dresident.StatusTpsLabel,
			Tps:            dresident.Tps,
			TanggalLahir:   dresident.TanggalLahir,
			Usia:           age,
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

func (r *resident) GetAllResidents(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultResident, error) {
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

	// Menggunakan EstimatedDocumentCount untuk menghitung total data
	totalResults, err := collection.EstimatedDocumentCount(ctx)
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
		age := dresident.Usia
		if age == 0 {
			dob, err := time.Parse("2006-01-02", dresident.TanggalLahir)
			if err != nil {
				fmt.Println(err)
			}
			now := time.Now()
			age = now.Year() - dob.Year()
			if now.YearDay() < dob.YearDay() {
				age--
			}
		}

		serverDTO := dto.FindAllResident{
			ID:             dresident.ID,
			Nama:           dresident.Nama,
			Alamat:         dresident.Alamat,
			JenisKelamin:   dresident.JenisKelamin,
			Kawin:          dresident.Kawin,
			NamaKabupaten:  dresident.NamaKabupaten,
			NamaKecamatan:  dresident.NamaKecamatan,
			NamaKelurahan:  dresident.NamaKelurahan,
			Nik:            dresident.Nik,
			Nkk:            dresident.Nkk,
			NoKtp:          dresident.NoKtp,
			Rt:             dresident.Rt,
			Rw:             dresident.Rw,
			Status:         dresident.Status,
			StatusTpsLabel: dresident.StatusTpsLabel,
			Tps:            dresident.Tps,
			TanggalLahir:   dresident.TanggalLahir,
			Usia:           age,
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

func (r *resident) Store(ctx context.Context, data model.Resident) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)
	person, err := r.getLastPerson(ctx, collection) // Mengambil ID terakhir
	if err != nil {
		return err // Mengembalikan error jika ada kesalahan dalam mengambil ID terakhir
	}
	newUserID := person.ID + 1 // Menambahkan 1 ke ID terakhir untuk membuat ID baru
	data.ID = newUserID        // Menetapkan ID baru pada data yang akan disimpan

	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err // Mengembalikan error jika ada kesalahan saat menyimpan data
	}
	return nil // Mengembalikan nilai nil jika penyimpanan berhasil
}

func (r *resident) GetTpsBySubDistrict(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"nama_kabupaten": filter.NamaKabupaten,
				"nama_kecamatan": filter.NamaKecamatan,
				"nama_kelurahan": filter.NamaKelurahan,
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$tps",
			},
		},
		bson.M{
			"$sort": bson.M{"_id": 1}, // Urutkan tps dari terkecil ke terbesar
		},
		bson.M{
			"$group": bson.M{
				"_id": nil,
				"tps": bson.M{"$push": "$_id"},
			},
		},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []string
	if cursor.Next(ctx) {
		var result struct {
			TPS []string `bson:"tps"`
		}
		if err := cursor.Decode(&result); err != nil {
			return nil, err
		}
		results = result.TPS
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return results, nil
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
					"nama_kabupaten": "$nama_kabupaten",
					"nama_kecamatan": "$nama_kecamatan",
					"nama_kelurahan": "$nama_kelurahan",
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
			NamaKelurahan: result["_id"].(bson.M)["nama_kelurahan"].(string),
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

func (r *resident) getLastPerson(ctx context.Context, collection *mongo.Collection) (model.Resident, error) {
	opts := options.FindOne().SetSort(bson.D{{"$natural", -1}})

	var person model.Resident
	if err := collection.FindOne(ctx, bson.D{}, opts).Decode(&person); err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Resident{}, nil // Jika tidak ada dokumen, kembalikan data kosong (tidak ada data)
		}
		return model.Resident{}, err // Kembalikan error jika ada kesalahan dalam mengambil data
	}

	return person, nil
}
