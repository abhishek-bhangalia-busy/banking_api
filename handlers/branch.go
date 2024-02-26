package handlers

import (
	"net/http"
	"strconv"

	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/abhishek-bhangalia-busy/banking-api/queries"
	"github.com/gin-gonic/gin"
)

func CreateBranch(c *gin.Context) {
	var branch models.Branch

	if err := c.ShouldBind(&branch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id, insertErr := queries.InsertBranch(&branch)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Branch created with id : ": id})
}


func BulkCreateBranch(c *gin.Context) {
	var body struct{
		Branches []models.Branch
	}

	if err := c.ShouldBind(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	insertErr := queries.BulkInsertBranch(body.Branches)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"Branch created successfully":""})
}

func GetAllBranches(c *gin.Context) {
	branches, err := queries.SelectAllBranches()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"branch": branches})
}

func GetAllBranchesWithBankAndAccounts(c *gin.Context) {
	branches, err := queries.SelectAllBranchesWithBankAndAccounts()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"branch": branches})
}

func GetBranchByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // 10 is for base 10 digits and 64 for uint64  bit size
	if err != nil {
		panic(err)
	}
	branchPtr, err := queries.SelectBranchByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"branch": *branchPtr})
}

func GetAllAccountsOfBranchByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // 10 is for base 10 digits and 64 for uint64  bit size
	if err != nil {
		panic(err)
	}
	accounts, err := queries.SelectAllAccountsOfBranchByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"accounts": accounts})
}

func GetAllCustomersOfBranchByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64) // 10 is for base 10 digits and 64 for uint64  bit size
	if err != nil {
		panic(err)
	}
	customers, err := queries.SelectAllCustomersOfBranchByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers})
}

func UpdateBranch(c *gin.Context) {
	branch := new(models.Branch)
	err := c.ShouldBindJSON(branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated_branch_id, err := queries.UpdateBranch(branch)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Update branch with id = ": updated_branch_id})
}

func DeleteAllBranches(c *gin.Context) {
	rows_deleted, err := queries.DeleteAllBranches()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total rows deleted ": rows_deleted})
}

func DeleteBranchByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	deleted_id, err := queries.DeleteBranchByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Deleted branch with id = ": deleted_id})
}
