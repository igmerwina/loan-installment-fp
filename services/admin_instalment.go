package services

import (
	"fmt"
	"goloan/config"
	"goloan/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllInstallment retrieves all installments
func GetAdminInstallment(c echo.Context) error {
	db := config.GetDB()

	var installments []model.Installment
	db.Find(&installments)

	fmt.Println("GetAdminInstallment")

	return c.JSON(http.StatusOK, installments)
}
