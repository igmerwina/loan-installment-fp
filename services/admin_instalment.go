package services

import (
	"fmt"
	"goloan/config"
	"goloan/model"
	"net/http"
	_ "goloan/docs"

	"github.com/labstack/echo/v4"
)


// GetAllInstallment retrieves all installments
// GetAdminInstallment Godoc
// @Summary GetAdminInstallment
// @Description GetAdminInstallment
// @Tags admin
// @Accept  json
// @Produce  json
// @Param model.Installment body model.Installment true "installment"
// @Success 200 {object} model.Installment
// @Router /admin-installment [get]
func GetAdminInstallment(c echo.Context) error {
	db := config.GetDB()

	var installments []model.Installment
	db.Find(&installments)

	fmt.Println("GetAdminInstallment")

	return c.JSON(http.StatusOK, installments)
}
