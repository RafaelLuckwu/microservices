package api

import (
	"context"
	"microservice-order/internal/application/core/domain"
	"microservice-order/internal/ports"
)

type Application struct {
	db      ports.DBPort
	payment ports.PaymentPort
}

func NewApplication(db ports.DBPort, pay ports.PaymentPort) *Application {
	return &Application{db: db, payment: pay}
}

func (a Application) PlaceOrder(o domain.Order) (domain.Order, error) {
	if err := a.db.Save(&o); err != nil {
		return domain.Order{}, err
	}
	if err := a.payment.Charge(&o); err != nil {
		return domain.Order{}, err
	}
	return o, nil
}
func (a Application) Charge(
	ctx context.Context,
	payment domain.Payment,
) (domain.Payment, error) {

	if payment.TotalPrice > 1000 {
		return domain.Payment{}, status.Errorf(
			codes.InvalidArgument,
			"payment over 1000 is not allowed",
		)
	}
	paymentResp, err := a.paymentClient.Create(ctx, paymentReq)
if err != nil {
	order.Status = domain.OrderCanceled
	_ = a.db.Update(ctx, &order)
	return domain.Order{}, err
}

order.Status = domain.OrderPaid
err = a.db.Update(ctx, &order)
if err != nil {
	return domain.Order{}, err
}


	

	result, err := a.api.Charge(ctx, newPayment)

	code := status.Code(err)
	if code == codes.InvalidArgument {
		return nil, err
	} else if err != nil {
		return nil, status.New(
			codes.Internal,
			fmt.Sprintf("failed to charge: %v", err),
		).Err()
	}

	return &payment.CreatePaymentResponse{
		PaymentId: result.ID,
	}, nil
}
result, err := a.api.CreateOrder(ctx, newOrder)

code := status.Code(err)
if code == codes.InvalidArgument {
	return nil, err
} else if err != nil {
	return nil, status.New(
		codes.Internal,
		fmt.Sprintf("failed to create order: %v", err),
	).Err()
}

	result, err := a.api.Charge(ctx, newPayment)
	if err != nil {
		return domain.Payment{}, err
	}

	err = a.db.Save(ctx, &payment)
	if err != nil {
		return domain.Payment{}, err
	}

	return payment, nil
}
