package handlers

import (
	"github.com/fatjan/gogomanager/internal/config"
	duckRepository "github.com/fatjan/gogomanager/internal/repositories/duck"
	duckUseCase "github.com/fatjan/gogomanager/internal/useCases/duck"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(cfg *config.Config, db *sqlx.DB, r *gin.Engine) {
	duckRepository := duckRepository.NewDuckRepository(db)
	duckUseCase := duckUseCase.NewUseCase(duckRepository)
	duckHandler := NewDuckHandler(duckUseCase)

	v1 := r.Group("v1")
	duckRouter := v1.Group("ducks")
	duckRouter.GET("/", duckHandler.Index)
	duckRouter.GET("/:id", duckHandler.Detail)
}
