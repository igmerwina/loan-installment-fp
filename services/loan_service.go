package services

import (
	"fmt"
	"goloan/config"
	"goloan/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetAllLoan(c echo.Context) error {
	db := config.GetDB()

	var loan []model.Loan
	db.Find(&loan)

	fmt.Println("GetAllLoan")

	return c.JSON(http.StatusOK, loan)
}

func GetDetailLoan(c echo.Context) error {
	db := config.GetDB()

	var loan model.Loan

	userId := c.Param("userId")
	result := db.Where("user_id = ?", userId).First(&loan)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	fmt.Println("GetAllLoan")

	return c.JSON(http.StatusOK, loan)
}

func CreateLoan(c echo.Context) error {
	db := config.GetDB()

	loan := model.Loan{}

	fmt.Println(loan)

	if err := c.Bind(&loan); err != nil {
		return err
	}

	db.Debug().Save(&loan)

	fmt.Println("CreateLoan")
	return c.JSON(http.StatusOK, loan)
}

func UpdateLoan(c echo.Context) error {
	db := config.GetDB()

	id := c.Param("Id")

	var loan model.Loan
	if err := db.First(&loan, id).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Loan not found"})
	}

	var updatedLoan model.Loan
	if err := c.Bind(&updatedLoan); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid input"})
	}

	fmt.Println("UpdateLoan")

	db.Model(&loan).Updates(&updatedLoan)
	return c.JSON(http.StatusOK, loan)
}

func DeleteLoan(c echo.Context) error {
	db := config.GetDB()

	loan := model.Loan{}

	delResp := model.Response{
		Status:  http.StatusOK,
		Message: "Delete Data Success",
	}

	if err := c.Bind(&loan); err != nil {
		return err
	}

	paramId := c.Param("orderId")

	// Then delete the order
	if err := db.Delete(&model.Loan{}, paramId).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	db.Model(&loan).Where("order_id = ?", paramId).Delete(&loan)

	fmt.Println("DeleteLoan")

	return c.JSON(http.StatusOK, delResp)
}
