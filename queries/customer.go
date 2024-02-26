package queries

import (
	"errors"

	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
)

func InsertCustomer(customer *models.Customer) (uint64, error) {
	//insert new customer returning primary keys
	_, insertErr := db.DB.Model(customer).Returning("id").Insert()

	if insertErr != nil {
		return 0, insertErr
	}
	return customer.ID, nil
}


func BulkInsertCustomer(customers []models.Customer) ( error) {
	_, insertErr := db.DB.Model(&customers).Insert()

	if insertErr != nil {
		return insertErr
	}
	return  nil
}


func SelectAllCustomers() ([]models.Customer, error) {
	var customers []models.Customer
	getErr := db.DB.Model(&customers).Select()

	if getErr != nil {
		return nil, getErr
	}
	return customers, nil
}


func SelectAllAccountsByCustomerID(cid uint64) ([]models.Account, error){
	var accounts []models.Account
	customer_accounts := db.DB.Model((*models.AccountToCustomer)(nil)).ColumnExpr("account_id").Where("customer_id = ?",cid)
	err := db.DB.Model(&accounts).Where("id IN (?)",customer_accounts).Select()
	if err != nil {
		return nil, err
	}
	return accounts,nil
}

func SelectCustomerByID(id uint64) (*models.Customer, error) {
	customer := new(models.Customer)
	err := db.DB.Model(customer).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func UpdateCustomer(customer *models.Customer) (uint64, error) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		return 0, txErr
	}

	res, err := tx.Model(customer).WherePK().Returning("id").Update()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if res.RowsAffected() == 0 {
		tx.Rollback()
		return 0, errors.New("no record found to update")
	}
	tx.Commit()
	return customer.ID, nil
}

func DeleteAllCustomers() (uint64, error) {
	customer := new(models.Customer)

	res, err := db.DB.Query(customer, "DELETE from customers")
	if err != nil {
		return 0, nil
	}
	return uint64(res.RowsAffected()), nil
}

func DeleteCustomerByID(id uint64) (uint64, error) {
	customer := new(models.Customer)

	_, err := db.DB.Model(customer).Where("id = ?", id).Returning("id").Delete()
	if err != nil {
		return 0, nil
	}
	return customer.ID, nil
}
