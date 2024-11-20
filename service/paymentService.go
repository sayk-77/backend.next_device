package service

import (
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"
	"next_device/backend/models"
	"next_device/backend/repository"
	"os"
	"strconv"
)

type PaymentService struct {
	paymentRepo *repository.PaymentRepository
}

func NewPaymentService(paymentRepo *repository.PaymentRepository) *PaymentService {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	return &PaymentService{paymentRepo: paymentRepo}
}

func (s *PaymentService) CreateCheckoutSession(orderID uint, totalPrice float64) (*stripe.CheckoutSession, error) {
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
					Currency: stripe.String("rub"),
					ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
						Name: stripe.String("Order #" + strconv.Itoa(int(orderID))),
					},
					UnitAmount: stripe.Int64(int64(totalPrice * 100)),
				},
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String("payment"),
		SuccessURL: stripe.String("https://localhost:3000/success?order_id=" + strconv.Itoa(int(orderID)) + "&session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String("https://localhost:3000/error?order_id=" + strconv.Itoa(int(orderID)) + "&session_id={CHECKOUT_SESSION_ID}"),
		Metadata: map[string]string{
			"order_id": strconv.Itoa(int(orderID)),
		},
	}

	session, err := session.New(params)
	if err != nil {
		return nil, err
	}

	return session, nil
}

func (s *PaymentService) SavePayment(payment *models.Payment) error {
	err := s.paymentRepo.SavePayment(payment)
	if err != nil {
		return err
	}
	return nil
}
