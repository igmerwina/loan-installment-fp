package services

import (
	"goloan/config"
	"goloan/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// GetAllTransactionsByUserID retrieves all transactions for a given user ID
func GetAllTransactionsByUserID(c echo.Context) error {
	// Parse user ID from the request parameters
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid user ID"})
	}

	// Get the database instance
	db := config.GetDB()

	// Retrieve all transactions for the given user ID
	var transactions []model.Transaction
	if err := db.Where("user_id = ?", userID).Find(&transactions).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get transactions"})
	}

	return c.JSON(http.StatusOK, transactions)
}
