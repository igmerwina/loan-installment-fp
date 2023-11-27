package model

import "time"

type Installment struct {
	ID        int64     `json:"id" gorm:"primaryKey"`
	UserID    int64     `json:"user_id"`
	LoanID    int64     `json:"loan_id"`
	Nominal   int64     `json:"nominal"`
	PaymentAt time.Time `json:"payment_at"`
}
