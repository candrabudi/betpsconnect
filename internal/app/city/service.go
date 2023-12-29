package city

import (
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
	"context"
)

type service struct {
	cityRepository repository.City
}

type Service interface {
	GetCity(ctx context.Context) ([]string, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		cityRepository: f.CityRepository,
	}
}

func (s *service) GetCity(ctx context.Context) ([]string, error) {
	// Menggunakan fungsi GetCity dari district untuk mendapatkan daftar nama kabupaten
	cities, err := s.cityRepository.GetCity(ctx)
	if err != nil {
		return nil, err
	}

	return cities, nil
}
