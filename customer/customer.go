package customer

import (
	"github.com/JittakornYeesuntes/finalexam/database"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Customer struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Status string `json:"status"`
}

func CreateCustomerHandler(c *gin.Context) {
	customer := Customer{}
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	row, err := database.InsertCustomer(customer.Name, customer.Email, customer.Status)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	log.Println("insert success : ", customer)
	c.JSON(http.StatusCreated, customer)
}

func GetCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	row, err := database.SelectByID(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	customer := Customer{}
	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	log.Println("select success : ", customer)
	c.JSON(http.StatusOK, customer)
}

func GetAllCustomerHandler(c *gin.Context) {
	rows, err := database.SelectAll()
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	customers := []Customer{}
	for rows.Next() {
		ctm := Customer{}
		err = rows.Scan(&ctm.ID, &ctm.Name, &ctm.Email, &ctm.Status)
		if err != nil {
			log.Println(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
			return
		}
		customers = append(customers, ctm)
	}
	log.Println("select all success : ", customers)
	c.JSON(http.StatusOK, customers)
}

func PutUpdateCustomerHandler(c *gin.Context) {
	id := c.Param("id")
	customer := Customer{}
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	row, err := database.UpdateByID(id, customer.Name, customer.Email, customer.Status)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}

	err = row.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.Status)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	log.Println("update success : ", customer)
	c.JSON(http.StatusOK, customer)
}

func DeleteCustomerByIDHandler(c *gin.Context) {
	id := c.Param("id")
	err := database.DeleteByID(id)
	if err != nil {
		log.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": http.StatusText(http.StatusInternalServerError)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":"customer deleted"})
}
