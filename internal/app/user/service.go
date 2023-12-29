package user

import (
	"betpsconnect/internal/dto"
	"betpsconnect/internal/factory"
	"betpsconnect/internal/model"
	"betpsconnect/internal/repository"
	"betpsconnect/pkg/constants"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type service struct {
	userRepository      repository.User
	userTokenRepository repository.UserToken
}

type Service interface {
	LoginUser(ctx context.Context, payload dto.PayloadLogin) (dto.ResultLogin, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		userRepository:      f.UserRepository,
		userTokenRepository: f.UserTokenRepository,
	}
}

func (s *service) LoginUser(ctx context.Context, payload dto.PayloadLogin) (dto.ResultLogin, error) {
	email := payload.Email
	password := payload.Password
	checkUser, err := s.userRepository.FindOneByEmail(ctx, email)
	if err == mongo.ErrNoDocuments {
		return dto.ResultLogin{}, constants.FailedLogin
	}
	token, err := s.userRepository.LoginUser(ctx, email, password)
	if err != nil {
		return dto.ResultLogin{}, constants.FailedLogin
	}

	payloadUserToken := model.UserToken{
		UserID: checkUser.ID,
		Token:  token,
	}
	err = s.userTokenRepository.Store(ctx, payloadUserToken)
	if err != nil {
		return dto.ResultLogin{}, constants.FailedLogin
	}
	resultLogin := dto.ResultLogin{
		ID:       checkUser.ID,
		FullName: checkUser.FullName,
		Email:    checkUser.Email,
		Regency:  checkUser.Regency,
		Status:   checkUser.Status,
		Role:     checkUser.Role,
		Token:    token,
	}
	return resultLogin, nil
}
