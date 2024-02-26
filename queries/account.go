package queries

import (
	"errors"
	// "fmt"
	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/go-pg/pg/v10"
)

func InsertAccount(account *models.Account) (uint64, error) {
	//insert new account returning primary keys
	_, insertErr := db.DB.Model(account).Returning("id").Insert()

	if insertErr != nil {
		return 0, insertErr
	}
	return account.ID, nil
}

func BulkInsertAccount(accounts []models.Account) (error) {
	_, insertErr := db.DB.Model(&accounts).Insert()

	if insertErr != nil {
		return insertErr
	}
	return nil
}

func SelectAllAccounts() ([]models.Account, error) {
	var accounts []models.Account
	getErr := db.DB.Model(&accounts).Select()

	if getErr != nil {
		return nil, getErr
	}
	return accounts, nil
}


func SelectAllCustomersByAccountID(aid uint64) ([]models.Customer, error){
	var customers []models.Customer
	customer_accounts := db.DB.Model((*models.AccountToCustomer)(nil)).ColumnExpr("customer_id").Where("account_id = ?",aid)
	err := db.DB.Model(&customers).Where("id IN (?)",customer_accounts).Select()
	if err != nil {
		return nil, err
	}
	return customers,nil
}


func SelectAccountByID(id uint64) (*models.Account, error) {
	account := new(models.Account)
	err := db.DB.Model(account).Where("id = ?", id).Select()
	if err != nil {
		return nil, err
	}
	return account, nil
}

func UpdateAccount(account *models.Account) (uint64, error) {
	tx, txErr := db.DB.Begin()
	if txErr != nil {
		return 0, txErr
	}

	res, err := tx.Model(account).WherePK().Returning("id").Update()
	if err != nil {
		tx.Rollback()
		return 0, err
	}
	if res.RowsAffected() == 0 {
		tx.Rollback()
		return 0, errors.New("no record found to update")
	}
	tx.Commit()
	return account.ID, nil
}

func DeleteAllAccounts() (uint64, error) {
	account := new(models.Account)

	res, err := db.DB.Query(account, "DELETE from accounts")
	if err != nil {
		return 0, nil
	}
	return uint64(res.RowsAffected()), nil
}

func DeleteAccountByID(id uint64) (uint64, error) {
	account := new(models.Account)

	_, err := db.DB.Model(account).Where("id = ?", id).Returning("id").Delete()
	if err != nil {
		return 0, nil
	}
	return account.ID, nil
}

func Deposit(tx *pg.Tx, id uint64, amount float64) (float64, error) {
	account := new(models.Account)
	account.ID = id
	err := tx.Model(account).WherePK().Select()
	if err != nil {
		return account.Balance, err
	}

	account.Balance += amount
	_, err = tx.Model(account).Column("balance").WherePK().Returning("balance").Update()
	if err != nil {
		return account.Balance, err
	}
	return account.Balance, nil
}

func Withdraw(tx *pg.Tx, id uint64, amount float64) (float64, error) {
	account := new(models.Account)
	account.ID = id
	err := tx.Model(account).WherePK().Select()
	if err != nil {
		return account.Balance, err
	}
	if account.Balance < amount {
		return account.Balance, errors.New("account balance is less than amount")
	}
	account.Balance -= amount
	_, err = tx.Model(account).Column("balance").WherePK().Returning("balance").Update()
	if err != nil {
		return account.Balance, err
	}
	return account.Balance, nil
}

func Transfer(tx *pg.Tx, sid uint64, rAccNo uint64, amount float64) error {
	// var saccount models.Account
	// var raccount models.Account
	// saccount.ID = sid

	// err := tx.Model(&saccount).WherePK().Select()
	// if err != nil {
	// 	return err
	// }
	// err = tx.Model(&raccount).Where("account_no = ?",rAccNo).Select()
	// if err != nil {
	// 	return err
	// }

	// if saccount.Balance < amount {
	// 	return errors.New("balance is less than amount")
	// }

	// saccount.Balance -= amount
	// fmt.Println("receiver acc no is ",rAccNo)

	// _, err = tx.Model(&saccount).Column("balance").WherePK().Update()
	// if err != nil {
	// 	return err
	// }
	// raccount.Balance += amount
	// _, err = tx.Model(&raccount).Column("balance").WherePK().Update()
	// if err != nil {
	// 	return err
	// }
	// // _, err = tx.Query(raccount, "UPDATE accounts SET balance = balance + ? WHERE account_no = ?",amount,rAccNo)

	// // // _, err = tx.Model(&raccount).Set("balance = balance + ?", amount).Where("account_no = ?", rAccNo).Update()
	// // if err != nil {
	// // 	return err
	// // }
	// fmt.Println("sender balance,",saccount.Balance)
	// fmt.Print("receiver balance",raccount.Balance)
	// return nil

	var account models.Account

	res,err := tx.Query(account, "UPDATE accounts SET balance = balance - ? WHERE id = ? AND balance - ?0 >= 0",amount, sid)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return errors.New("not a valid request (either id is wrong or balance is not sufficient)")
	}

	_,err = tx.Query(account, "UPDATE accounts SET balance = balance + ? WHERE account_no = ?",amount, rAccNo)
	if err != nil {
		return err
	}
	return nil
}
