package db

import (

	"github.com/abhishek-bhangalia-busy/banking-api/models"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

var DB *pg.DB

func ConnectToDB() *pg.DB{
	DB = pg.Connect(&pg.Options{
		Addr:     ":5432",
		User:     "postgres",
		Password: "1234",
		Database: "bank",
	})
	return DB
}


func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*models.Bank)(nil),
		(*models.Branch)(nil),
		(*models.Customer)(nil),
		(*models.Account)(nil),
		(*models.AccountToCustomer)(nil),
		(*models.Transaction)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
			FKConstraints: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
