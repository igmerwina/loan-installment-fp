package main

import (
	"goloan/config"
	"goloan/services"

	"github.com/labstack/echo/v4"
)

func main() {
	config.Connect()

	e := echo.New()

	cust := e.Group("/customer")

	cust.POST("/create-loan", services.CreateLoan)
	cust.PUT("/update/:Id", services.UpdateLoan)
	// cust.GET("/detail-loan/:id", services.CreateLoan) // TODO

	cust.GET("/create-installment", services.CreateInstallment)
	cust.GET("/history-instalment", services.GetAllInstallment)

	e.Logger.Fatal(e.Start(":8021"))
}
