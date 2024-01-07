package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrueResident interface {
	Store(ctx context.Context, newData dto.TrueResidentPayload) error
	GetAll(ctx context.Context, limit, offset int64, filter dto.TrueResidentFilter) (dto.ResultAllTrueResident, error)
	Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateTrueResident) error
	GetTpsOnValidResident(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error)
	Delete(ctx context.Context, ID int) error
	Export(ctx context.Context, filter dto.TrueResidentFilter) ([]byte, error)
}

type trueresident struct {
	MongoConn *mongo.Client
}

func NewTrueResidentRepository(mongoConn *mongo.Client) TrueResident {
	return &trueresident{
		MongoConn: mongoConn,
	}
}

func (tr *trueresident) GetAll(ctx context.Context, limit, offset int64, filter dto.TrueResidentFilter) (dto.ResultAllTrueResident, error) {

	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		matchStage["subdistrict"] = filter.NamaKelurahan
	}

	if filter.Jaringan != "" {
		matchStage["network"] = filter.Jaringan
	}

	if filter.TPS != "" {
		matchStage["tps"] = filter.TPS
	}

	if filter.IsManual != "" {
		isManualInt, err := strconv.Atoi(filter.IsManual)
		if err == nil {
			matchStage["is_manual"] = isManualInt
		}
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["full_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":          1,
			"id":           1,
			"full_name":    1,
			"gender":       1,
			"city":         1,
			"district":     1,
			"subdistrict":  1,
			"address":      1,
			"birth_date":   1,
			"age":          1,
			"tps":          1,
			"nik":          1,
			"no_handphone": 1,
			"network":      1,
			"is_manual":    1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultAllTrueResident{}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindTrueAllResident
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultAllTrueResident{}, err
	}

	totalResults, err := tr.GetTotalFilteredResidentCount(ctx, filter)
	if err != nil {
		return dto.ResultAllTrueResident{}, err
	}

	result := dto.ResultAllTrueResident{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindTrueAllResident{}
	}

	return result, nil
}

