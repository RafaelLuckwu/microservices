package main

import (
    "log"
    "github.com/ruandg/microservices/order/config"
    db "github.com/ruandg/microservices/order/internal/adapters/db"
    grpcAdapter "github.com/ruandg/microservices/order/internal/adapters/grpc"
    "github.com/ruandg/microservices/order/internal/application/core/api"
)

func main() {
    dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
    if err != nil {
        log.Fatalf("Failed to connect to database. Error: %v", err)
    }
    application := api.NewApplication(dbAdapter)
    grpcSrv := grpcAdapter.NewAdapter(application, config.GetApplicationPort())
    grpcSrv.Run()
}
