package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
)

type UserService interface {
	GetUser(string) (*dto.UserResponse, *error)
}

type DefaultUserService struct {
	Repo domain.UserRepository
}

func (service DefaultUserService) GetUser(id string) (*dto.UserResponse, *error) {
	return nil, nil
}

func NewUserService(repo domain.UserRepository) DefaultUserService {
	return DefaultUserService{repo}
}
