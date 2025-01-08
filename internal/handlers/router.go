package handlers

import (
	"github.com/fatjan/gogomanager/internal/config"
	authRepository "github.com/fatjan/gogomanager/internal/repositories/auth"
	departmentRepository "github.com/fatjan/gogomanager/internal/repositories/department"
	duckRepository "github.com/fatjan/gogomanager/internal/repositories/duck"
	employeeRepository "github.com/fatjan/gogomanager/internal/repositories/employee"
	authUseCase "github.com/fatjan/gogomanager/internal/useCases/auth"
	departmentUseCase "github.com/fatjan/gogomanager/internal/useCases/department"
	duckUseCase "github.com/fatjan/gogomanager/internal/useCases/duck"
	employeeUseCase "github.com/fatjan/gogomanager/internal/useCases/employee"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(cfgData *config.Config, db *sqlx.DB, r *gin.Engine) {
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

	employeeRepository := employeeRepository.NewEmployeeRepository(db)
	employeeUseCase := employeeUseCase.NewUseCase(employeeRepository)
	employeeHandler := NewEmployeeHandler(employeeUseCase)

	employeeRouter := v1.Group("employee")
	employeeRouter.GET("/", employeeHandler.Get)

	authRepository := authRepository.NewAuthRepository(db)
	authUseCase := authUseCase.NewUseCase(authRepository, cfgData)
	authHandler := NewAuthHandler(authUseCase)

	authRouter := v1.Group("auth")
	authRouter.POST("/", authHandler.Post)
}
