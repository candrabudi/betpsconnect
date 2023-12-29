package repository

import (
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type City interface {
	GetCity(ctx context.Context) ([]string, error)
}

type city struct {
	MongoConn *mongo.Client
}

func NewCityRepository(mongoConn *mongo.Client) City {
	return &city{
		MongoConn: mongoConn,
	}
}
func (c *city) GetCity(ctx context.Context) ([]string, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "cities"

	collection := c.MongoConn.Database(dbName).Collection(collectionName)

	// Melakukan pencarian tanpa filter
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Gunakan map untuk menghindari duplikat nama kabupaten
	cityMap := make(map[string]bool)
	for cursor.Next(ctx) {
		var ddistrict model.District
		if err := cursor.Decode(&ddistrict); err != nil {
			return nil, err
		}
		cityMap[ddistrict.NamaKabupaten] = true
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	// Ubah map menjadi array string yang diinginkan
	var cities []string
	for city := range cityMap {
		cities = append(cities, city)
	}

	return cities, nil
}
