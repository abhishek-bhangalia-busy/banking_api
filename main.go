package main

import (
	"github.com/abhishek-bhangalia-busy/banking-api/db"
	"github.com/abhishek-bhangalia-busy/banking-api/initializers"
	"github.com/abhishek-bhangalia-busy/banking-api/routes"
	"github.com/go-pg/pg/v10"
)

var DB *pg.DB

func init(){
	initializers.LoadEnvVariables()
}


func main() {
	//connect to DB
	DB = db.ConnectToDB()
	defer DB.Close()

	//creating schema
	err := db.CreateSchema(DB)
	if err != nil {
		panic(err)
	}

	routes.CreateRoutes()
}

