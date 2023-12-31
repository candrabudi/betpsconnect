package trueresident

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
	"context"
)

type service struct {
	trueResidentRepository repository.TrueResident
}

type Service interface {
	Store(ctx context.Context, payload dto.TrueResidentPayload) error
}

func NewService(f *factory.Factory) Service {
	return &service{
		trueResidentRepository: f.TrueResidentRepository,
	}
}

func (s *service) Store(ctx context.Context, payload dto.TrueResidentPayload) error {

	err := s.trueResidentRepository.Store(ctx, payload)
	if err != nil {
		return err
	}

	return nil
}
