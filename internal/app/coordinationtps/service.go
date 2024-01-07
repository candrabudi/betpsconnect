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
	coordinationTps repository.CoordinationTps
}

type Service interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationTpsFilter, userSess any) (dto.ResultAllCoordinatorTps, error)
	Store(ctx context.Context, payload dto.PayloadStoreCoordinatorTps) error
	Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorTps) error
	Delete(ctx context.Context, ID int) error
	Export(ctx context.Context, filter dto.CoordinationTpsFilter, userSess any) ([]byte, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		coordinationTps: f.CoordinationTpsRepository,
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

	resultTpsResidents, err := s.coordinationTps.GetAll(ctx, limit, offset, filter)
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

func (s *service) Export(ctx context.Context, filter dto.CoordinationTpsFilter, userSess any) ([]byte, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return nil, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	export, err := s.coordinationTps.Export(ctx, filter)
	if err != nil {
		return nil, err
	}

	return export, nil
}

func (s *service) Store(ctx context.Context, payload dto.PayloadStoreCoordinatorTps) error {

	err := s.coordinationTps.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorTps) error {
	err := s.coordinationTps.Update(ctx, ID, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, ID int) error {
	err := s.coordinationTps.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
