package handlers

import (
	"github.com/fatjan/gogomanager/internal/config"
	"github.com/fatjan/gogomanager/internal/pkg/jwt_helper"
	authRepository "github.com/fatjan/gogomanager/internal/repositories/auth"
	departmentRepository "github.com/fatjan/gogomanager/internal/repositories/department"
	duckRepository "github.com/fatjan/gogomanager/internal/repositories/duck"
	employeeRepository "github.com/fatjan/gogomanager/internal/repositories/employee"
	authUseCase "github.com/fatjan/gogomanager/internal/useCases/auth"
	departmentUseCase "github.com/fatjan/gogomanager/internal/useCases/department"
	duckUseCase "github.com/fatjan/gogomanager/internal/useCases/duck"
	employeeUseCase "github.com/fatjan/gogomanager/internal/useCases/employee"

	userRepository "github.com/fatjan/gogomanager/internal/repositories/user"
	userUseCase "github.com/fatjan/gogomanager/internal/useCases/user"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(cfgData *config.Config, db *sqlx.DB, r *gin.Engine) {
	// integrasi jwt
	jwtMiddleware := jwt_helper.JWTMiddleware(cfgData.JwtKey)

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
	departmentRouter.Use(jwtMiddleware)
	departmentRouter.POST("/", departmentHandler.Post)
	departmentRouter.GET("/", departmentHandler.Index)
	departmentRouter.PATCH("/:id", departmentHandler.Update)
	departmentRouter.DELETE("/:id", departmentHandler.Delete)

	employeeRepository := employeeRepository.NewEmployeeRepository(db)
	employeeUseCase := employeeUseCase.NewUseCase(employeeRepository)
	employeeHandler := NewEmployeeHandler(employeeUseCase)

	employeeRouter := v1.Group("employee")
	employeeRouter.Use(jwtMiddleware)
	employeeRouter.GET("/", employeeHandler.Get)
	employeeRouter.POST("/", employeeHandler.Post)

	authRepository := authRepository.NewAuthRepository(db)
	authUseCase := authUseCase.NewUseCase(authRepository, cfgData)
	authHandler := NewAuthHandler(authUseCase)

	authRouter := v1.Group("auth")
	authRouter.POST("/", authHandler.Post)

	userRepository := userRepository.NewUserRepository(db)
	userUseCase := userUseCase.NewUseCase(userRepository)
	userHandler := NewUserHandler(userUseCase)

	userRouter := v1.Group("user")
	userRouter.Use(jwtMiddleware)
	userRouter.GET("/", userHandler.Get)
}
