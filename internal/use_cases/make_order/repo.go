package makeorder

import "github.com/pashapdev/applicationDesignTest/internal/entities"

type repo interface {
	InsertOrders(orders *entities.Order)
}
