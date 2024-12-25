package main

import (
	"main/payment"
)

func main() {

	stripe := payment.Stripe{}
	paypal := payment.Paypal{}

	gateways := []payment.UnifiedPaymentGateway{
		&payment.StripeAdapter{
			Adaptee: &stripe,
		},
		payment.PaypalAdapter{
			Adaptee: &paypal,
		},
	}

	for _, gateway := range gateways {
		gateway.Pay(100)
	}
}
