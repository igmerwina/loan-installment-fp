package services

import (
	"fmt"
	"goloan/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAllInstallment retrieves all installments
func GetAllInstallment(c echo.Context) error {
	var installments []model.Installment
	db.Find(&installments)

	fmt.Println("GetAllInstallment")

	return c.JSON(http.StatusOK, installments)
}

// CreateInstallment creates a new installment
func CreateInstallment(c echo.Context) error {
	installment := model.Installment{}

	fmt.Println(installment)

	if err := c.Bind(&installment); err != nil {
		return err
	}

	db.Debug().Create(&installment)

	fmt.Println("CreateInstallment")
	return c.JSON(http.StatusOK, installment)
}

// UpdateInstallment updates an existing installment
func UpdateInstallment(c echo.Context) error {
	id := c.Param("installmentId")

	var installment model.Installment
	if err := db.First(&installment, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Installment not found"})
	}

	var updatedInstallment model.Installment
	if err := c.Bind(&updatedInstallment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	fmt.Println("UpdateInstallment")

	db.Model(&installment).Updates(&updatedInstallment)
	return c.JSON(http.StatusOK, installment)
}

// DeleteInstallment deletes an installment
func DeleteInstallment(c echo.Context) error {
	installment := model.Installment{}

	delResp := model.Response{
		Status:  http.StatusOK,
		Message: "Delete Data Success",
	}

	if err := c.Bind(&installment); err != nil {
		return err
	}

	paramID := c.Param("installmentId")

	// Then delete the installment
	if err := db.Delete(&model.Installment{}, paramID).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	db.Model(&installment).Where("installment_id = ?", paramID).Delete(&installment)

	fmt.Println("DeleteInstallment")

	return c.JSON(http.StatusOK, delResp)
}
