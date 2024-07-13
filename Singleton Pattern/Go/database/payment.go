package database

import (
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Payment struct {
	ID            uint `gorm:"primaryKey"`
	UserId        uint `gorm:"index"`
	Amount        int
	Currency      string
	Status        string
	PaymentMethod string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func CreatePayment(newPayment Payment) error {

	tx := database.Begin()

	if tx.Error != nil {
		return fmt.Errorf("could not begin transaction: %v", tx.Error)
	}

	err := tx.Model(&User{}).Where("id = ?", newPayment.UserId).UpdateColumn("credit", gorm.Expr("credit + ?", newPayment.Amount)).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update credit: %v", err)
	}

	err = tx.Create(&newPayment).Error
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to create payment record: %v", err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("transaction commit failed: %v", err)
	}

	return nil
}
