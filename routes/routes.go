package routes

import (
	"github.com/abhishek-bhangalia-busy/banking-api/handlers"
	"github.com/gin-gonic/gin"
)

func CreateRoutes() {
	router := gin.Default()

	bankRoutes := router.Group("/bank")
	bankRoutes.POST("", handlers.CreateBank)
	bankRoutes.GET("", handlers.GetAllBanks)
	bankRoutes.GET("/branch", handlers.GetAllBanksWithBranches)
	bankRoutes.GET("/:id", handlers.GetBankByID)
	bankRoutes.GET("/:id/branch", handlers.GetAllBranchesOfBankByID)
	bankRoutes.PATCH("", handlers.UpdateBank)
	bankRoutes.DELETE("", handlers.DeleteAllBanks)
	bankRoutes.DELETE("/:id", handlers.DeleteBankByID)

	branchRoutes := router.Group("/branch")
	branchRoutes.POST("", handlers.CreateBranch)
	branchRoutes.GET("", handlers.GetAllBranches)
	branchRoutes.GET("/bank/account", handlers.GetAllBranchesWithBankAndAccounts)
	branchRoutes.GET("/:id", handlers.GetBranchByID)
	branchRoutes.GET("/:id/account", handlers.GetAllAccountsOfBranchByID)
	branchRoutes.PATCH("", handlers.UpdateBranch)
	branchRoutes.DELETE("", handlers.DeleteAllBranches)
	branchRoutes.DELETE("/:id", handlers.DeleteBranchByID)
	//get all customers and accounts

	accountRoutes := router.Group("/account")
	accountRoutes.POST("", handlers.CreateAccount)
	accountRoutes.GET("", handlers.GetAllAccounts)
	accountRoutes.GET("/:id", handlers.GetAccountByID)
	accountRoutes.GET("/:id/customer", handlers.GetAllCustomersByAccountID)
	accountRoutes.PATCH("", handlers.UpdateAccount)
	accountRoutes.DELETE("", handlers.DeleteAllAccounts)
	accountRoutes.DELETE("/:id", handlers.DeleteAccountByID)
	//add joint account

	customerRoutes := router.Group("/customer")
	customerRoutes.POST("", handlers.CreateCustomer)
	customerRoutes.GET("", handlers.GetAllCustomers)
	customerRoutes.GET("/:id", handlers.GetCustomerByID)
	customerRoutes.GET("/:id/account", handlers.GetAllAccountsByCustomerID)
	customerRoutes.PATCH("", handlers.UpdateCustomer)
	customerRoutes.DELETE("", handlers.DeleteAllCustomers)
	customerRoutes.DELETE("/:id", handlers.DeleteCustomerByID)

	mappingRoutes := router.Group("/account_to_customer")
	mappingRoutes.POST("", handlers.CreateMapping)
	mappingRoutes.GET("", handlers.GetAllMappings)
	mappingRoutes.GET("/:id", handlers.GetMappingByID)
	mappingRoutes.PATCH("", handlers.UpdateMapping)
	mappingRoutes.DELETE("", handlers.DeleteAllMappings)
	mappingRoutes.DELETE("/:id", handlers.DeleteMappingByID)

	transactionRoutes := router.Group("/transaction")
	transactionRoutes.POST("", handlers.CreateTransaction)
	transactionRoutes.GET("", handlers.GetAllTransactions)
	transactionRoutes.GET("/:id", handlers.GetTransactionByID)
	transactionRoutes.GET("/account/:id", handlers.GetTransactionByAccountID)
	transactionRoutes.DELETE("", handlers.DeleteAllTransactions)
	transactionRoutes.DELETE("/:id", handlers.DeleteTransactionByID)

	router.Run(":8080")
}
