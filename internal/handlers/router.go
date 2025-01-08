package handlers

import (
	"github.com/fatjan/gogomanager/internal/config"
	departmentRepository "github.com/fatjan/gogomanager/internal/repositories/department"
	duckRepository "github.com/fatjan/gogomanager/internal/repositories/duck"
	userRepository "github.com/fatjan/gogomanager/internal/repositories/user"
	departmentUseCase "github.com/fatjan/gogomanager/internal/useCases/department"
	duckUseCase "github.com/fatjan/gogomanager/internal/useCases/duck"
	userUseCase "github.com/fatjan/gogomanager/internal/useCases/user"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(_ *config.Config, db *sqlx.DB, r *gin.Engine) {
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
	departmentRouter.DELETE("/:id", departmentHandler.Delete)

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUseCase(userRepository)
	userHandler := NewUserHandler(userUseCase)

	userRouter := v1.Group("user")
	userRouter.GET("/", userHandler.Get)
}
