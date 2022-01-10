package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/stefanusong/posly-backend/dto"
	"github.com/stefanusong/posly-backend/entities"
	"github.com/stefanusong/posly-backend/helpers"
	"github.com/stefanusong/posly-backend/services"
)

type IAuthController interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type authController struct {
	authService services.IAuthService
	jwtService  services.IJwtService
}

func NewAuthController(authService services.IAuthService, jwtService services.IJwtService) IAuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx echo.Context) error {

	// Bind req body to LoginDTO
	var loginDTO dto.LoginDTO
	if errDto := ctx.Bind(&loginDTO); errDto != nil {
		response := helpers.CreateErrorResponse("Invalid request body", errDto.Error(), helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Validate input fields
	if errValidate := helpers.ValidateBody(loginDTO); errValidate != "" {
		response := helpers.CreateErrorResponse("Invalid request body", errValidate, helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Verify credential
	authResult := c.authService.VerifyCredential(loginDTO.Email, loginDTO.Password)

	// Valid Credential, return resto with token
	if v, ok := authResult.(entities.Resto); ok {
		token := c.jwtService.GenerateToken(strconv.FormatUint(v.ID, 10))
		v.Token = token
		response := helpers.CreateResponse(true, "Logged in", v)
		return ctx.JSON(http.StatusOK, response)
	}

	//  Invalid Credential
	response := helpers.CreateErrorResponse("Invalid credential", "Invalid email or password", helpers.EmptyObj{})
	return ctx.JSON(http.StatusUnauthorized, response)
}

func (c *authController) Register(ctx echo.Context) error {

	// Bind req body to register DTO
	var registerDTO dto.RegisterDTO
	if errDto := ctx.Bind(&registerDTO); errDto != nil {
		response := helpers.CreateErrorResponse("Invalid request body", errDto.Error(), helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Validate input fields
	if errValidate := helpers.ValidateBody(registerDTO); errValidate != "" {
		response := helpers.CreateErrorResponse("Invalid request body", errValidate, helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Check if email already exists in the database
	if c.authService.IsEmailExists(registerDTO.Email) {
		response := helpers.CreateErrorResponse("Please use another email", "Email already exists", helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Check if slug already exists in the database
	if c.authService.IsSlugExists(registerDTO.RestoSlug) {
		response := helpers.CreateErrorResponse("Please use another slug", "Slug already exists", helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Create new resto
	registerResult := c.authService.CreateResto(registerDTO)

	if registerResult != (entities.Resto{}) {
		token := c.jwtService.GenerateToken(strconv.FormatUint(registerResult.ID, 10))
		registerResult.Token = token
		response := helpers.CreateResponse(true, "Resto Registered", registerResult)
		return ctx.JSON(http.StatusCreated, response)
	}

	//  Failed to create resto
	response := helpers.CreateErrorResponse("Failed to create new resto", "Internal Server Error", helpers.EmptyObj{})
	return ctx.JSON(http.StatusUnauthorized, response)
}
