package handlers

import (
	"net/http"
	"strconv"

	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/abhishek-bhangalia-busy/banking-api/queries"
	"github.com/gin-gonic/gin"
)

func CreateMapping(c *gin.Context) {
	var mapping models.AccountToCustomer

	if err := c.ShouldBind(&mapping); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	id, insertErr := queries.InsertMapping(&mapping)
	if insertErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": insertErr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Mapping created with id : ": id})
}

func GetAllMappings(c *gin.Context) {
	mappings, err := queries.SelectAllMappings()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mappings": mappings})
}

func GetMappingByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	mappingPtr, err := queries.SelectMappingByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Mapping": *mappingPtr})
}

func UpdateMapping(c *gin.Context) {
	mapping := new(models.AccountToCustomer)
	err := c.ShouldBindJSON(mapping)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated_mapping_id, err := queries.UpdateMapping(mapping)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"Update mapping with id = ": updated_mapping_id})
}

func DeleteAllMappings(c *gin.Context) {
	rows_deleted, err := queries.DeleteAllMappings()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"total rows deleted ": rows_deleted})
}

func DeleteMappingByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		panic(err)
	}
	deleted_id, err := queries.DeleteMappingByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Deleted Mapping with id = ": deleted_id})
}
