package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"errors"
	"fmt"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Resident interface {
	GetResidentTps(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultTpsResidents, error)
	ResidentValidate(ctx context.Context, newData dto.PayloadUpdateValidInvalid) ([]int, error)
	DetailResident(ctx context.Context, ResidentID int) (dto.DetailResident, error)
	CheckResidentByNik(ctx context.Context, Nik string) (dto.DetailResident, error)
	GetTpsBySubDistrict(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error)
	Store(ctx context.Context, data model.TrueResident) error
	GetKecamatanByKabupaten(ctx context.Context, kabupatenName string) ([]dto.FindAllResidentGrouped, error)
	GetListValidate(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultValidateResidents, error)
}

type resident struct {
	MongoConn *mongo.Client
}

func NewResidentRepository(mongoConn *mongo.Client) Resident {
	return &resident{
		MongoConn: mongoConn,
	}
}

func (r *resident) GetResidentTps(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultTpsResidents, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)

	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["nama_kabupaten"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["nama_kecamatan"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		matchStage["nama_kelurahan"] = filter.NamaKelurahan
	}

	if filter.TPS != "" {
		matchStage["tps"] = filter.TPS
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["$or"] = []bson.M{{"nama": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}
	matchStage["is_deleted"] = 0

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":              1,
			"id":               1,
			"nama":             1,
			"jenis_kelamin":    1,
			"nama_kecamatan":   1,
			"nama_kelurahan":   1,
			"nama_kabupaten":   1,
			"tanggal_lahir":    1,
			"tps":              1,
			"nik":              1,
			"nkk":              1,
			"status":           1,
			"is_verification":  1,
			"status_tps_label": 1,
			"tempat_lahir":     1,
			"telp":             1,
			"no_ktp":           1,
			"difabel":          1,
			"kawin":            1,
			"rt":               1,
			"rw":               1,
			"alamat":           1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultTpsResidents{
			Items:    []dto.FindTpsResidents{},
			Metadata: dto.MetaData{},
		}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindTpsResidents
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultTpsResidents{
			Items:    []dto.FindTpsResidents{},
			Metadata: dto.MetaData{},
		}, err
	}

	totalResults, err := r.GetTotalFilteredResidentCount(ctx, filter)
	if err != nil {
		return dto.ResultTpsResidents{
			Items:    []dto.FindTpsResidents{},
			Metadata: dto.MetaData{},
		}, err
	}

	result := dto.ResultTpsResidents{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindTpsResidents{}
	}

	return result, nil
}

func (r *resident) GetListValidate(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultValidateResidents, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)

	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["nama_kabupaten"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["nama_kecamatan"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		matchStage["nama_kelurahan"] = filter.NamaKelurahan
	}

	if filter.TPS != "" {
		matchStage["tps"] = filter.TPS
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["$or"] = []bson.M{{"nama": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}
	matchStage["is_deleted"] = 0

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":              1,
			"id":               1,
			"nama":             1,
			"jenis_kelamin":    1,
			"nama_kecamatan":   1,
			"nama_kelurahan":   1,
			"nama_kabupaten":   1,
			"tanggal_lahir":    1,
			"tps":              1,
			"nik":              1,
			"nkk":              1,
			"status":           1,
			"is_verification":  1,
			"status_tps_label": 1,
			"tempat_lahir":     1,
			"telp":             1,
			"no_ktp":           1,
			"difabel":          1,
			"kawin":            1,
			"rt":               1,
			"rw":               1,
			"alamat":           1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultValidateResidents{
			Items:    []dto.FindValidateResidents{},
			Metadata: dto.MetaData{},
		}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindValidateResidents
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultValidateResidents{
			Items:    []dto.FindValidateResidents{},
			Metadata: dto.MetaData{},
		}, err
	}

	totalResults, err := r.GetTotalFilteredResidentCount(ctx, filter)
	if err != nil {
		return dto.ResultValidateResidents{
			Items:    []dto.FindValidateResidents{},
			Metadata: dto.MetaData{},
		}, err
	}

	result := dto.ResultValidateResidents{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindValidateResidents{}
	}

	return result, nil
}
// func (r *resident) GetListValidate(ctx context.Context, limit, offset int64, filter dto.ResidentFilter) (dto.ResultValidateResidents, error) {
// 	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
// 	collectionName := "residents"

// 	collection := r.MongoConn.Database(dbName).Collection(collectionName)

// 	pipeline := []bson.M{}

// 	matchStage := bson.M{}

// 	if filter.NamaKabupaten != "" {
// 		matchStage["nama_kabupaten"] = filter.NamaKabupaten
// 	}

// 	if filter.NamaKecamatan != "" {
// 		matchStage["nama_kecamatan"] = filter.NamaKecamatan
// 	}

// 	if filter.NamaKelurahan != "" {
// 		matchStage["nama_kelurahan"] = filter.NamaKelurahan
// 	}

// 	if filter.TPS != "" {
// 		matchStage["tps"] = filter.TPS
// 	}

// 	if filter.Nama != "" {
// 		regexPattern := regexp.QuoteMeta(filter.Nama)
// 		matchStage["$or"] = []bson.M{{"nama": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
// 	}
// 	matchStage["is_deleted"] = 0

// 	if len(matchStage) > 0 {
// 		pipeline = append(pipeline, bson.M{"$match": matchStage})
// 	}

// 	projectStage := bson.M{
// 		"$project": bson.M{
// 			"_id":              1,
// 			"id":               1,
// 			"nama":             1,
// 			"jenis_kelamin":    1,
// 			"nama_kecamatan":   1,
// 			"tanggal_lahir":    1,
// 			"tps":              1,
// 			"nik":              1,
// 			"nkk":              1,
// 			"status":           1,
// 			"is_verification":  1,
// 			"status_tps_label": 1,
// 			"tempat_lahir":     1,
// 			"telp":             1,
// 			"no_ktp":           1,
// 			"difabel":          1,
// 			"kawin":            1,
// 			"rt":               1,
// 			"rw":               1,
// 			"alamat":           1,
// 		},
// 	}

// 	pipeline = append(pipeline, projectStage)

// 	pipeline = append(pipeline, bson.M{"$skip": offset})
// 	pipeline = append(pipeline, bson.M{"$limit": limit})

// 	cursor, err := collection.Aggregate(ctx, pipeline)
// 	if err != nil {
// 		return dto.ResultValidateResidents{
// 			Items:    []dto.FindValidateResidents{},
// 			Metadata: dto.MetaData{},
// 		}, err
// 	}
// 	defer cursor.Close(ctx)
// 	var dataAllResident []dto.FindValidateResidents
// 	if err = cursor.All(ctx, &dataAllResident); err != nil {
// 		return dto.ResultValidateResidents{
// 			Items:    []dto.FindValidateResidents{},
// 			Metadata: dto.MetaData{},
// 		}, err
// 	}

// 	totalResults, err := r.GetTotalFilteredResidentCount(ctx, filter)
// 	if err != nil {
// 		return dto.ResultValidateResidents{
// 			Items:    []dto.FindValidateResidents{},
// 			Metadata: dto.MetaData{},
// 		}, err
// 	}

// 	result := dto.ResultValidateResidents{
// 		Items: dataAllResident,
// 		Metadata: dto.MetaData{
// 			TotalResults: int(totalResults),
// 			Limit:        int(limit),
// 			Offset:       int(offset),
// 			Count:        len(dataAllResident),
// 		},
// 	}

// 	if totalResults == 0 {
// 		result.Items = []dto.FindValidateResidents{}
// 	}

// 	return result, nil
// }

func (r *resident) GetTotalFilteredResidentCount(ctx context.Context, filter dto.ResidentFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.NamaKabupaten != "" {
		filterOptions["nama_kabupaten"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		filterOptions["nama_kecamatan"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		filterOptions["nama_kelurahan"] = filter.NamaKelurahan
	}

	if filter.TPS != "" {
		filterOptions["tps"] = filter.TPS
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"nama": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}
	filterOptions["is_deleted"] = 0
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

func (r *resident) DetailResident(ctx context.Context, ResidentID int) (dto.DetailResident, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "kaznet")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := make([]bson.M, 0)

	matchStage := bson.M{
		"$match": bson.M{
			"id": ResidentID,
		},
	}
	pipeline = append(pipeline, matchStage)

	projectStage := bson.M{
		"$project": bson.M{
			"_id":              1,
			"id":               1,
			"nama":             1,
			"jenis_kelamin":    1,
			"nama_kecamatan":   1,
			"nama_kabupaten":   1,
			"nama_kelurahan":   1,
			"email":            1,
			"ektp":             1,
			"saringan_id":      1,
			"tanggal_lahir":    1,
			"tps":              1,
			"nik":              1,
			"nkk":              1,
			"status":           1,
			"is_verification":  1,
			"status_tps_label": 1,
			"tempat_lahir":     1,
			"telp":             1,
			"no_ktp":           1,
			"difabel":          1,
			"kawin":            1,
			"rt":               1,
			"rw":               1,
			"alamat":           1,
		},
	}
	pipeline = append(pipeline, projectStage)

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.DetailResident{}, err
	}
	defer cursor.Close(ctx)

	var dataResident dto.DetailResident
	if cursor.Next(ctx) {
		if err := cursor.Decode(&dataResident); err != nil {
			return dto.DetailResident{}, err
		}
	} else {
		return dto.DetailResident{}, errors.New("Resident not found")
	}

	return dataResident, nil
}

func (r *resident) CheckResidentByNik(ctx context.Context, Nik string) (dto.DetailResident, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	bsonFilter["nik"] = Nik
	var dresident model.Resident
	err := collection.FindOne(ctx, bsonFilter).Decode(&dresident)
	if err != nil {
		return dto.DetailResident{}, err
	}
	residentDTO := dto.DetailResident{
		ID:             dresident.ID,
		Nama:           dresident.Nama,
		Alamat:         dresident.Alamat,
		Difabel:        dresident.Difabel,
		Ektp:           dresident.Ektp,
		Email:          dresident.Email,
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
		SaringanID:     dresident.SaringanID,
		Status:         dresident.Status,
		StatusTpsLabel: dresident.StatusTpsLabel,
		TanggalLahir:   dresident.TanggalLahir,
		Usia:           dresident.Usia,
		TempatLahir:    dresident.TempatLahir,
		Telp:           dresident.Telp,
		Tps:            removeLeadingZeros(dresident.Tps),
		IsVerification: dresident.IsVerification,
	}

	return residentDTO, nil
}

func (r *resident) Store(ctx context.Context, data model.TrueResident) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := r.MongoConn.Database(dbName).Collection(collectionName)

	trueFilter := bson.M{"nik": data.Nik}

	var existingResident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&existingResident)
	if err == nil {
		return errors.New("NIK already exists")
	}

	person, err := r.getLastPerson(ctx, collection)
	if err != nil {
		return err
	}
	newUserID := person.ID + 1
	data.ID = newUserID
	_, err = collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (r *resident) ResidentValidate(ctx context.Context, newData dto.PayloadUpdateValidInvalid) ([]int, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := r.MongoConn.Database(dbName).Collection("residents")
	var duplicateData []int

	for _, dataResident := range newData.Items {
		filter := bson.D{
			{"$and", bson.A{
				bson.D{{"id", dataResident.ID}},
				bson.D{{"is_verification", 0}},
			}},
		}

		update := bson.D{}
		if newData.IsTrue == true {
			dresident, err := r.DetailResident(ctx, dataResident.ID)
			if err != nil {
				return duplicateData, err
			}
			trueResidentCollection := r.MongoConn.Database(dbName).Collection("true_residents")
			trueFilter := bson.M{"nik": dresident.Nik}

			var mresident model.TrueResident
			err = trueResidentCollection.FindOne(ctx, trueFilter).Decode(&mresident)

			if err == nil {
				if dresident.IsVerification == 0 {
					filterUpdate := bson.M{"id": dresident.ID}

					update := bson.M{
						"$set": bson.M{
							"is_deleted": 1,
						},
					}
					_, err = collection.UpdateOne(ctx, filterUpdate, update)
					if err != nil {
						return duplicateData, err
					}
				}
				duplicateData = append(duplicateData, dataResident.ID)
				continue
			}

			person, err := r.getLastPerson(ctx, trueResidentCollection)
			if err != nil {
				return duplicateData, err
			}
			newUserID := person.ID + 1
			filterUpdate := bson.M{"id": dresident.ID}

			update := bson.M{
				"$set": bson.M{
					"is_verification": 1,
				},
			}
			result, err := collection.UpdateOne(ctx, filterUpdate, update)
			if err != nil {
				return duplicateData, err
			}
			if result.ModifiedCount > 0 {
				birthDate, _ := time.Parse("2006-01-02", dresident.TanggalLahir)
				age := r.calculateAge(birthDate)
				var tps string
				if dataResident.Tps != "" {
					tps = dataResident.Tps
				} else {
					tps = dresident.Tps
				}
				TrueResident := model.TrueResident{
					ID:          newUserID,
					FullName:    dresident.Nama,
					Address:     dresident.Alamat,
					Nik:         dresident.Nik,
					NoHandphone: dataResident.NoHandphone,
					Age:         age,
					Gender:      dresident.JenisKelamin,
					City:        dresident.NamaKabupaten,
					District:    dresident.NamaKecamatan,
					SubDistrict: dresident.NamaKelurahan,
					Tps:         removeLeadingZeros(tps),
					Jaringan:    dataResident.Jaringan,
					IsManual:    0,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				_, errInsert := trueResidentCollection.InsertOne(ctx, TrueResident)
				if errInsert != nil {
					return duplicateData, errInsert
				}
			} else {
				return duplicateData, errors.New("No data update resident.")
			}
		}

		if len(update) > 0 {
			_, err := collection.UpdateOne(ctx, filter, update)
			if err != nil {
				return duplicateData, err
			}
		}
	}
	if len(duplicateData) == 0 {
		return []int{}, nil
	}
	return duplicateData, nil
}

func (r *resident) calculateAge(TanggalLahir time.Time) int {
	waktuSekarang := time.Now()
	usia := waktuSekarang.Year() - TanggalLahir.Year()
	if waktuSekarang.Month() < TanggalLahir.Month() || (waktuSekarang.Month() == TanggalLahir.Month() && waktuSekarang.Day() < TanggalLahir.Day()) {
		usia--
	}
	return usia
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
			"$sort": bson.M{"_id": 1},
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
		return []string{}, err
	}
	defer cursor.Close(ctx)

	var results []string

	if cursor.Next(ctx) {
		var result struct {
			TPS []string `bson:"tps"`
		}
		if err := cursor.Decode(&result); err != nil {
			return []string{}, err
		}

		var nonEmptyResults []string
		for _, tps := range result.TPS {
			if tps != "" {
				nonEmptyResults = append(nonEmptyResults, removeLeadingZeros(tps))
			}
		}

		if len(nonEmptyResults) == 0 {
			return []string{}, nil
		}
		fmt.Println("KODOK LAST")
		return nonEmptyResults, nil
	}

	if err := cursor.Err(); err != nil {
		return []string{}, err
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
				"count": bson.M{"$sum": 1},
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
			Count:         result["count"].(int32),
		}

		dataKecamatanInKabupaten = append(dataKecamatanInKabupaten, groupedData)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return dataKecamatanInKabupaten, nil
}

func (r *resident) getLastPerson(ctx context.Context, collection *mongo.Collection) (model.Resident, error) {
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

func (r *resident) GetLastValidPerson(ctx context.Context, collection *mongo.Collection) (model.TrueResident, error) {
	opts := options.FindOne().SetSort(bson.D{{"$natural", -1}})

	var person model.TrueResident
	if err := collection.FindOne(ctx, bson.D{}, opts).Decode(&person); err != nil {
		if err == mongo.ErrNoDocuments {
			return model.TrueResident{}, nil
		}
		return model.TrueResident{}, err
	}

	return person, nil
}
