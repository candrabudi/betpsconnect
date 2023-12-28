package subdistrict

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
	"context"
)

type service struct {
	subDistrictRepository repository.SubDistrict
}

type Service interface {
	GetByDistrict(ctx context.Context, filter dto.GetByDistrict) ([]dto.GetByDistrict, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		subDistrictRepository: f.SubDistrictRepository,
	}
}

func (s *service) GetByDistrict(ctx context.Context, filter dto.GetByDistrict) ([]dto.GetByDistrict, error) {
	resultSubDistrictByDistrict, err := s.subDistrictRepository.GetByDistrict(ctx, filter)
	if err != nil {
		return nil, err
	}

	return resultSubDistrictByDistrict, nil
}
