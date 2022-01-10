package services

import (
	"github.com/stefanusong/posly-backend/dto"
	"github.com/stefanusong/posly-backend/entities"
	"github.com/stefanusong/posly-backend/repositories"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateResto(resto dto.RegisterDTO) entities.Resto
	IsEmailExists(email string) bool
	IsSlugExists(slug string) bool

}

type authService struct {
	restoRepository repositories.IRestoRepository
}

func NewAuthService(restoRepo repositories.IRestoRepository) IAuthService {
	return &authService{
		restoRepository: restoRepo,
	}
}

func (s *authService) VerifyCredential(email string, password string) interface{} {
	if email == "" || password == "" {
		return false
	}

	res := s.restoRepository.FindRestoByEmail(email)
	if res != (entities.Resto{}) {
		err := bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password))
		if err != nil {
			return false
		}
		return res
	}
	return false
}

func (s *authService) CreateResto(resto dto.RegisterDTO) entities.Resto {
	newResto := entities.Resto{
		OwnerName:    resto.OwnerName,
		RestoName:    resto.RestoName,
		RestoSlug:    resto.RestoSlug,
		RestoAddress: resto.RestoAddress,
		Phone:        resto.Phone,
		Email:        resto.Email,
		Password:     resto.Password,
	}

	res := s.restoRepository.CreateResto(newResto)
	return res
}

func (s *authService) IsEmailExists(email string) bool {
	res := s.restoRepository.FindRestoByEmail(email)
	return res != (entities.Resto{})
}

func (s *authService) IsSlugExists(slug string) bool {
	res := s.restoRepository.FindRestoBySlug(slug)
	return res != (entities.Resto{})
}