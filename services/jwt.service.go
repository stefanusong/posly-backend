package services

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stefanusong/posly-backend/helpers"
)

type IJwtService interface {
	GenerateToken(userID string) string
	GetUserClaims(ctx echo.Context) *jwtCustomClaims
}

type jwtService struct {
	issuer    string
	secretKey string
}

// Custom claim struct
type jwtCustomClaims struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

// Create new service instance
func NewJWTService() IJwtService {
	return &jwtService{
		issuer:    "stefanusong",
		secretKey: helpers.GetJWTSecret(),
	}
}

//Generate JWT Token with custom claims
func (j *jwtService) GenerateToken(userID string) string {
	// Set custom claims
	claims := &jwtCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(j.secretKey))

	helpers.HandleError(err)

	return t
}

func GetCustomClaim() *jwtCustomClaims {
	return &jwtCustomClaims{}
}

func (j *jwtService) GetUserClaims(ctx echo.Context) *jwtCustomClaims {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	return claims
}
