package main

import (
	"goloan/config"
	"goloan/services"

	"github.com/labstack/echo/v4"

	echoSwagger "github.com/swaggo/echo-swagger"
)

func main() {
	env := config.New().SetArgs()

	config.Connect()

	e := echo.New()

	cust := e.Group("/customer")
	cust.POST("/create-loan", services.CreateLoan)
	cust.PUT("/update/:Id", services.UpdateLoan)
	cust.GET("/detail-loan/:userId", services.GetDetailLoan)

	cust.GET("/history-instalment:/:userId", services.GetHistInstallment)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.Fatal(e.Start(":" + env.AppsPort))
}
