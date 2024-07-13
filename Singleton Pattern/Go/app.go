package main

import (
	"main/database"
	"math/rand"
	"time"
)

func main() {
	database.ConnectToDatabase()

	for i := 0; i < 100; i++ {

		paymentMethod := ChooseFromSlice([]string{"deposit", "withdrawal"})

		paymentOperator := -1
		if paymentMethod == "deposit" {
			paymentOperator = 1
		}

		go GetPaymentInstance().Pay(database.Payment{
			Amount:        rand.Intn(10000) * paymentOperator,
			UserId:        1,
			Currency:      "toman",
			Status:        "success",
			PaymentMethod: paymentMethod,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		})
	}
}

func ChooseFromSlice(options []string) string {
	return options[rand.Intn(len(options))]
}
