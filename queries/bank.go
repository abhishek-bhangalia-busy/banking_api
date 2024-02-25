package queries

import (
	"errors"

	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
)

func InsertBank(bank *models.Bank) (uint64, error) {
	//insert new bank returning primary keys
	_, insertErr := db.DB.Model(bank).Returning("id").Insert()

	if insertErr != nil {
		return 0, insertErr
	}
	return bank.ID, nil
}

func SelectAllBanks() ([]models.Bank, error) {
	var banks []models.Bank
	getErr := db.DB.Model(&banks).Select()

	if getErr != nil {
		return nil, getErr
	}
	return banks, nil
}

func SelectAllBanksWithBranches() ([]models.Bank, error) {
	var banks []models.Bank
	getErr := db.DB.Model(&banks).Relation("Branches").Select()

	if getErr != nil {
		return nil, getErr
	}
	return banks, nil
}

func SelectBankByID(id uint64) (*models.Bank, error) {
	bank := new(models.Bank)
	err := db.DB.Model(bank).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return bank, nil
}

func SelectAllBranchesOfBankByID(id uint64) ([]models.Branch, error){
	branches := new([]models.Branch)
	err := db.DB.Model(branches).Where("bank_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return *branches, nil
}

func UpdateBank(bank *models.Bank) (uint64, error) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		return 0, txErr
	}

	res, err := tx.Model(bank).WherePK().Returning("id").Update()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if res.RowsAffected() == 0 {
		tx.Rollback()
		return 0, errors.New("no record found to update")
	}
	tx.Commit()
	return bank.ID, nil
}

func DeleteAllBanks() (uint64, error) {
	bank := new(models.Bank)

	// both methods works
	// res,err := db.DB.Model(bank).Where("true").Delete()
	res, err := db.DB.Query(bank, "DELETE from banks")
	if err != nil {
		return 0, nil
	}
	return uint64(res.RowsAffected()), nil
}

func DeleteBankByID(id uint64) (uint64, error) {
	bank := new(models.Bank)

	_, err := db.DB.Model(bank).Where("id = ?", id).Returning("id").Delete()
	if err != nil {
		return 0, nil
	}
	return bank.ID, nil
}
