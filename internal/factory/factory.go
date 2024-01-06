package factory

import (
	"betpsconnect/database"
	"betpsconnect/internal/repository"
)

type Factory struct {
	ResidentRepository                repository.Resident
	DistrictRepository                repository.District
	SubDistrictRepository             repository.SubDistrict
	CityRepository                    repository.City
	UserRepository                    repository.User
	UserTokenRepository               repository.UserToken
	TrueResidentRepository            repository.TrueResident
	CoordinationCityRepository        repository.CoordinationCity
	CoordinationDistrictRepository    repository.CoordinationDistrict
	CoordinationSubdistrictRepository repository.CoordinationSubdistrict
	CoordinationTpsRepository         repository.CoordinationTps
}

func NewFactory() *Factory {
	mongoConn := database.GetMongoConnection()
	return &Factory{
		ResidentRepository:                repository.NewResidentRepository(mongoConn),
		DistrictRepository:                repository.NewDistrictRepository(mongoConn),
		SubDistrictRepository:             repository.NewSubDistrictRepository(mongoConn),
		CityRepository:                    repository.NewCityRepository(mongoConn),
		UserRepository:                    repository.NewUserRepository(mongoConn),
		UserTokenRepository:               repository.NewUserTokenRepository(mongoConn),
		TrueResidentRepository:            repository.NewTrueResidentRepository(mongoConn),
		CoordinationCityRepository:        repository.NewCoordinationCityRepository(mongoConn),
		CoordinationDistrictRepository:    repository.NewCoordinationDistrictRepository(mongoConn),
		CoordinationSubdistrictRepository: repository.NewCoordinationSubdistrictRepository(mongoConn),
		CoordinationTpsRepository:         repository.NewCoordinationTpsRepository(mongoConn),
	}
}
