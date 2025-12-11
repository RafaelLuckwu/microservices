package grpc

import (
	"log"
	"net"
	"microservice-order/internal/application/core/api"

	pb "github.com/SEU_USUARIO/microservices-proto/golang/order"
	"google.golang.org/grpc"
)

type Adapter struct {
	app *api.Application
	port string
	pb.UnimplementedOrderServiceServer
}

func NewAdapter(app *api.Application, port string) *Adapter {
	return &Adapter{app: app, port: port}
}

func (a *Adapter) Run() {
	lis, err := net.Listen("tcp", a.port)
	if err != nil { log.Fatalf("Error: %v", err) }

	server := grpc.NewServer()
	pb.RegisterOrderServiceServer(server, a)
	log.Printf("Order running %s", a.port)
	server.Serve(lis)
}
