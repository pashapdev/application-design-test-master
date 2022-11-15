package getorders

import (
	"fmt"

	"github.com/pashapdev/applicationDesignTest/internal/entities"
	"github.com/pashapdev/applicationDesignTest/internal/errors"
)

type Service struct {
	repo repo
}

func NewSvc(repo repo) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) Do(request *request) (*response, error) {
	email := entities.Email(request.Email)
	if err := email.Valid(); err != nil {
		return nil, fmt.Errorf("%s: %w", request.Email, errors.ErrEmailValidation)
	}

	orders := s.repo.SelectOrdersByEmail(email)
	response := new(response)
	response.Orders = make([]Order, len(orders))
	for i := range orders {
		response.Orders[i] = NewOrder(&orders[i])
	}

	return response, nil
}
