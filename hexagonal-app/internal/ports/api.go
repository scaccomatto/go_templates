package ports

import "paru.net/hexagonalapp/internal/application/domain"

type APIPort interface {
	PlaceOrder(order domain.Order) (domain.Order, error)
}
