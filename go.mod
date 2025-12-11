module microservice-order

go 1.22

require (
  github.com/SEU_USUARIO/microservices-proto/golang/payment v0.0.0-00010101000000-000000000000
  google.golang.org/grpc v1.63.0
)

replace github.com/SEU_USUARIO/microservices-proto/golang/payment => ../microservices-proto/golang/payment
