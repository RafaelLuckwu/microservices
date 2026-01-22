package service

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

func CallPayment(conn *grpc.ClientConn) error {
    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    // Simulação de chamada gRPC
    err := status.Error(codes.DeadlineExceeded, "timeout simulado")

    if err != nil {
        st, ok := status.FromError(err)
        if ok && st.Code() == codes.DeadlineExceeded {
            log.Println("Timeout ao chamar o serviço Payment (DeadlineExceeded)")
        } else {
            log.Printf("Erro ao chamar Payment: %v", err)
        }
        return err
    }

    return nil
}
