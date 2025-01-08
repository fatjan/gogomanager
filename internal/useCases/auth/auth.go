package auth

import (
	"github.com/fatjan/gogomanager/internal/config"
	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/models"
	"github.com/fatjan/gogomanager/internal/pkg/jwt_helper"
	"github.com/fatjan/gogomanager/internal/repositories/auth"
)

type useCase struct {
	authRepository auth.Repository
	cfg            *config.Config
}

func NewUseCase(authRepository auth.Repository, cfg *config.Config) UseCase {
	return &useCase{
		authRepository: authRepository,
		cfg:            cfg,
	}
}

func (uc *useCase) Login(authRequest *dto.AuthRequest) (*dto.AuthResponse, error) {
	err := authRequest.ValidatePayloadAuth()
	if err != nil {
		return nil, err
	}
	authRequest.SetName()
	authRequest.HashPassword()

	manager, err := uc.authRepository.FindByEmail(authRequest.Email)
	if err != nil {
		return nil, err
	}

	err = authRequest.ComparePassword(manager.Password)
	if err != nil {
		return nil, err
	}

	token, err := jwt_helper.SignJwt(uc.cfg.JwtKey, manager.ID)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Email: authRequest.Email,
		Token: token,
	}, nil
}

func (uc *useCase) Register(authRequest *dto.AuthRequest) (*dto.AuthResponse, error) {
	err := authRequest.ValidatePayloadAuth()
	if err != nil {
		return nil, err
	}
	authRequest.SetName()
	err = authRequest.HashPassword()
	if err != nil {
		return nil, err
	}

	newAuth := &models.Manager{
		Email:    authRequest.Email,
		Password: authRequest.Password,
		Name:     authRequest.Name,
	}

	id, err := uc.authRepository.Post(newAuth)
	if err != nil {
		return nil, err
	}

	token, err := jwt_helper.SignJwt(uc.cfg.JwtKey, id)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Email: authRequest.Email,
		Token: token,
	}, nil
}
