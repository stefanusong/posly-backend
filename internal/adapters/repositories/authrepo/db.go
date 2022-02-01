package authrepo

import (
	"fmt"

	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repo {
	return &repo{
		db: db,
	}
}

func (repo *repo) CreateUser() {
	fmt.Println("Create User")
}
