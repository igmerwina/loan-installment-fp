package services

import (
	"fmt"
	"goloan/config"
	"goloan/model"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

// GetAllInstallment retrieves all installments
// CreateEmployee Godoc
// @Summary GetAllInstallment
// @Description GetAllInstallment
// @Tags installment
// @Accept  json
// @Produce  json
// @Param model.Installment body model.Installment true "installment"
// @Success 200 {object} model.Installment
// @Router /all-installment [get]
func GetAllInstallment(c echo.Context) error {
	db := config.GetDB()

	var installments []model.Installment
	db.Find(&installments)

	fmt.Println("GetAllInstallment")

	return c.JSON(http.StatusOK, installments)
}

// GetAllInstallment retrieves all installments
// CreateEmployee Godoc
// @Summary GetHistInstallment
// @Description GetHistInstallment
// @Tags installment
// @Accept  json
// @Produce  json
// @Param model.Installment body model.Installment true "installment"
// @Success 200 {object} model.Installment
// @Router /history-instalment/{userId} [get]
func GetHistInstallment(c echo.Context) error {
	db := config.GetDB()

	userId := c.Param("userId")

	var installment []model.Installment

	if err := db.Where("user_id = ?", userId).Find(&installment).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	fmt.Println("GetHistInstallment", installment)

	resp := model.Response{
		ResponseCode: config.Success,
		ResponseDesc: config.MessageResponse(config.Success),
		Data:         installment,
	}

	return c.JSON(http.StatusOK, resp)
}

// PayInstallment pay the installment made
// @Summary PayInstallment
// @Description Pay the installment made
// @Tags installment
// @Accept  json
// @Produce  json
// @Param model.Transaction body model.Transaction true "Transaction"
// @Success 200 {object} model.Transaction
// @Router /installment [post]
func PayInstallment(c echo.Context) error {
	// Parse the JSON request body into an Installment struct
	var newTransaction model.Transaction

	if err := c.Bind(&newTransaction); err != nil {
		return err
	}

	// Get the database instance
	db := config.GetDB()

	// Call the CreateTransaction function from the model package
	if err := db.Create(&newTransaction).Error; err != nil {
		return err
		// return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create transaction"})
	}

	// Update the loan's remaining amount by subtracting the installment's nominal
	if err := updateRemainingAmount(db, newTransaction.LoanID, newTransaction.Nominal); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update remaining amount"})
	}

	return c.JSON(http.StatusCreated, newTransaction)
}

func updateRemainingAmount(db *gorm.DB, loanID, installmentAmount int64) error {
	var loan model.Loan

	// Find the loan by ID
	if err := db.First(&loan, loanID).Error; err != nil {
		return err
	}

	fmt.Println("Get success id: ", loan.ID)

	// Subtract the installment amount from the loan's remaining amount
	loan.RemainInstallment -= installmentAmount

	fmt.Println("Installment: ", loan.RemainInstallment)
	// Save the updated loan back to the database
	if err := db.Save(&loan).Error; err != nil {
		return err
	}

	return nil
}
