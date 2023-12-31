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

	"github.com/360EntSecGroup-Skylar/excelize"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CoordinationDistrict interface {
	Store(ctx context.Context, newData dto.PayloadStoreCoordinatorDistrict) error
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationDistrictFilter) (dto.ResultAllCoordinatorDistrict, error)
	Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateCoordinatorDistrict) error
	Delete(ctx context.Context, ID int) error
	Export(ctx context.Context, filter dto.CoordinationDistrictFilter) ([]byte, error)
}

type coordinationdistrict struct {
	MongoConn *mongo.Client
}

func NewCoordinationDistrictRepository(mongoConn *mongo.Client) CoordinationDistrict {
	return &coordinationdistrict{
		MongoConn: mongoConn,
	}
}

func (cd *coordinationdistrict) GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationDistrictFilter) (dto.ResultAllCoordinatorDistrict, error) {

	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_district"

	collection := cd.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["korcam_city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["korcam_district"] = filter.NamaKecamatan
	}

	if filter.Jaringan != "" {
		matchStage["korcam_network"] = filter.Jaringan
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["korcam_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	projectStage := bson.M{
		"$project": bson.M{
			"_id":             1,
			"id":              1,
			"korcam_name":     1,
			"korcam_nik":      1,
			"korcam_phone":    1,
			"korcam_age":      1,
			"korcam_gender":   1,
			"korcam_address":  1,
			"korcam_city":     1,
			"korcam_district": 1,
			"korcam_network":  1,
		},
	}

	pipeline = append(pipeline, projectStage)

	pipeline = append(pipeline, bson.M{"$skip": offset})
	pipeline = append(pipeline, bson.M{"$limit": limit})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return dto.ResultAllCoordinatorDistrict{}, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.FindCoordinatorDistrict
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return dto.ResultAllCoordinatorDistrict{}, err
	}

	totalResults, err := cd.GetTotalFilteredCoordinationCount(ctx, filter)
	if err != nil {
		return dto.ResultAllCoordinatorDistrict{}, err
	}

	result := dto.ResultAllCoordinatorDistrict{
		Items: dataAllResident,
		Metadata: dto.MetaData{
			TotalResults: int(totalResults),
			Limit:        int(limit),
			Offset:       int(offset),
			Count:        len(dataAllResident),
		},
	}

	if totalResults == 0 {
		result.Items = []dto.FindCoordinatorDistrict{}
	}

	return result, nil
}

func (cd *coordinationdistrict) GetTotalFilteredCoordinationCount(ctx context.Context, filter dto.CoordinationDistrictFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_district"

	collection := cd.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.NamaKabupaten != "" {
		filterOptions["korcam_city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		filterOptions["korcam_district"] = filter.NamaKecamatan
	}

	if filter.Jaringan != "" {
		filterOptions["korcam_network"] = filter.Jaringan
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"korcam_name": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
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

func (cd *coordinationdistrict) Store(ctx context.Context, newData dto.PayloadStoreCoordinatorDistrict) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := cd.MongoConn.Database(dbName).Collection("coordination_district")
	trueFilter := bson.M{"korcam_nik": newData.KorcamNik}

	var existingResident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&existingResident)
	if err == nil {
		return errors.New("NIK already exists")
	}

	lastData, err := cd.getLastData(ctx, collection)
	if err != nil {
		return err
	}
	newID := lastData.ID + 1

	TrueResident := model.CoordinationDistrict{
		ID:             newID,
		KorcamName:     newData.KorcamName,
		KorcamNik:      newData.KorcamNik,
		KorcamPhone:    newData.KorcamPhone,
		KorcamAge:      newData.KorcamAge,
		KorcamGender:   newData.KorcamGender,
		KorcamAddress:  newData.KorcamAddress,
		KorcamCity:     newData.KorcamCity,
		KorcamDistrict: newData.KorcamDistrict,
		KorcamNetwork:  newData.KorcamNetwork,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (cd *coordinationdistrict) Update(ctx context.Context, ID int, updatedData dto.PayloadUpdateCoordinatorDistrict) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_district"

	collection := cd.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	update := bson.M{
		"$set": bson.M{
			"korcam_name":     updatedData.KorcamName,
			"korcam_nik":      updatedData.KorcamNik,
			"korcam_phone":    updatedData.KorcamPhone,
			"korcam_age":      updatedData.KorcamAge,
			"korcam_gender":   updatedData.KorcamGender,
			"korcam_address":  updatedData.KorcamAddress,
			"korcam_city":     updatedData.KorcamCity,
			"korcam_district": updatedData.KorcamDistrict,
			"korcam_network":  updatedData.KorcamNetwork,
			"updated_at":      time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cs *coordinationdistrict) Delete(ctx context.Context, ID int) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_district"

	collection := cs.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (cd *coordinationdistrict) Export(ctx context.Context, filter dto.CoordinationDistrictFilter) ([]byte, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_district"

	collection := cd.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := make([]bson.M, 0)

	matchStage := bson.M{}

	if filter.NamaKabupaten != "" {
		matchStage["korkab_city"] = filter.NamaKabupaten
	}

	if filter.NamaKecamatan != "" {
		matchStage["korkab_district"] = filter.NamaKecamatan
	}

	if filter.Jaringan != "" {
		matchStage["korkab_network"] = filter.Jaringan
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["korkab_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}

	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var dataAllResident []dto.ExportCoordinatorDistrict
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return nil, err
	}

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	xlsx.NewSheet(sheetName)

	if err := cd.setHeaderAndStyle(xlsx, sheetName); err != nil {
		return nil, err
	}

	cd.setDataRows(xlsx, sheetName, dataAllResident)

	fileBytes, err := xlsx.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return fileBytes.Bytes(), nil
}

func (cd *coordinationdistrict) setHeaderAndStyle(xlsx *excelize.File, sheetName string) error {
	headers := []string{"NAMA", "NIK", "NOMOR HANDPHONE", "JK", "UMUR", "ALAMAT", "KABUPATEN", "KECAMATAN", "JARINGAN", "TANGGAL BUAT", "TANGGAL UPDATE"}
	widths := []float64{30, 30, 25, 10, 10, 35, 20, 20, 15, 20, 20}

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

func (cd *coordinationdistrict) setDataRows(xlsx *excelize.File, sheetName string, dataAllResident []dto.ExportCoordinatorDistrict) {
	style, _ := xlsx.NewStyle(`{"alignment":{"horizontal":"center"}}`)

	for i, data := range dataAllResident {
		rowData := []interface{}{data.Nama, data.Nik, data.NoHandphone, data.Gender, data.Usia, data.Alamat, data.NamaKabupaten, data.NamaKecamatan, data.Jaringan, data.CreatedAt.Format("2006-01-02"), data.UpdatedAt.Format("2006-01-02")}
		for j, value := range rowData {
			cell := fmt.Sprintf("%c%d", 'A'+j, i+2)
			xlsx.SetCellValue(sheetName, cell, value)
			xlsx.SetCellStyle(sheetName, cell, cell, style)
		}
	}
}

func (cd *coordinationdistrict) getLastData(ctx context.Context, collection *mongo.Collection) (model.Resident, error) {
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
