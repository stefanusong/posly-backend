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

type IRestoController interface {
	UpdateResto(c echo.Context) error
	GetRestoProfile(c echo.Context) error
}

type restoController struct {
	restoService services.IRestoService
	jwtService   services.IJwtService
}

func NewRestoController(restoService services.IRestoService, jwtService services.IJwtService) IRestoController {
	return &restoController{
		restoService: restoService,
		jwtService:   jwtService,
	}
}

func (c *restoController) UpdateResto(ctx echo.Context) error {
	var RestoUpdateDTO = dto.RestoUpdateDTO{}

	// Bind request body to restoUpdateDto
	if errDTO := ctx.Bind(&RestoUpdateDTO); errDTO != nil {
		response := helpers.CreateErrorResponse("Invalid request body", errDTO.Error(), helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Validate request body
	if errValidate := helpers.ValidateBody(RestoUpdateDTO); errValidate != "" {
		response := helpers.CreateErrorResponse("Invalid request body", errValidate, helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Get id from token claims
	claims := c.jwtService.GetUserClaims(ctx)
	id, err := strconv.ParseUint(claims.UserID, 10, 64)
	helpers.HandleError(err)
	RestoUpdateDTO.ID = id

	// Get current user data
	currentResto := c.restoService.GetRestoProfile(id)

	// Check if email already exists in the database
	if RestoUpdateDTO.Email != "" && RestoUpdateDTO.Email != currentResto.Email && c.restoService.IsEmailExists(RestoUpdateDTO.Email) {
		response := helpers.CreateErrorResponse("Please use another email", "Email already exists", helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Check if slug already exists in the database
	if RestoUpdateDTO.RestoSlug != "" && RestoUpdateDTO.RestoSlug != currentResto.RestoSlug && c.restoService.IsSlugExists(RestoUpdateDTO.RestoSlug) {
		response := helpers.CreateErrorResponse("Please use another slug", "Slug already exists", helpers.EmptyObj{})
		return ctx.JSON(http.StatusBadRequest, response)
	}

	// Update resto
	updatedResto := c.restoService.UpdateResto(RestoUpdateDTO)
	res := helpers.CreateResponse(true, "Resto updated", updatedResto)
	return ctx.JSON(http.StatusOK, res)
}

func (c *restoController) GetRestoProfile(ctx echo.Context) error {
	// Get id from token claims
	claims := c.jwtService.GetUserClaims(ctx)
	id, err := strconv.ParseUint(claims.UserID, 10, 64)
	helpers.HandleError(err)

	resto := c.restoService.GetRestoProfile(id)
	if resto == (entities.Resto{}) {
		response := helpers.CreateErrorResponse("Resto not found", "Resto not found", helpers.EmptyObj{})
		return ctx.JSON(http.StatusNotFound, response)
	}
	return ctx.JSON(http.StatusOK, helpers.CreateResponse(true, "Resto retrieved", resto))
}
