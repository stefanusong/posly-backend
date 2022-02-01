package authcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ports "github.com/stefanusong/posly-backend/internal/ports/services"
)

type httpController struct {
	authService ports.AuthService
}

func NewHttpController(authService ports.AuthService) *httpController {
	return &httpController{
		authService: authService,
	}
}

func (httpController *httpController) Login(c *gin.Context) {
	httpController.authService.Login()
	c.JSON(http.StatusAccepted, gin.H{
		"data": "Login",
	})
}
