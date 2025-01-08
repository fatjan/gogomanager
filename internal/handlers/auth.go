package handlers

import (
	"net/http"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/useCases/auth"
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Post(ginCtx *gin.Context)
}

type authHandler struct {
	authUseCase auth.UseCase
}

func NewAuthHandler(authUsecase auth.UseCase) AuthHandler {
	return &authHandler{authUseCase: authUsecase}
}

func (r *authHandler) Post(ginCtx *gin.Context) {
	var authRequest dto.AuthRequest
	if err := ginCtx.BindJSON(&authRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if authRequest.Action == string(dto.Create) {
		authResponse, err := r.authUseCase.Register(&authRequest)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ginCtx.JSON(http.StatusCreated, authResponse)
		return
	} else if authRequest.Action == string(dto.Login) {
		authResponse, err := r.authUseCase.Login(&authRequest)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ginCtx.JSON(http.StatusOK, authResponse)
		return
	} else {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid action"})
		return
	}
}
