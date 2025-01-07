package handlers

import (
	"github.com/fatjan/gogomanager/internal/config"
	departmentRepository "github.com/fatjan/gogomanager/internal/repositories/department"
	duckRepository "github.com/fatjan/gogomanager/internal/repositories/duck"
	departmentUseCase "github.com/fatjan/gogomanager/internal/useCases/department"
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

	departmentRepository := departmentRepository.NewDepartmentRepository(db)
	departmentUseCase := departmentUseCase.NewUseCase(departmentRepository)
	departmentHandler := NewDepartmentHandler(departmentUseCase)

	departmentRouter := v1.Group("department")
	departmentRouter.POST("/", departmentHandler.Post)
	departmentRouter.PATCH("/:id", departmentHandler.Update)
}
