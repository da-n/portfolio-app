package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type UserService interface {
	GetUser(string) (*dto.UserResponse, *errs.AppError)
}

type DefaultUserService struct {
	Repo domain.UserRepository
}

func (service DefaultUserService) GetUser(id string) (*dto.UserResponse, *errs.AppError) {
	return nil, nil
}

func NewUserService(r domain.UserRepository) DefaultUserService {
	return DefaultUserService{r}
}
