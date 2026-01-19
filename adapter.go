package adapter
import (
	"context"
	"microservice-order/internal/application/core/domain"
	"microservice-order/internal/ports/payment"
)
func (a Adapter) Create(
	ctx context.Context,
	request *payment.CreatePaymentRequest,
) (*payment.CreatePaymentResponse, error) {

	log.WithContext(ctx).Info("Creating payment...")
	newPayment := domain.NewPayment(
		request.UserId,
		request.OrderId,
		request.TotalPrice,
	)