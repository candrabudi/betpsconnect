package database

import (
	"betpsconnect/pkg/util"
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoConn *mongo.Client
	once      sync.Once
)

type mongoConfig struct {
	URI      string
	Name     string
	User     string
	Password string
}

func CreateMongoConnection() {
	// Create MongoDB configuration information
	mongoConf := mongoConfig{
		URI:      "mongodb://localhost:27017",
		Name:     "tpsconnect_dev",
		User:     "",
		Password: "",
	}

	if uri := util.GetEnv("MONGO_URI", ""); uri != "" {
		mongoConf.URI = uri
	}
	if dbName := util.GetEnv("MONGO_DB_NAME", ""); dbName != "" {
		mongoConf.Name = dbName
	}
	if user := util.GetEnv("MONGO_USER", ""); user != "" {
		mongoConf.User = user
	}
	if pass := util.GetEnv("MONGO_PASS", ""); pass != "" {
		mongoConf.Password = pass
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoConf.URI)

	// Check if user and password are provided
	if mongoConf.User != "" && mongoConf.Password != "" {
		clientOptions.Auth = &options.Credential{
			Username: mongoConf.User,
			Password: mongoConf.Password,
		}
	} else {
		fmt.Println("Skipping authentication as username or password not provided")
	}

	// Set connection context timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	mongoConn, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}

	// Check the connection
	err = mongoConn.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!")
}

func GetMongoConnection() *mongo.Client {
	// Check MongoDB connection, if exists return the client instance
	if mongoConn == nil {
		CreateMongoConnection()
	}
	return mongoConn
}
