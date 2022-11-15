package getorders

import "github.com/pashapdev/applicationDesignTest/internal/entities"

type repo interface {
	SelectOrdersByEmail(email entities.Email) []entities.Order
}
