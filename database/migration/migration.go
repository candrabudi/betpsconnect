package migration

import (
	"betpsconnect/database"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collections = []struct {
	Name   string
	Schema interface{}
}{}

func Migrate() {
	conn := database.GetMongoConnection()
	ctx := context.TODO()

	for _, col := range collections {
		collection := conn.Database("your_database").Collection(col.Name)
		createCollectionIfNotExist(ctx, collection, col.Schema)
	}
}

func createCollectionIfNotExist(ctx context.Context, collection *mongo.Collection, schema interface{}) {
	names, err := collection.Indexes().List(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	exists := false
	for names.Next(ctx) {
		exists = true
		break
	}

	if !exists {
		opts := options.CreateIndexes().SetMaxTime(10 * time.Second)
		indexes := prepareIndexes()
		_, err := collection.Indexes().CreateMany(ctx, indexes, opts)
		if err != nil {
			fmt.Println("Failed to create indexes:", err)
			return
		}
	}
}

func prepareIndexes() []mongo.IndexModel {
	index1 := mongo.IndexModel{
		Keys:    bson.D{{Key: "field_name", Value: 1}},
		Options: options.Index().SetUnique(true),
	}

	return []mongo.IndexModel{index1}
}
