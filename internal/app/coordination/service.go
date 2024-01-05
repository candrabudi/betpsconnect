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
	coordinationRepository repository.Coordination
}

type Service interface {
	GetListCoordinationCity(ctx context.Context, limit, offset int64, filter dto.ResidentFilter, userSess any) (dto.ResultAllCoordinatorCity, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		coordinationRepository: f.CoordinationRepository,
	}
}

func (s *service) GetListCoordinationCity(ctx context.Context, limit, offset int64, filter dto.ResidentFilter, userSess any) (dto.ResultAllCoordinatorCity, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultAllCoordinatorCity{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultTpsResidents, err := s.coordinationRepository.GetAll(ctx, limit, offset, filter)
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
