package service

import (
	"microservice-payment/internal/domain"
	"microservice-payment/internal/ports"
)

type PaymentService struct { db ports.DBPort }

func NewPaymentService(db ports.DBPort) *PaymentService {
	return &PaymentService{db: db}
}

func (s *PaymentService) Create(uid, oid int64, total float32) (*domain.Payment, error) {
	p := &domain.Payment{
		UserID: uid, OrderID: oid, Total: total,
		BillID: oid + 1000,
	}
	err := s.db.Save(p)
	return p, err
}
