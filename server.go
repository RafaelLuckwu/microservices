package grpc

import (
	"context"
	"log"
	"net"

	pb "github.com/SEU_USUARIO/microservices-proto/golang/payment"
	"microservice-payment/internal/service"
	"google.golang.org/grpc"
)

type Adapter struct {
	port string
	svc  *service.PaymentService
	pb.UnimplementedPaymentServer
}

func NewAdapter(port string, svc *service.PaymentService) *Adapter {
	return &Adapter{port: port, svc: svc}
}

func (a *Adapter) Create(ctx context.Context, req *pb.CreatePaymentRequest) (*pb.CreatePaymentResponse, error) {
	p, err := a.svc.Create(req.UserId, req.OrderId, req.TotalPrice)
	if err != nil { return nil, err }

	return &pb.CreatePaymentResponse{
		PaymentId: p.ID,
		BillId:    p.BillID,
	}, nil
}

func (a *Adapter) Run() {
	lis, err := net.Listen("tcp", a.port)
	if err != nil { log.Fatalf("Listen error: %v", err) }

	server := grpc.NewServer()
	pb.RegisterPaymentServer(server, a)

	log.Printf("Payment service running on %s", a.port)
	server.Serve(lis)
}
