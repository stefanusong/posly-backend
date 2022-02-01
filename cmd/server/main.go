package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stefanusong/posly-backend/configs"
	"github.com/stefanusong/posly-backend/internal/adapters/controllers/authcontroller"
	"github.com/stefanusong/posly-backend/internal/adapters/repositories/authrepo"
	"github.com/stefanusong/posly-backend/internal/core/services/authservice"
)

func main() {
	db := configs.OpenDBConnection()

	// Auth
	authRepo := authrepo.NewRepo(db)
	authService := authservice.NewService(authRepo)
	authController := authcontroller.NewHttpController(authService)

	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.GET("/login", authController.Login)
		}
	}

	r.Run(":8080")
}
