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
	GetAll(ctx context.Context, limit, offset int64, filter dto.TrueResidentFilter, userSess any) (dto.ResultAllTrueResident, error)
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

func (s *service) Store(ctx context.Context, payload dto.TrueResidentPayload) error {

	err := s.trueResidentRepository.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
