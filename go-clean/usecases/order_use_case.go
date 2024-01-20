package usecases

import (
	"errors"

	"github.com/kitpk/go-architecture105/entities"
)

type OrderUseCase interface {
	CreateOrder(order entities.Order) error
}

type OrderService struct {
	repo OrderRepository
}

func NewOrderService(repo OrderRepository) OrderUseCase {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(order entities.Order) error {
	if order.Total <= 0 {
		return errors.New("Total must be positive")
	}

	if err := s.repo.Save(order); err != nil {
		return err
	}

	return nil
}
