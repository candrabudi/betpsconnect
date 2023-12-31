package trueresident

import (
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
)

type service struct {
	trueResidentRepository repository.TrueResident
}

type Service interface {
}

func NewService(f *factory.Factory) Service {
	return &service{
		trueResidentRepository: f.TrueResidentRepository,
	}
}
