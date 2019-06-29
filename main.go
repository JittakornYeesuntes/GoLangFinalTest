package main

import (
	"github.com/JittakornYeesuntes/finalexam/customer"
	"github.com/JittakornYeesuntes/finalexam/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	r := setupRoute()
	r.Run(":2019")
}

func setupRoute() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Authorization)
	r.POST("/customers", customer.CreateCustomerHandler)
	r.GET("/customers/:id", customer.GetCustomerByIDHandler)
	r.GET("/customers", customer.GetAllCustomerHandler)
	r.PUT("/customers/:id", customer.PutUpdateCustomerHandler)
	r.DELETE("/customers/:id", customer.DeleteCustomerByIDHandler)
	return r
}