package repositories

import (
	"github.com/stefanusong/posly-backend/entities"
	"github.com/stefanusong/posly-backend/helpers"
	"gorm.io/gorm"
)

//Resto repository will be communicating with database
type IRestoRepository interface {
	CreateResto(resto entities.Resto) entities.Resto
	UpdateResto(resto entities.Resto) entities.Resto
	FindRestoByEmail(email string) entities.Resto
	FindRestoBySlug(slug string) entities.Resto
	FindRestoByID(restoID uint64) entities.Resto
}

type restoConnection struct {
	connection *gorm.DB
}

func NewRestoRepository(db *gorm.DB) IRestoRepository {
	return &restoConnection{
		connection: db,
	}
}

func (db *restoConnection) CreateResto(resto entities.Resto) entities.Resto {
	resto.Password = helpers.HashAndSalt([]byte(resto.Password))
	db.connection.Save(&resto)
	return resto
}

func (db *restoConnection) UpdateResto(resto entities.Resto) entities.Resto {
	var tempResto entities.Resto
	db.connection.Find(&tempResto, resto.ID)

	if resto.Password != "" {
		resto.Password = helpers.HashAndSalt([]byte(resto.Password))
	} else {
		resto.Password = tempResto.Password
	}

	db.connection.Model(&resto).Updates(resto)
	return resto
}

func (db *restoConnection) FindRestoByEmail(email string) entities.Resto {
	var resto entities.Resto
	db.connection.Where("email = ?", email).Take(&resto)
	return resto
}

func (db *restoConnection) FindRestoBySlug(slug string) entities.Resto {
	var resto entities.Resto
	db.connection.Where("resto_slug = ?", slug).Take(&resto)
	return resto
}

func (db *restoConnection) FindRestoByID(restoID uint64) entities.Resto {
	var resto entities.Resto
	db.connection.Find(&resto, restoID)
	return resto
}
