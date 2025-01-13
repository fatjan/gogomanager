package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/fatjan/gogomanager/internal/dto"
	urlValidator "github.com/fatjan/gogomanager/internal/pkg/validator"
	"github.com/fatjan/gogomanager/internal/useCases/user"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler interface {
	Get(ginCtx *gin.Context)
	Update(ginCtx *gin.Context)
}

type userHandler struct {
	userUseCase user.UseCase
}

func (r *userHandler) Get(ginCtx *gin.Context) {
	managerId, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}

	id := managerId.(int)
	userRequest := dto.UserRequest{
		UserID: id,
	}

	userResponse, err := r.userUseCase.GetUser(&userRequest)
	if err != nil {
		ginCtx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, userResponse)
}
func (r *userHandler) Update(ginCtx *gin.Context) {
	if ginCtx.ContentType() != "application/json" {
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid content type"})
		return
	}

	var userRequest dto.UserPatchRequest

	managerId, exists := ginCtx.Get("manager_id")
	if !exists {
		ginCtx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid manager id"})
		return
	}

	userIDInt := managerId.(int)

	if err := ginCtx.ShouldBindJSON(&userRequest); err != nil {
		log.Println(err.Error())
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var validate = validator.New()
	_ = validate.RegisterValidation("url", urlValidator.StrictURLValidation)

	if err := validate.Struct(userRequest); err != nil {
		log.Println(err.Error())
		ginCtx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := r.userUseCase.UpdateUser(ginCtx.Request.Context(), userIDInt, &userRequest)
	if err != nil {
		statusRes := http.StatusInternalServerError
		errorMessageRes := errors.New("internal server error")
		if errors.Is(err, sql.ErrNoRows) {
			log.Println(fmt.Sprintf("user with id %d not found", userIDInt))
			statusRes = http.StatusNotFound
			errorMessageRes = errors.New(fmt.Sprintf("user with id %d not found", userIDInt))
		}

		switch err.Error() {
		case "duplicate email":
			statusRes = http.StatusConflict
			errorMessageRes = errors.New("email already exists")
		}

		ginCtx.JSON(statusRes, gin.H{"error": errorMessageRes.Error()})
		return
	}

	ginCtx.JSON(http.StatusOK, userResponse)
}

func NewUserHandler(userUseCase user.UseCase) UserHandler {
	return &userHandler{userUseCase: userUseCase}
}
