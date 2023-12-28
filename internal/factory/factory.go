package factory

import (
	"betpsconnect/database"
	"betpsconnect/internal/repository"
)

type Factory struct {
	ResidentRepository    repository.Resident
	DistrictRepository    repository.District
	SubDistrictRepository repository.SubDistrict
}

func NewFactory() *Factory {
	mongoConn := database.GetMongoConnection()
	return &Factory{
		ResidentRepository:    repository.NewResidentRepository(mongoConn),
		DistrictRepository:    repository.NewDistrictRepository(mongoConn),
		SubDistrictRepository: repository.NewSubDistrictRepository(mongoConn),
	}
}
