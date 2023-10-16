package ports

import "paru.net/exagonalapp/internal/application/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
