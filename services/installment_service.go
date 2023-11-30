package services

import (
	"fmt"
	"goloan/config"
	"goloan/model"
	"net/http"

	"github.com/labstack/echo/v4"
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
