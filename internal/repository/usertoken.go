package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserToken interface {
	Store(ctx context.Context, data model.UserToken) error
	FindToken(ctx context.Context, Token string) (dto.FindOneUserToken, error)
}

type usertoken struct {
	MongoConn *mongo.Client
}

func NewUserTokenRepository(mongoConn *mongo.Client) UserToken {
	return &usertoken{
		MongoConn: mongoConn,
	}
}

func (ut *usertoken) Store(ctx context.Context, data model.UserToken) error {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "user_tokens"

	collection := ut.MongoConn.Database(dbName).Collection(collectionName)

	_, err := collection.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func (ut *usertoken) FindToken(ctx context.Context, Token string) (dto.FindOneUserToken, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "user_tokens"

	collection := ut.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	bsonFilter["token"] = Token
	var duser model.UserToken
	err := collection.FindOne(ctx, bsonFilter).Decode(&duser)
	if err != nil {
		return dto.FindOneUserToken{}, err
	}
	userDTO := dto.FindOneUserToken{
		ID:     duser.ID,
		UserID: duser.UserID,
		Token:  duser.Token,
	}

	return userDTO, nil
}
