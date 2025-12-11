package db

import (
	"microservice-payment/internal/domain"
	"microservice-payment/internal/ports"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Adapter struct { db *gorm.DB }

var _ ports.DBPort = (*Adapter)(nil)

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil { return nil, err }
	db.AutoMigrate(&domain.Payment{})
	return &Adapter{db: db}, nil
}

func (a *Adapter) Save(payment *domain.Payment) error {
	return a.db.Create(payment).Error
}
