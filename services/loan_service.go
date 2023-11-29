package services

import (
	"fmt"
	"goloan/config"
	"goloan/model"
	"goloan/util"
	"log"
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
	fmt.Println("GetDetailLoan")

	db := config.GetDB()

	var loan model.Loan

	userId := c.Param("userId")
	result := db.Where("user_id = ?", userId).First(&loan)
	if result.Error != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	}

	resp := model.Response{
		ResponseCode: config.Success,
		ResponseDesc: config.MessageResponse(config.Success),
		Data:         result,
	}

	return c.JSON(http.StatusOK, resp)
}

func CreateLoan(c echo.Context) error {
	db := config.GetDB()

	loan := model.Loan{}

	if err := c.Bind(&loan); err != nil {
		return err
	}

	util.GenerateInstallment(&loan)

	db.Debug().Save(&loan)

	resp := model.Response{
		ResponseCode: config.Success,
		ResponseDesc: config.MessageResponse(config.Success),
		Data:         loan,
	}

	fmt.Println("CreateLoan")
	return c.JSON(http.StatusOK, resp)
}

func UpdateLoan(c echo.Context) error {
	log.Println("UpdateLoan")
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

	db.Model(&loan).Updates(&updatedLoan)

	resp := model.Response{
		ResponseCode: config.Success,
		ResponseDesc: config.MessageResponse(config.Success),
		Data:         loan,
	}

	return c.JSON(http.StatusOK, resp)
}

/* unused
func DeleteLoan(c echo.Context) error {
	db := config.GetDB()

	loan := model.Loan{}

	delResp := model.Response{
		ResponseCode: config.Success,
		Message:      "Delete Data Success",
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
*/
