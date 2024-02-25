package queries

import (
	"errors"
	"fmt"

	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
)

func InsertMapping(mapping *models.AccountToCustomer) (uint64, error) {
	//insert new mapping returning primary keys
	_, insertErr := db.DB.Model(mapping).Returning("id").Insert()

	if insertErr != nil {
		fmt.Println("erro si ", insertErr)
		return 0, insertErr
	}
	return mapping.ID, nil
}

func SelectAllMappings() ([]models.AccountToCustomer, error) {
	var mappings []models.AccountToCustomer
	getErr := db.DB.Model(&mappings).Select()

	if getErr != nil {
		return nil, getErr
	}
	return mappings, nil
}

func SelectMappingByID(id uint64) (*models.AccountToCustomer, error) {
	mapping := new(models.AccountToCustomer)
	err := db.DB.Model(mapping).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return mapping, nil
}

func UpdateMapping(mapping *models.AccountToCustomer) (uint64, error) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		return 0, txErr
	}

	res, err := tx.Model(mapping).WherePK().Returning("id").Update()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if res.RowsAffected() == 0 {
		tx.Rollback()
		return 0, errors.New("no record found to update")
	}
	tx.Commit()
	return mapping.ID, nil
}

func DeleteAllMappings() (uint64, error) {
	mapping := new(models.AccountToCustomer)

	res, err := db.DB.Query(mapping, "DELETE from mappings")
	if err != nil {
		return 0, nil
	}
	return uint64(res.RowsAffected()), nil
}

func DeleteMappingByID(id uint64) (uint64, error) {
	mapping := new(models.AccountToCustomer)

	_, err := db.DB.Model(mapping).Where("id = ?", id).Returning("id").Delete()
	if err != nil {
		return 0, nil
	}
	return mapping.ID, nil
}
