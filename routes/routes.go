package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/stefanusong/posly-backend/configs"
	"github.com/stefanusong/posly-backend/controllers"
	"github.com/stefanusong/posly-backend/middlewares"
	"github.com/stefanusong/posly-backend/repositories"
	"github.com/stefanusong/posly-backend/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                      = configs.ConnectDB()
	restoRepository repositories.IRestoRepository = repositories.NewRestoRepository(db)
	authService    services.IAuthService         = services.NewAuthService(restoRepository)
	jwtService     services.IJwtService          = services.NewJWTService()
	restoService    services.IRestoService        = services.NewRestoService(restoRepository)
	authController controllers.IAuthController   = controllers.NewAuthController(authService, jwtService)
	restoController controllers.IRestoController  = controllers.NewRestoController(restoService, jwtService)
)

func Init() (*echo.Echo, *gorm.DB) {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Masuk pak echo !")
	})

	// auth routes
	authRoute := e.Group("/api/auth")
	authRoute.POST("/login", authController.Login)
	authRoute.POST("/register", authController.Register)

	// resto routes
	restoRoute := e.Group("/api/resto")
	restoRoute.Use(middlewares.IsAuthenticated)
	restoRoute.GET("/profile", restoController.GetRestoProfile)
	restoRoute.PUT("/profile", restoController.UpdateResto)

	return e, db
}
