package district

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
	"errors"
)

type service struct {
	districtRepository repository.District
}

type Service interface {
	GetDistrictByCity(ctx context.Context, filter dto.GetByCity, userSess any) ([]string, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		districtRepository: f.DistrictRepository,
	}
}

func (s *service) GetDistrictByCity(ctx context.Context, filter dto.GetByCity, userSess any) ([]string, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return []string{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultDistrictByCity, err := s.districtRepository.GetByCity(ctx, filter)
	if err != nil {
		return nil, err
	}
	return resultDistrictByCity, nil
}
