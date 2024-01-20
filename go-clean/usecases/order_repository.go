package usecases

import "github.com/kitpk/go-architecture105/entities"

type OrderRepository interface {
	Save(order entities.Order) error
}
