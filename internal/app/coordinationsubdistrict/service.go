package coordinationsubdistrict

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
	"errors"
)

type service struct {
	coordinationSubdistrict repository.CoordinationSubdistrict
}

type Service interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationSubdistrictFilter, userSess any) (dto.ResultAllCoordinatorSubdistrict, error)
	Store(ctx context.Context, payload dto.PayloadStoreCoordinatorSubdistrict) error
	Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorSubdistrict) error
	Delete(ctx context.Context, ID int) error
	Export(ctx context.Context, filter dto.CoordinationSubdistrictFilter, userSess any) ([]byte, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		coordinationSubdistrict: f.CoordinationSubdistrictRepository,
	}
}

func (s *service) GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationSubdistrictFilter, userSess any) (dto.ResultAllCoordinatorSubdistrict, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultAllCoordinatorSubdistrict{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultTpsResidents, err := s.coordinationSubdistrict.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultAllCoordinatorSubdistrict{
			Items:    []dto.FindCoordinatorSubdistrict{},
			Metadata: dto.MetaData{},
		}, err
	}

	if len(resultTpsResidents.Items) == 0 {
		return dto.ResultAllCoordinatorSubdistrict{
			Items:    []dto.FindCoordinatorSubdistrict{},
			Metadata: dto.MetaData{},
		}, nil
	}
	return resultTpsResidents, nil
}

func (s *service) Export(ctx context.Context, filter dto.CoordinationSubdistrictFilter, userSess any) ([]byte, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return nil, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	export, err := s.coordinationSubdistrict.Export(ctx, filter)
	if err != nil {
		return nil, err
	}

	return export, nil
}

func (s *service) Store(ctx context.Context, payload dto.PayloadStoreCoordinatorSubdistrict) error {

	err := s.coordinationSubdistrict.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorSubdistrict) error {
	err := s.coordinationSubdistrict.Update(ctx, ID, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, ID int) error {
	err := s.coordinationSubdistrict.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
