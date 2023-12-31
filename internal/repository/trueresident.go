package repository

import (
	"betpsconnect/internal/dto"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type TrueResident interface {
}

type trueresident struct {
	MongoConn *mongo.Client
}

func NewTrueResidentRepository(mongoConn *mongo.Client) TrueResident {
	return &trueresident{
		MongoConn: mongoConn,
	}
}

func (tr *trueresident) ResidentValidate(ctx context.Context, newData dto.TrueResidentPayload) error {
	return nil
}
