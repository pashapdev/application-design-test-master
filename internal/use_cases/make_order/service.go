package makeorder

import (
	"fmt"
	"time"

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

func (s *Service) Do(request *request) error {
	order, err := s.orderFromRequestItem(request)
	if err != nil {
		return err
	}

	s.repo.InsertOrders(order)
	return nil
}

func (s *Service) orderFromRequestItem(request *request) (*entities.Order, error) {
	const timeFormat = "2006-01-02"
	orderFrom, err := time.Parse(timeFormat, request.From)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", request.From, errors.ErrDateValidation)
	}

	orderTo, err := time.Parse(timeFormat, request.To)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", request.To, errors.ErrDateValidation)
	}

	email := entities.Email(request.UserEmail)
	if err := email.Valid(); err != nil {
		return nil, fmt.Errorf("%s: %w", email, errors.ErrEmailValidation)
	}

	if _, exists := entities.ValidRoomTypes[request.Room]; !exists {
		return nil, fmt.Errorf("%s: %w", request.Room, errors.ErrRoomValidation)
	}

	room := entities.RoomType(request.Room)

	return &entities.Order{
		Room:      room,
		UserEmail: email,
		From:      orderFrom,
		To:        orderTo,
	}, nil
}
