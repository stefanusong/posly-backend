package authservice

import (
	ports "github.com/stefanusong/posly-backend/internal/ports/repositories"
)

type service struct {
	authRepo ports.AuthRepo
}

func NewService(authRepo ports.AuthRepo) *service {
	return &service{
		authRepo: authRepo,
	}
}

func (srv *service) Login() {
	srv.authRepo.CreateUser()
}
