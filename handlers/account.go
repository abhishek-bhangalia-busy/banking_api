package handlers

import (
	"net/http"
	"strconv"

	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/abhishek-bhangalia-busy/banking-api/queries"
	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var account models.Account
	if err := c.ShouldBind(&account); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, insertErr := queries.InsertAccount(&account)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Account created with id : ": id})
}

func BulkCreateAccount(c *gin.Context) {
	var body struct{
		Accounts []models.Account
	}
	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertErr := queries.BulkInsertAccount(body.Accounts)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Accounts created successfully : ":""})
}

func GetAllAccounts(c *gin.Context) {
	accounts, err := queries.SelectAllAccounts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Accounts": accounts})
}

func GetAllCustomersByAccountID(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	customers, err := queries.SelectAllCustomersByAccountID(id)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Customers of accounts ": customers})
}


func GetAccountByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	accountPtr, err := queries.SelectAccountByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Account": *accountPtr})
}

func UpdateAccount(c *gin.Context) {
	account := new(models.Account)
	err := c.ShouldBindJSON(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated_account_id, err := queries.UpdateAccount(account)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Update account with id = ": updated_account_id})
}

func DeleteAllAccounts(c *gin.Context) {
	rows_deleted, err := queries.DeleteAllAccounts()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total rows deleted ": rows_deleted})
}

func DeleteAccountByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	deleted_id, err := queries.DeleteAccountByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Deleted Account with id = ": deleted_id})
}

