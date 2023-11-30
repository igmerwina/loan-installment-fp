package model

import "time"

type Installment struct {
	ID        int64     `json:"id" gorm:"primaryKey" swagger:"description(ID)"`
	UserID    int64     `json:"user_id" swagger:"description(UserID)"`
	LoanID    int64     `json:"loan_id" swagger:"description(LoanID)"`
	Nominal   int64     `json:"nominal" swagger:"description(Nominal)"`
	PaymentAt time.Time `json:"payment_at,omitempty" swagger:"description(PaymentAt)"`
}
