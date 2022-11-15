package getorders

import "github.com/pashapdev/applicationDesignTest/internal/entities"

type request struct {
	Email string `json:"email"`
}

type response struct {
	Orders []Order `json:"orders"`
}

type Order struct {
	Room      string
	UserEmail string
	From      string
	To        string
}

func NewOrder(order *entities.Order) Order {
	const timeFormat = "2006-01-02"
	return Order{
		Room:      string(order.Room),
		UserEmail: string(order.UserEmail),
		From:      order.From.Format(timeFormat),
		To:        order.To.Format(timeFormat),
	}
}
