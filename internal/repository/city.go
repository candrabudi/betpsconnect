package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type City interface {
}

type city struct {
	MongoConn *mongo.Client
}

func NewCityRepository(mongoConn *mongo.Client) City {
	return &city{
		MongoConn: mongoConn,
	}
}
