package district

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
	"context"
)

type service struct {
	districtRepository repository.District
}

type Service interface {
	GetDistrictByCity(ctx context.Context, filter dto.GetByCity) ([]string, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		districtRepository: f.DistrictRepository,
	}
}

func (s *service) GetDistrictByCity(ctx context.Context, filter dto.GetByCity) ([]string, error) {
	resultDistrictByCity, err := s.districtRepository.GetByCity(ctx, filter)
	if err != nil {
		return nil, err
	}
	return resultDistrictByCity, nil
}
