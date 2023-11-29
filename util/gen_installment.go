package util

import (
	"goloan/config"
	"goloan/model"
	"math"
)

// Generate Installment Based On Period
func GenerateInstallment(loan *model.Loan) {
	db := config.GetDB()
	// Generate Installments based on PeriodInstallment
	for i := 0; i < int(loan.PeriodInstallment); i++ {
		nomInstallment := int64(math.Round(float64(loan.Nominal / loan.PeriodInstallment)))
		installment := model.Installment{
			UserID:  loan.UserID,
			LoanID:  loan.ID,
			Nominal: nomInstallment,
		}

		db.Create(&installment)
	}
}
