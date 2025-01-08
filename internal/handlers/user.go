package handlers

import (
	"net/http"

	"github.com/fatjan/gogomanager/internal/dto"
	"github.com/fatjan/gogomanager/internal/useCases/user"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Get(ginCtx *gin.Context)
}

type userHandler struct {
	userUseCase user.UseCase
}

func (r *userHandler) Get(ginCtx *gin.Context) {
	var userRequest dto.UserRequest

	if err := ginCtx.BindJSON(&userRequest); err != nil {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userResponse, err := r.userUseCase.GetUser(&userRequest)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, userResponse)
}

func NewUserHandler(userUseCase user.UseCase) UserHandler {
	return &userHandler{userUseCase: userUseCase}
}
