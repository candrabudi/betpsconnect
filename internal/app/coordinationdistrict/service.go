package coordinationdistrict

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
	"errors"
)

type service struct {
	coordinationDistrict repository.CoordinationDistrict
}

type Service interface {
	GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationDistrictFilter, userSess any) (dto.ResultAllCoordinatorDistrict, error)
	Store(ctx context.Context, payload dto.PayloadStoreCoordinatorDistrict) error
	Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorDistrict) error
	Delete(ctx context.Context, ID int) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		coordinationDistrict: f.CoordinationDistrictRepository,
	}
}

func (s *service) GetAll(ctx context.Context, limit, offset int64, filter dto.CoordinationDistrictFilter, userSess any) (dto.ResultAllCoordinatorDistrict, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultAllCoordinatorDistrict{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultTpsResidents, err := s.coordinationDistrict.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultAllCoordinatorDistrict{
			Items:    []dto.FindCoordinatorDistrict{},
			Metadata: dto.MetaData{},
		}, err
	}

	if len(resultTpsResidents.Items) == 0 {
		return dto.ResultAllCoordinatorDistrict{
			Items:    []dto.FindCoordinatorDistrict{},
			Metadata: dto.MetaData{},
		}, nil
	}
	return resultTpsResidents, nil
}

func (s *service) Store(ctx context.Context, payload dto.PayloadStoreCoordinatorDistrict) error {

	err := s.coordinationDistrict.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, ID int, payload dto.PayloadUpdateCoordinatorDistrict) error {
	err := s.coordinationDistrict.Update(ctx, ID, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Delete(ctx context.Context, ID int) error {
	err := s.coordinationDistrict.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
