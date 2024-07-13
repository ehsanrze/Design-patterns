package main

import (
	"main/database"
	"sync"
)

type PaymentSingleton struct {
}

var lock = &sync.Mutex{}

var paymentInstance *PaymentSingleton

func GetPaymentInstance() *PaymentSingleton {
	lock.Lock()
	defer lock.Unlock()

	if paymentInstance == nil {
		paymentInstance = &PaymentSingleton{}
	}

	return paymentInstance
}

func (payment PaymentSingleton) Pay(newPayment database.Payment) {

	err := database.CreatePayment(newPayment)

	if err != nil {
		panic(err)
	}

}
