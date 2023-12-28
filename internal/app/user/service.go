package user

import (
	"betpsconnect/internal/factory"
	"betpsconnect/internal/repository"
)

type service struct {
	userRepository repository.City
}

type Service interface {
}

func NewService(f *factory.Factory) Service {
	return &service{
		userRepository: f.UserRepository,
	}
}
