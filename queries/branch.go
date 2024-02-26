package queries

import (
	"errors"

	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
)

func InsertBranch(branch *models.Branch) (uint64, error) {
	//insert new branch returning primary keys
	_, insertErr := db.DB.Model(branch).Returning("id").Insert()

	if insertErr != nil {
		return 0, insertErr
	}
	return branch.ID, nil
}

func BulkInsertBranch(branches []models.Branch) (error) {
	_, insertErr := db.DB.Model(&branches).Insert()

	if insertErr != nil {
		return  insertErr
	}
	return  nil
}

func SelectAllBranches() ([]models.Branch, error) {
	var branches []models.Branch
	getErr := db.DB.Model(&branches).Select()

	if getErr != nil {
		return nil, getErr
	}
	return branches, nil
}

func SelectAllBranchesWithBankAndAccounts() ([]models.Branch, error) {
	var branches []models.Branch
	getErr := db.DB.Model(&branches).Relation("Accounts").Relation("Bank").Select()

	if getErr != nil {
		return nil, getErr
	}
	return branches, nil
}

func SelectBranchByID(id uint64) (*models.Branch, error) {
	branch := new(models.Branch)
	err := db.DB.Model(branch).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return branch, nil
}

func SelectAllAccountsOfBranchByID(id uint64) ([]models.Account, error){
	accounts := new([]models.Account)
	err := db.DB.Model(accounts).Where("branch_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return *accounts, nil
}

func SelectAllCustomersOfBranchByID(bid uint64) ([]models.Customer, error){
	customers := new([]models.Customer)
	accountIDsOfBranch := db.DB.Model((*models.Account)(nil)).Column("id").Where("branch_id = ?",bid)
	customerIDs := db.DB.Model((*models.AccountToCustomer)(nil)).Column("customer_id").Where("account_id IN (?)",accountIDsOfBranch)
	
	err := db.DB.Model(customers).Where("id IN (?)", customerIDs).Select()
	if err != nil {
		return nil, err
	}
	return *customers, nil
}

func UpdateBranch(branch *models.Branch) (uint64, error) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		return 0, txErr
	}

	res, err := tx.Model(branch).WherePK().Returning("id").Update()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if res.RowsAffected() == 0 {
		tx.Rollback()
		return 0, errors.New("no record found to update")
	}
	tx.Commit()
	return branch.ID, nil
}

func DeleteAllBranches() (uint64, error) {
	branch := new(models.Branch)

	// both methods works
	// res,err := db.DB.Model(branch).Where("true").Delete()
	res, err := db.DB.Query(branch, "DELETE from branches")
	if err != nil {
		return 0, nil
	}
	return uint64(res.RowsAffected()), nil
}

func DeleteBranchByID(id uint64) (uint64, error) {
	branch := new(models.Branch)

	_, err := db.DB.Model(branch).Where("id = ?", id).Returning("id").Delete()
	if err != nil {
		return 0, err
	}
	return branch.ID, nil
}
