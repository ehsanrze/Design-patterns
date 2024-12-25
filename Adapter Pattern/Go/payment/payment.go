package payment

type UnifiedPaymentGateway interface {
	Pay(amount int)
}

type StripeAdapter struct {
	Adaptee *Stripe
}

func (s *StripeAdapter) Pay(amount int) {
	s.Adaptee.Pay(amount * 10)
}

type PaypalAdapter struct {
	Adaptee *Paypal
	Test    int
}

func (p PaypalAdapter) Pay(amount int) {
	p.Test = 10
	p.Adaptee.Pay(amount)
}
