package queries

import (
	"errors"
	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/go-pg/pg/v10"
)

func InsertTransaction(tx *pg.Tx,transaction *models.Transaction) (uint64, error) {
	//insert new transaction returning primary keys
	_, insertErr := db.DB.Model(transaction).Returning("id").Insert()

	if insertErr != nil {
		return 0, insertErr
	}
	return transaction.ID, nil
}

func SelectAllTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	getErr := db.DB.Model(&transactions).Select()

	if getErr != nil {
		return nil, getErr
	}
	return transactions, nil
}

func SelectTransactionByID(id uint64) (*models.Transaction, error) {
	transaction := new(models.Transaction)
	err := db.DB.Model(transaction).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}


func SelectTransactionByAccountID(id uint64) ([]models.Transaction, error) {
	transaction := new([]models.Transaction)
	err := db.DB.Model(transaction).Where("account_id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return *transaction, nil
}

func UpdateTransaction(transaction *models.Transaction) (uint64, error) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		return 0, txErr
	}

	res, err := tx.Model(transaction).WherePK().Returning("id").Update()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if res.RowsAffected() == 0 {
		tx.Rollback()
		return 0, errors.New("no record found to update")
	}
	tx.Commit()
	return transaction.ID, nil
}

func DeleteAllTransactions() (uint64, error) {
	transaction := new(models.Transaction)

	res, err := db.DB.Query(transaction, "DELETE from transactions")
	if err != nil {
		return 0, nil
	}
	return uint64(res.RowsAffected()), nil
}

func DeleteTransactionByID(id uint64) (uint64, error) {
	transaction := new(models.Transaction)

	_, err := db.DB.Model(transaction).Where("id = ?", id).Returning("id").Delete()
	if err != nil {
		return 0, nil
	}
	return transaction.ID, nil
}
