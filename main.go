package main

import (
	"log"
	"microservice-order/config"
	db "microservice-order/internal/adapters/db"
	pay "microservice-order/internal/adapters/payment"
	grpcServer "microservice-order/internal/adapters/grpc"
	"microservice-order/internal/application/core/api"
)

func main() {
	dbAdapter, _ := db.NewAdapter(config.GetDataSourceURL())
	paymentAdapter, err := pay.NewAdapter(config.GetPaymentServiceUrl())
	if err != nil { log.Fatalf("Payment error: %v", err) }

	app := api.NewApplication(dbAdapter, paymentAdapter)
	server := grpcServer.NewAdapter(app, config.GetApplicationPort())
	server.Run()
}