func (tr *trueresident) Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateTrueResident) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	update := bson.M{
		"$set": bson.M{
			"full_name":    updatedData.FullName,
			"nik":          updatedData.Nik,
			"no_handphone": updatedData.NoHandphone,
			"age":          updatedData.Age,
			"address":      updatedData.Address,
			"subdistrict":  updatedData.Subdistrict,
			"district":     updatedData.District,
			"city":         updatedData.City,
			"tps":          removeLeadingZeros(updatedData.TPS),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (tr *trueresident) GetTotalFilteredResidentCount(ctx context.Context, filter dto.TrueResidentFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.NamaKabupaten != "" {
		filterOptions["city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		filterOptions["district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		filterOptions["subdistrict"] = filter.NamaKelurahan
	}

	if filter.TPS != "" {
		filterOptions["tps"] = filter.TPS
	}

	if filter.IsManual != "" {
		isManualInt, err := strconv.Atoi(filter.IsManual)
		if err != nil {
			fmt.Println("Failed to convert IsManual to int:", err)
		} else {
			filterOptions["is_manual"] = isManualInt
		}
	}

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

func (tr *trueresident) Store(ctx context.Context, newData dto.TrueResidentPayload) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := tr.MongoConn.Database(dbName).Collection("true_residents")
	trueFilter := bson.M{"nik": newData.Nik}

	var existingResident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&existingResident)
	if err == nil {
		return errors.New("NIK already exists")
	}

	person, err := tr.getLastPerson(ctx, collection)
	if err != nil {
		return err
	}
	newUserID := person.ID + 1
	TrueResident := model.TrueResident{
		ID:          newUserID,
		FullName:    newData.FullName,
		Nik:         newData.Nik,
		NoHandphone: newData.NoHandphone,
		Age:         newData.Age,
		Gender:      newData.Gender,
		City:        newData.City,
		District:    newData.District,
		SubDistrict: newData.Subdistrict,
		Address:     newData.Address,
		Jaringan:    newData.Jaringan,
		Tps:         removeLeadingZeros(newData.TPS),
		IsManual:    1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func removeLeadingZeros(s string) string {
	for strings.HasPrefix(s, "0") && len(s) > 1 {
		s = s[1:]
	}
	return s
}

func (tr *trueresident) getLastPerson(ctx context.Context, collection *mongo.Collection) (model.Resident, error) {
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

func (tr *trueresident) GetTpsOnValidResident(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)

	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"city":        filter.NamaKabupaten,
				"district":    filter.NamaKecamatan,
				"subdistrict": filter.NamaKelurahan,
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
				nonEmptyResults = append(nonEmptyResults, tps)
			}
		}

		if len(nonEmptyResults) == 0 {
			return []string{}, nil
		}

		return nonEmptyResults, nil
	}

	if err := cursor.Err(); err != nil {
		return []string{}, err
	}

	return results, nil
}

func (tr *trueresident) Delete(ctx context.Context, ID int) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	var resident model.TrueResident
	err := collection.FindOne(ctx, filter).Decode(&resident)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return errors.New("data not found")
		}
		return err
	}

	NIK := resident.Nik

	collectionResident := "residents"

	collectionUpdateResident := tr.MongoConn.Database(dbName).Collection(collectionResident)
	filterUpdate := bson.M{"nik": NIK}

	update := bson.M{
		"$set": bson.M{
			"is_verification": 0,
			"is_delete":       0,
			"updated_at":      time.Now(),
		},
	}

	_, err = collectionUpdateResident.UpdateOne(ctx, filterUpdate, update)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (tr *trueresident) Export(ctx context.Context, filter dto.TrueResidentFilter) ([]byte, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "true_residents"

	collection := tr.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := make([]bson.M, 0)

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["district"] = filter.NamaKecamatan
	}

	if filter.NamaKelurahan != "" {
		matchStage["subdistrict"] = filter.NamaKelurahan
	}

	if filter.Jaringan != "" {
		matchStage["network"] = filter.Jaringan
	}

	if filter.TPS != "" {
		matchStage["tps"] = filter.TPS
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["full_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.ExportTrueAllResident
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return nil, err
	}

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	xlsx.NewSheet(sheetName)

	if err := tr.setHeaderAndStyle(xlsx, sheetName); err != nil {
		return nil, err
	}

	tr.setDataRows(xlsx, sheetName, dataAllResident)

	fileBytes, err := xlsx.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return fileBytes.Bytes(), nil
}

func (tr *trueresident) setHeaderAndStyle(xlsx *excelize.File, sheetName string) error {
	headers := []string{"NAMA", "NIK", "NOMOR HANDPHONE", "JK", "UMUR", "ALAMAT", "KABUPATEN", "KECAMATAN", "KELURAHAN", "TPS", "JARINGAN", "TANGGAL BUAT", "TANGGAL UPDATE"}
	widths := []float64{30, 30, 25, 10, 10, 35, 20, 20, 20, 15, 15, 20, 20}

	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	if err != nil {
		return err
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%c%d", 'A'+i, 1)
		xlsx.SetCellValue(sheetName, cell, header)
		xlsx.SetCellStyle(sheetName, cell, cell, style)
		xlsx.SetColWidth(sheetName, cell[:1], cell[:1], widths[i])
	}
	return nil
}

func (tr *trueresident) setDataRows(xlsx *excelize.File, sheetName string, dataTrueResident []dto.ExportTrueAllResident) {
	style, _ := xlsx.NewStyle(`{"alignment":{"horizontal":"center"}}`)

	for i, data := range dataTrueResident {
		rowData := []interface{}{
			data.Nama,
			data.Nik,
			data.NoHandphone,
			data.JenisKelamin,
			data.Usia,
			data.Address,
			data.NamaKabupaten,
			data.NamaKecamatan,
			data.NamaKelurahan,
			data.Tps,
			data.Jaringan,
			data.CreatedAt.Format("2006-01-02"),
			data.UpdatedAt.Format("2006-01-02"),
		}
		for j, value := range rowData {
			cell := fmt.Sprintf("%c%d", 'A'+j, i+2)
			xlsx.SetCellValue(sheetName, cell, value)
			xlsx.SetCellStyle(sheetName, cell, cell, style)
		}
	}
}
