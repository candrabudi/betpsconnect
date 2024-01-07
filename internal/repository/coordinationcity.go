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

type CoordinationCity interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationCityFilter) (dto.ResultAllCoordinatorCity, error)
	Store(ctx context.Context, newData dto.PayloadStoreCoordinatorCity) error
	Update(ctx context.Context, ID int, updatedData dto.PayloadStoreCoordinatorCity) error
	Delete(ctx context.Context, ID int) error
	Export(ctx context.Context, filter dto.CoordinationCityFilter) ([]byte, error)
}

type coordinationcity struct {
	MongoConn *mongo.Client
}

func NewCoordinationCityRepository(mongoConn *mongo.Client) CoordinationCity {
	return &coordinationcity{
		MongoConn: mongoConn,
	}
}

func (cc *coordinationcity) GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationCityFilter) (dto.ResultAllCoordinatorCity, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := []bson.M{}

	matchStage := bson.M{}

	if filter.KorkabCity != "" {
		matchStage["korkab_city"] = filter.KorkabCity
	}

	if filter.Jaringan != "" {
		matchStage["korkab_network"] = filter.Jaringan
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		matchStage["korkab_name"] = primitive.Regex{Pattern: regexPattern, Options: "i"}
	}
	fmt.Println(matchStage)
	if len(matchStage) > 0 {
		pipeline = append(pipeline, bson.M{"$match": matchStage})
	}

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

func (cc *coordinationcity) GetTotalFilteredCoordinationCount(ctx context.Context, filter dto.CoordinationCityFilter) (int32, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)

	filterOptions := bson.M{}

	if filter.KorkabCity != "" {
		regexPattern := regexp.QuoteMeta(filter.KorkabCity)
		filterOptions["$or"] = []bson.M{{"korkab_city": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
	}

	if filter.Jaringan != "" {
		filterOptions["kordes_network"] = filter.Jaringan
	}

	if filter.Nama != "" {
		regexPattern := regexp.QuoteMeta(filter.Nama)
		filterOptions["$or"] = []bson.M{{"korkab_name": primitive.Regex{Pattern: regexPattern, Options: "i"}}}
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

func (cc *coordinationcity) Store(ctx context.Context, newData dto.PayloadStoreCoordinatorCity) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := cc.MongoConn.Database(dbName).Collection("coordination_city")
	trueFilter := bson.M{"kab_nik": newData.KorkabNik}

	var existingResident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&existingResident)
	if err == nil {
		return errors.New("NIK already exists")
	}

	lastData, err := cc.getLastData(ctx, collection)
	if err != nil {
		return err
	}
	newID := lastData.ID + 1

	TrueResident := model.CoordinationCity{
		ID:            newID,
		KorkabName:    newData.KorkabName,
		KorkabNik:     newData.KorkabNik,
		KorkabPhone:   newData.KorkabPhone,
		KorkabAge:     newData.KorkabAge,
		KorkabGender:  newData.KorkabGender,
		KorkabAddress: newData.KorkabAddress,
		KorkabCity:    newData.KorkabCity,
		KorkabNetwork: newData.KorkabNetwork,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}

	return nil
}

func (cc *coordinationcity) Update(ctx context.Context, ID int, updatedData dto.PayloadStoreCoordinatorCity) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	update := bson.M{
		"$set": bson.M{
			"korkab_name":    updatedData.KorkabName,
			"korkab_nik":     updatedData.KorkabNik,
			"korkab_phone":   updatedData.KorkabPhone,
			"korkab_age":     updatedData.KorkabAge,
			"korkab_gender":  updatedData.KorkabGender,
			"korkab_address": updatedData.KorkabAddress,
			"korkab_city":    updatedData.KorkabCity,
			"korkab_network": updatedData.KorkabNetwork,
			"updated_at":     time.Now(),
		},
	}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (cc *coordinationcity) Delete(ctx context.Context, ID int) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)
	filter := bson.M{"id": ID}

	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}

	return nil
}

func (cc *coordinationcity) Export(ctx context.Context, filter dto.CoordinationCityFilter) ([]byte, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "coordination_city"

	collection := cc.MongoConn.Database(dbName).Collection(collectionName)
	pipeline := make([]bson.M, 0)

	matchStage := bson.M{}

	if filter.KorkabCity != "" {
		matchStage["korkab_city"] = filter.KorkabCity
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

	var dataAllResident []dto.ExportCoordinatorCity
	if err = cursor.All(ctx, &dataAllResident); err != nil {
		return nil, err
	}

	xlsx := excelize.NewFile()
	sheetName := "Sheet1"
	xlsx.NewSheet(sheetName)

	if err := setHeaderAndStyle(xlsx, sheetName); err != nil {
		return nil, err
	}

	setDataRows(xlsx, sheetName, dataAllResident)

	fileBytes, err := xlsx.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return fileBytes.Bytes(), nil
}

func setHeaderAndStyle(xlsx *excelize.File, sheetName string) error {
	headers := []string{"NAMA", "NIK", "NOMOR HANDPHONE", "JK", "UMUR", "ALAMAT", "KABUPATEN", "JARINGAN", "TANGGAL BUAT", "TANGGAL UPDATE"}
	widths := []float64{30, 30, 25, 10, 10, 35, 20, 15, 20, 20}

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

func setDataRows(xlsx *excelize.File, sheetName string, dataAllResident []dto.ExportCoordinatorCity) {
	style, _ := xlsx.NewStyle(`{"alignment":{"horizontal":"center"}}`)

	for i, data := range dataAllResident {
		rowData := []interface{}{data.Nama, data.Nik, data.NoHandphone, data.Gender, data.Usia, data.Alamat, data.NamaKabupaten, data.Jaringan, data.CreatedAt.Format("2006-01-02"), data.UpdatedAt.Format("2006-01-02")}
		for j, value := range rowData {
			cell := fmt.Sprintf("%c%d", 'A'+j, i+2)
			xlsx.SetCellValue(sheetName, cell, value)
			xlsx.SetCellStyle(sheetName, cell, cell, style)
		}
	}
}

func (cc *coordinationcity) getLastData(ctx context.Context, collection *mongo.Collection) (model.Resident, error) {
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
