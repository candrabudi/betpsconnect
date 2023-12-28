package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type User interface {
}

type user struct {
	MongoConn *mongo.Client
}

func NewUserRepository(mongoConn *mongo.Client) User {
	return &user{
		MongoConn: mongoConn,
	}
}
