package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/stefanusong/posly-backend/helpers"
	"github.com/stefanusong/posly-backend/services"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(helpers.GetJWTSecret()),
	Claims:     services.GetCustomClaim(),
})
