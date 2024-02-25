package handlers

import (
	"net/http"
	"strconv"

	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/abhishek-bhangalia-busy/banking-api/queries"
	"github.com/gin-gonic/gin"
)

func CreateBank(c *gin.Context) {
	var bank models.Bank
	if err := c.ShouldBind(&bank); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, insertErr := queries.InsertBank(&bank)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Bank created with id : ": id})
}

func GetAllBanks(c *gin.Context) {
	banks, err := queries.SelectAllBanks()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Banks": banks})
}

func GetAllBanksWithBranches(c *gin.Context) {
	banks, err := queries.SelectAllBanksWithBranches()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Banks": banks})
}

func GetBankByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	bankPtr, err := queries.SelectBankByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Bank": *bankPtr})
}

func GetAllBranchesOfBankByID(c *gin.Context){
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	branches, err := queries.SelectAllBranchesOfBankByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Branches : ":branches})
}

func UpdateBank(c *gin.Context) {
	bank := new(models.Bank)
	err := c.ShouldBindJSON(bank)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated_bank_id, err := queries.UpdateBank(bank)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Update bank with id = ": updated_bank_id})
}

func DeleteAllBanks(c *gin.Context) {
	rows_deleted, err := queries.DeleteAllBanks()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total rows deleted ": rows_deleted})
}

func DeleteBankByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	deleted_id, err := queries.DeleteBankByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Deleted Bank with id = ": deleted_id})
}
