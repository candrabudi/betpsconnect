package coordinationcity

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
	"errors"
)

type service struct {
	coordinationCityRepository repository.CoordinationCity
}

type Service interface {
	GetListCoordinationCity(ctx context.Context, limit, offset int64, filter dto.CoordinationCityFilter, userSess any) (dto.ResultAllCoordinatorCity, error)
	Store(ctx context.Context, payload dto.PayloadStoreCoordinatorCity) error
	Update(ctx context.Context, ID int, payload dto.PayloadStoreCoordinatorCity) error
	Delete(ctx context.Context, ID int) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		coordinationCityRepository: f.CoordinationCityRepository,
	}
}

func (s *service) GetListCoordinationCity(ctx context.Context, limit, offset int64, filter dto.CoordinationCityFilter, userSess any) (dto.ResultAllCoordinatorCity, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultAllCoordinatorCity{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.KorkabCity = user.Regency
	}

	resultTpsResidents, err := s.coordinationCityRepository.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultAllCoordinatorCity{
			Items:    []dto.FindCoordinatorCity{},
			Metadata: dto.MetaData{},
		}, err
	}

	if len(resultTpsResidents.Items) == 0 {
		return dto.ResultAllCoordinatorCity{
			Items:    []dto.FindCoordinatorCity{},
			Metadata: dto.MetaData{},
		}, nil
	}
	return resultTpsResidents, nil
}

func (s *service) Store(ctx context.Context, payload dto.PayloadStoreCoordinatorCity) error {

	err := s.coordinationCityRepository.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, ID int, payload dto.PayloadStoreCoordinatorCity) error {

	err := s.coordinationCityRepository.Update(ctx, ID, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, ID int) error {
	err := s.coordinationCityRepository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
