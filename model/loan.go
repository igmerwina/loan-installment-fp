package model

import "time"

type Loan struct {
	ID                int64     `json:"id" gorm:"primaryKey"`
	UserID            int64     `json:"user_id"`
	Item              string    `json:"item"`
	Nominal           int64     `json:"nominal"`
	PeriodInstallment int64     `json:"period_installment"`
	CostInstallment   int64     `json:"cost_installment"`
	StatusLoan        string    `json:"status_loan"`
	Paid              int64     `json:"paid"`
	RemainInstallment int64     `json:"remain_installment"`
	CreatedAt         time.Time `json:"created_at"`
}
