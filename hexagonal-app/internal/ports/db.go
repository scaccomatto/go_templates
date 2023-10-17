package ports

import "paru.net/hexagonalapp/internal/application/domain"

type DBPort interface {
	Get(id string) (domain.Order, error)
	Save(*domain.Order) error
}
