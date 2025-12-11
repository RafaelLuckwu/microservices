package ports

import "microservice-order/internal/application/core/domain"

type PaymentPort interface {
	Charge(*domain.Order) error
}
