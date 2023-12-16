package domain

import "context"

type MidtransService interface {
	GenerateSnapURL(c context.Context, topUp *TopUp) error
	VerifyPayment(c context.Context, orderId string) (bool, error)
}
