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

// GetAllLoan Godoc
// @Summary GetAllLoan
// @Description GetAllLoan
// @Tags loan
// @Accept  json
// @Produce  json
// @Param model.Loan body model.Loan true "Loan"
// @Success 200 {object} model.Loan
// @Router /all-loan [get]
func GetAllLoan(c echo.Context) error {
	db := config.GetDB()

	var loan []model.Loan
	db.Find(&loan)

	fmt.Println("GetAllLoan")

	return c.JSON(http.StatusOK, loan)
}

// GetDetailLoan Godoc
// @Summary GetDetailLoan
// @Description GetDetailLoan
// @Tags loan
// @Accept  json
// @Produce  json
// @Param model.Loan body model.Loan true "Loan"
// @Success 200 {object} model.Loan
// @Router /detail-loan/{userId} [get]
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

// CreateLoan Godoc
// @Summary Create Loan
// @Description Create Loan
// @Tags loan
// @Accept  json
// @Produce  json
// @Param model.Loan body model.Loan true "Loan"
// @Success 200 {object} model.Loan
// @Router /create-loan [post]
func CreateLoan(c echo.Context) error {
	db := config.GetDB()

	loan := model.Loan{}

	if err := c.Bind(&loan); err != nil {
		return err
	}

	db.Debug().Save(&loan)

	//loan get dari hasil save diatas, bisa ngeget langsungd ari save
	//gorm.lastinsertid
	util.GenerateInstallment(&loan)

	resp := model.Response{
		ResponseCode: config.Success,
		ResponseDesc: config.MessageResponse(config.Success),
		Data:         loan,
	}

	fmt.Println("CreateLoan")
	return c.JSON(http.StatusOK, resp)
}

// UpdateLoan Godoc
// @Summary Update Loan
// @Description Update Loan
// @Tags loan
// @Accept  json
// @Produce  json
// @Param model.Loan body model.Loan true "Loan"
// @Success 200 {object} model.Loan
// @Router /update/{Id} [put]
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
