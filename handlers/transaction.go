package handlers

import (
	// "fmt"
	// "fmt"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/abhishek-bhangalia-busy/banking-api/queries"
	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		// tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": "can't create transaction"})
		return
	}

	var transaction models.Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		// tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	_, insertErr := queries.InsertTransaction(tx, &transaction)
	if insertErr != nil {
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}
	// fmt.Print(transaction.AccountID)
	var newBal float64
	var err error

	switch transaction.Mode {
	case "deposit":
		newBal, err = queries.Deposit(tx, transaction.AccountID, transaction.Amount)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
			return
		}
		tx.Commit()
		c.JSON(http.StatusCreated, gin.H{"Transaction created successfully and new balance is : ": newBal})
	case "withdraw":
		newBal, err = queries.Withdraw(tx, transaction.AccountID, transaction.Amount)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tx.Commit()
		c.JSON(http.StatusCreated, gin.H{"Transaction created successfully and new balance is : ": newBal})
	case "transfer":
		fmt.Println(transaction.ReceiverAccountNo)
		if transaction.ReceiverAccountNo == 0 {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": errors.New("receiver account no is not sent").Error()})
			return
		}
		err = queries.Transfer(tx, transaction.AccountID, transaction.ReceiverAccountNo, transaction.Amount)
		if err != nil {
			tx.Rollback()
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		tx.Commit()
		c.JSON(http.StatusCreated, gin.H{"Transaction created successfully": ""})
	}
}

func GetAllTransactions(c *gin.Context) {
	transactions, err := queries.SelectAllTransactions()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Transactions": transactions})
}

func GetTransactionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	transactionPtr, err := queries.SelectTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Transaction": *transactionPtr})
}

func GetTransactionByAccountID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	transaction, err := queries.SelectTransactionByAccountID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Transaction": transaction})
}


func DeleteAllTransactions(c *gin.Context) {
	rows_deleted, err := queries.DeleteAllTransactions()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total rows deleted ": rows_deleted})
}


func DeleteTransactionByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	deleted_id, err := queries.DeleteTransactionByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Deleted Transaction with id = ": deleted_id})
}
