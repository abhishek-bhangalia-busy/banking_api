package handlers

import (
	"net/http"
	"strconv"

	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/abhishek-bhangalia-busy/banking-api/queries"
	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBind(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, insertErr := queries.InsertCustomer(&customer)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Customer created with id : ": id})
}

func GetAllCustomers(c *gin.Context) {
	customers, err := queries.SelectAllCustomers()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customers": customers})
}

func GetAllAccountsByCustomerID(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	accounts, err := queries.SelectAllAccountsByCustomerID(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Accounts of customer ": accounts})
}

func GetCustomerByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	customerPtr, err := queries.SelectCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customer": *customerPtr})
}

func UpdateCustomer(c *gin.Context) {
	customer := new(models.Customer)
	err := c.ShouldBindJSON(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated_customer_id, err := queries.UpdateCustomer(customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Update customer with id = ": updated_customer_id})
}

func DeleteAllCustomers(c *gin.Context) {
	rows_deleted, err := queries.DeleteAllCustomers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total rows deleted ": rows_deleted})
}

func DeleteCustomerByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	deleted_id, err := queries.DeleteCustomerByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Deleted Customer with id = ": deleted_id})
}

