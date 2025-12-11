package ports

import "microservice-order/internal/application/core/domain"

type DBPort interface {
	Save(*domain.Order) error
}
