package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type UserToken interface {
}

type usertoken struct {
	MongoConn *mongo.Client
}

func NewUserTokenRepository(mongoConn *mongo.Client) UserToken {
	return &usertoken{
		MongoConn: mongoConn,
	}
}
