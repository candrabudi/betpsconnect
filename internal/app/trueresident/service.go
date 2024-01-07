package trueresident

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"context"
	"errors"
)

type service struct {
	trueResidentRepository repository.TrueResident
}

type Service interface {
	Store(ctx context.Context, payload dto.TrueResidentPayload) error
	Update(ctx context.Context, ID int, payload dto.PayloadUpdateTrueResident) error
	GetAll(ctx context.Context, limit, offset int64, filter dto.TrueResidentFilter, userSess any) (dto.ResultAllTrueResident, error)
	GetTpsOnValidResident(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error)
	Delete(ctx context.Context, ID int) error
	Export(ctx context.Context, filter dto.TrueResidentFilter, userSess any) ([]byte, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		trueResidentRepository: f.TrueResidentRepository,
	}
}

func (s *service) GetAll(ctx context.Context, limit, offset int64, filter dto.TrueResidentFilter, userSess any) (dto.ResultAllTrueResident, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return dto.ResultAllTrueResident{}, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	resultTpsResidents, err := s.trueResidentRepository.GetAll(ctx, limit, offset, filter)
	if err != nil {
		return dto.ResultAllTrueResident{}, err
	}
	return resultTpsResidents, nil
}

func (s *service) Export(ctx context.Context, filter dto.TrueResidentFilter, userSess any) ([]byte, error) {
	user, ok := userSess.(model.User)
	if !ok {
		return nil, errors.New("invalid user session data")
	}

	if user.Role == "admin" {
		filter.NamaKabupaten = user.Regency
	}

	export, err := s.trueResidentRepository.Export(ctx, filter)
	if err != nil {
		return nil, err
	}

	return export, nil
}

func (s *service) Store(ctx context.Context, payload dto.TrueResidentPayload) error {

	err := s.trueResidentRepository.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) Update(ctx context.Context, ID int, payload dto.PayloadUpdateTrueResident) error {

	err := s.trueResidentRepository.Update(ctx, ID, payload)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetTpsOnValidResident(ctx context.Context, filter dto.FindTpsByDistrict) ([]string, error) {
	resultTps, err := s.trueResidentRepository.GetTpsOnValidResident(ctx, filter)
	if err != nil {
		return []string{}, err
	}
	if len(resultTps) == 0 {
		return []string{}, nil
	}
	return resultTps, nil
}

func (s *service) Delete(ctx context.Context, ID int) error {
	err := s.trueResidentRepository.Delete(ctx, ID)
	if err != nil {
		return err
	}

	return nil
}
