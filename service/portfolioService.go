package service

import (
	"github.com/da-n/portfolio-app/domain"
	"github.com/da-n/portfolio-app/dto"
	"github.com/da-n/portfolio-app/errs"
)

type PortfolioService interface {
	GetPortfolio(int64) (*dto.PortfolioResponse, *errs.AppError)
}

type DefaultPortfolioService struct {
	repo domain.PortfolioRepository
}

func (service DefaultPortfolioService) GetPortfolio(portfolioId int64) (*dto.PortfolioResponse, *errs.AppError) {
	p, err := service.repo.FindPortfolioById(portfolioId)
	if err != nil {
		return nil, err
	}
	response := p.ToDto()
	return &response, nil
}

func NewPortfolioService(r domain.PortfolioRepository) DefaultPortfolioService {
	return DefaultPortfolioService{r}
}
