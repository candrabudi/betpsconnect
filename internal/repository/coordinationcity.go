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

	xlsx.SetColWidth(sheetName, "A", "A", 30)
	xlsx.SetColWidth(sheetName, "B", "B", 30)
	xlsx.SetColWidth(sheetName, "C", "C", 25)
	xlsx.SetColWidth(sheetName, "D", "D", 10)
	xlsx.SetColWidth(sheetName, "E", "E", 35)
	xlsx.SetColWidth(sheetName, "F", "F", 20)
	xlsx.SetColWidth(sheetName, "G", "G", 15)
	xlsx.SetColWidth(sheetName, "H", "H", 20)
	xlsx.SetColWidth(sheetName, "I", "I", 20)

	style, err := xlsx.NewStyle(`{"alignment":{"horizontal":"center"}}`)
	if err != nil {
		return nil, err
	}

	xlsx.SetCellValue(sheetName, "A1", "NAMA")
	xlsx.SetCellStyle(sheetName, "A1", "A1", style)
	xlsx.SetCellValue(sheetName, "B1", "NIK")
	xlsx.SetCellStyle(sheetName, "B1", "B1", style)
	xlsx.SetCellValue(sheetName, "C1", "NOMOR HANDPHONE")
	xlsx.SetCellStyle(sheetName, "C1", "C1", style)
	xlsx.SetCellValue(sheetName, "D1", "UMUR")
	xlsx.SetCellStyle(sheetName, "D1", "D1", style)
	xlsx.SetCellValue(sheetName, "E1", "ALAMAT")
	xlsx.SetCellStyle(sheetName, "E1", "E1", style)
	xlsx.SetCellValue(sheetName, "F1", "KABUPATEN")
	xlsx.SetCellStyle(sheetName, "F1", "F1", style)
	xlsx.SetCellValue(sheetName, "G1", "JARINGAN")
	xlsx.SetCellStyle(sheetName, "G1", "G1", style)
	xlsx.SetCellValue(sheetName, "H1", "TANGGAL BUAT")
	xlsx.SetCellStyle(sheetName, "H1", "H1", style)
	xlsx.SetCellValue(sheetName, "I1", "TANGGAL UPDATE")
	xlsx.SetCellStyle(sheetName, "I1", "I1", style)

	for i, data := range dataAllResident {
		xlsx.SetCellValue(sheetName, fmt.Sprintf("A%d", i+2), data.Nama)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("B%d", i+2), data.Nik)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("C%d", i+2), data.NoHandphone)
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("C%d", i+2), fmt.Sprintf("C%d", i+2), style)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("D%d", i+2), data.Usia)
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("D%d", i+2), fmt.Sprintf("D%d", i+2), style)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("E%d", i+2), data.Alamat)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("F%d", i+2), data.NamaKabupaten)
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("F%d", i+2), fmt.Sprintf("F%d", i+2), style)
		xlsx.SetCellValue(sheetName, fmt.Sprintf("G%d", i+2), data.Jaringan)
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("G%d", i+2), fmt.Sprintf("G%d", i+2), style)

		createdAtFormatted := data.CreatedAt.Format("2006-01-02")
		xlsx.SetCellValue(sheetName, fmt.Sprintf("H%d", i+2), createdAtFormatted)
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("H%d", i+2), fmt.Sprintf("H%d", i+2), style)

		updatedAtFormatted := data.UpdatedAt.Format("2006-01-02")
		xlsx.SetCellValue(sheetName, fmt.Sprintf("I%d", i+2), updatedAtFormatted)
		xlsx.SetCellStyle(sheetName, fmt.Sprintf("I%d", i+2), fmt.Sprintf("I%d", i+2), style)
	}

	fileBytes, err := xlsx.WriteToBuffer()
	if err != nil {
		return nil, err
	}

	return fileBytes.Bytes(), nil
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
