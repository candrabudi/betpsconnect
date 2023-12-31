package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TrueResident interface {
	Store(ctx context.Context, newData dto.TrueResidentPayload) error
}

type trueresident struct {
	MongoConn *mongo.Client
}

func NewTrueResidentRepository(mongoConn *mongo.Client) TrueResident {
	return &trueresident{
		MongoConn: mongoConn,
	}
}

func (tr *trueresident) Store(ctx context.Context, newData dto.TrueResidentPayload) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collection := tr.MongoConn.Database(dbName).Collection("true_residents")
	trueFilter := bson.M{"nik": newData.Nik}
	var mresident model.TrueResident
	err := collection.FindOne(ctx, trueFilter).Decode(&mresident)

	if err == nil {
		return err
	}
	person, err := tr.getLastPerson(ctx, collection)
	if err != nil {
		return err
	}
	newUserID := person.ID + 1
	TrueResident := model.TrueResident{
		ID:          newUserID,
		ResidentID:  0,
		FullName:    newData.FullName,
		Nik:         newData.Nik,
		NoHandphone: newData.NoHandphone,
		Age:         newData.Age,
		Gender:      newData.Gender,
		BirthDate:   "",
		BirthPlace:  "",
		City:        newData.City,
		District:    newData.District,
		SubDistrict: newData.Subdistrict,
		Tps:         newData.TPS,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, errInsert := collection.InsertOne(ctx, TrueResident)
	if errInsert != nil {
		return errInsert
	}
	return nil
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
