package city

import (
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
)

type service struct {
	cityRepository repository.City
}

type Service interface {
}

func NewService(f *factory.Factory) Service {
	return &service{
		cityRepository: f.CityRepository,
	}
}
