package coordinationtps

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
	"errors"
)

type service struct {
	CoordinationTps repository.CoordinationTps
}

type Service interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationTpsFilter, userSess any) (dto.ResultAllCoordinatorTps, error)
	Store(ctx context.Context, payload dto.PayloadStoreCoordinatorTps) error
	Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorTps) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		CoordinationTps: f.CoordinationTpsRepository,
	}
}

func (s *service) GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationTpsFilter, userSess any) (dto.ResultAllCoordinatorTps, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultAllCoordinatorTps{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultTpsResidents, err := s.CoordinationTps.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultAllCoordinatorTps{
			Items:    []dto.FindCoordinatorTps{},
			Metadata: dto.MetaData{},
		}, err
	}

	if len(resultTpsResidents.Items) == 0 {
		return dto.ResultAllCoordinatorTps{
			Items:    []dto.FindCoordinatorTps{},
			Metadata: dto.MetaData{},
		}, nil
	}
	return resultTpsResidents, nil
}

func (s *service) Store(ctx context.Context, payload dto.PayloadStoreCoordinatorTps) error {

	err := s.CoordinationTps.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorTps) error {
	err := s.CoordinationTps.Update(ctx, ID, payload)
	if err != nil {
		return err
	}

	return nil
}
