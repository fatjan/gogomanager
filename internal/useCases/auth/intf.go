package auth

import (
	"github.com/fatjan/gogomanager/internal/dto"
)

type UseCase interface {
	Login(*dto.AuthRequest) (*dto.AuthResponse, error)
	Register(*dto.AuthRequest) (*dto.AuthResponse, error)
}
