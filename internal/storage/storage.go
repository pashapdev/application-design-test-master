package storage

import (
	"sync"

	"github.com/pashapdev/applicationDesignTest/internal/entities"
)

type Storage struct {
	orders []entities.Order
	mx     *sync.RWMutex
}

func New() *Storage {
	return &Storage{
		mx: &sync.RWMutex{},
	}
}

func (s *Storage) InsertOrders(order *entities.Order) {
	s.mx.Lock()
	defer s.mx.Unlock()

	s.orders = append(s.orders, *order)
}

func (s *Storage) SelectOrdersByEmail(email entities.Email) []entities.Order {
	s.mx.Lock()
	defer s.mx.Unlock()

	var orders []entities.Order
	for i := range s.orders {
		if s.orders[i].UserEmail == email {
			orders = append(orders, s.orders[i])
		}
	}

	return orders
}
