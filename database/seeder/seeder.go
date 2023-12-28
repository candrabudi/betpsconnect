package seeder

import (
	"betpsconnect/database"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uint      `bson:"_id,omitempty"`
	Name      string    `bson:"name"`
	Email     string    `bson:"email"`
	Role      string    `bson:"role"`
	Password  string    `bson:"password"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func Seed() {
	mongoConn := database.GetMongoConnection()
	ctx := context.Background()

	passwordSA := []byte("superadminpassword")
	hashedPasswordSA, err := bcrypt.GenerateFromPassword(passwordSA, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	passwordAdmin := []byte("adminpassword")
	hashedPasswordAdmin, err := bcrypt.GenerateFromPassword(passwordAdmin, bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	seedData := []User{
		{
			Name:      "Super Admin",
			Email:     "superadmin@gmail.com",
			Role:      "superadmin",
			Password:  string(hashedPasswordSA),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Name:      "Admin",
			Email:     "admin@gmail.com",
			Role:      "admin",
			Password:  string(hashedPasswordAdmin),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	usersCollection := mongoConn.Database("your_database").Collection("users") // Ganti dengan nama database dan koleksi Anda

	for _, data := range seedData {
		var user User
		err := usersCollection.FindOne(ctx, bson.M{"email": data.Email}).Decode(&user)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				_, err := usersCollection.InsertOne(ctx, data)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else {
				fmt.Println(err)
				return
			}
		}
	}
}
