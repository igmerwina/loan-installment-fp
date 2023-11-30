package model

import "time"

type Loan struct {
	ID                int64     `json:"id" gorm:"primaryKey" swagger:"description(ID)"`
	UserID            int64     `json:"user_id" swagger:"description(UserID)"`
	Item              string    `json:"item" swagger:"description(Item)"`
	Nominal           int64     `json:"nominal" swagger:"description(Nominal)"`
	PeriodInstallment int64     `json:"period_installment" swagger:"description(PeriodInstallment)"`
	CostInstallment   int64     `json:"cost_installment" swagger:"description(CostInstallment)"`
	StatusLoan        string    `json:"status_loan" swagger:"description(StatusLoan)"`
	Paid              int64     `json:"paid" swagger:"description(Paid)"`
	RemainInstallment int64     `json:"remain_installment" swagger:"description(RemainInstallment)"`
	CreatedAt         time.Time `json:"created_at" swagger:"description(CreatedAt)"`
}
