package payment

import "fmt"

type Paypal struct {
}

func (paypal *Paypal) Pay(amount int) {
	fmt.Printf("Paypal's payment was successfully done. Amount: %d dollars\n", amount)
}
