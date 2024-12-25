package payment

import "fmt"

type Stripe struct {
}

func (stripe *Stripe) Pay(amount int) {
	fmt.Printf("Stripe's payment was successfully done. Amount: %d cents\n", amount)
}
