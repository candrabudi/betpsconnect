package factory

import (
	"betpsconnect/database"
	"betpsconnect/internal/repository"
)

type Factory struct {
	ResidentRepository repository.Resident
}

func NewFactory() *Factory {
	mongoConn := database.GetMongoConnection()
	return &Factory{
		ResidentRepository: repository.NewResidentRepository(mongoConn),
	}
}
