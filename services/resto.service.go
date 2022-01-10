package services

import (
	"github.com/stefanusong/posly-backend/dto"
	"github.com/stefanusong/posly-backend/entities"
	"github.com/stefanusong/posly-backend/repositories"
)

type IRestoService interface {
	UpdateResto(resto dto.RestoUpdateDTO) entities.Resto
	GetRestoProfile(restoID uint64) entities.Resto
	IsEmailExists(email string) bool
	IsSlugExists(slug string) bool
}

type restoService struct {
	restoRepository repositories.IRestoRepository
}

func NewRestoService(restoRepo repositories.IRestoRepository) IRestoService {
	return &restoService{
		restoRepository: restoRepo,
	}
}

func (s *restoService) UpdateResto(resto dto.RestoUpdateDTO) entities.Resto {
	
	restoToUpdate := entities.Resto{
		ID:        resto.ID,
		OwnerName: resto.OwnerName,
		RestoName: resto.RestoName,
		RestoSlug: resto.RestoSlug,
		RestoAddress: resto.RestoAddress,
		Phone: resto.Phone,
		Email:     resto.Email,
		Password:  resto.Password,
	}

	updatedResto := s.restoRepository.UpdateResto(restoToUpdate)
	return updatedResto
}

func (s *restoService) GetRestoProfile(restoID uint64) entities.Resto {
	return s.restoRepository.FindRestoByID(restoID)
}


func (s *restoService) IsEmailExists(email string) bool {
	res := s.restoRepository.FindRestoByEmail(email)
	return res != (entities.Resto{})
}

func (s *restoService) IsSlugExists(slug string) bool {
	res := s.restoRepository.FindRestoBySlug(slug)
	return res != (entities.Resto{})
}