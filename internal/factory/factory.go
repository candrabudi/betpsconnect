package factory

import (
	"betpsconnect/database"
	"betpsconnect/internal/repository"
)

type Factory struct {
	ResidentRepository repository.Resident
	DistrictRepository repository.District
}

func NewFactory() *Factory {
	mongoConn := database.GetMongoConnection()
	return &Factory{
		DistrictRepository: repository.NewDistrictRepository(mongoConn),
	}
}
