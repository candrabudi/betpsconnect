package repository

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/model"
	"betpsconnect/pkg/util"
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User interface {
	LoginUser(ctx context.Context, email, password string) (string, error)
	FindOneByEmail(ctx context.Context, email string) (dto.FindOneUser, error)
	FindOne(ctx context.Context, UserID int) (dto.FindOneUser, error)
}

type user struct {
	MongoConn *mongo.Client
}

func NewUserRepository(mongoConn *mongo.Client) User {
	return &user{
		MongoConn: mongoConn,
	}
}

type Claims struct {
	UserID string `json:"userID"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func (u *user) LoginUser(ctx context.Context, email, password string) (string, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "users"

	collection := u.MongoConn.Database(dbName).Collection(collectionName)

	var user model.User
	filter := bson.M{"email": email}
	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", err
	}

	userID := strconv.Itoa(user.ID)
	secretKey := []byte(util.GetEnv("JWT_KEY", "fallback"))
	jwt, err := u.GenerateToken(secretKey, userID, user.Email)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func (u *user) GenerateToken(secretKey []byte, userID string, email string) (string, error) {
	loc, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		return "", err
	}

	var expiredTime int64
	jwtMode := util.GetEnv("JWT_MODE", "fallback")
	if jwtMode == "release" {
		expiredTime = time.Now().In(loc).Add(time.Hour * 2191).Unix()
	} else {
		expiredTime = time.Now().In(loc).Add(time.Hour * 730).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     expiredTime,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (u *user) FindOneByEmail(ctx context.Context, email string) (dto.FindOneUser, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "users"

	collection := u.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	bsonFilter["email"] = email
	var duser model.User
	err := collection.FindOne(ctx, bsonFilter).Decode(&duser)
	if err != nil {
		return dto.FindOneUser{}, err
	}
	fmt.Println(duser)
	userDTO := dto.FindOneUser{
		ID:       duser.ID,
		FullName: duser.FullName,
		Regency:  duser.Regency,
		Status:   duser.Status,
		Role:     duser.Role,
		Email:    duser.Email,
	}

	return userDTO, nil
}

func (u *user) FindOne(ctx context.Context, UserID int) (dto.FindOneUser, error) {
	dbName := util.GetEnv("MONGO_DB_NAME", "tpsconnect_dev")
	collectionName := "users"

	collection := u.MongoConn.Database(dbName).Collection(collectionName)
	bsonFilter := bson.M{}
	bsonFilter["id"] = UserID
	var duser model.User
	err := collection.FindOne(ctx, bsonFilter).Decode(&duser)
	if err != nil {
		return dto.FindOneUser{}, err
	}
	fmt.Println(duser)
	userDTO := dto.FindOneUser{
		ID:       duser.ID,
		FullName: duser.FullName,
		Regency:  duser.Regency,
		Status:   duser.Status,
		Role:     duser.Role,
		Email:    duser.Email,
	}

	return userDTO, nil
}
