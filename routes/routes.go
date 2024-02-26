package routes

import (
	"github.com/abhishek-bhangalia-busy/banking-api/handlers"
	"github.com/abhishek-bhangalia-busy/banking-api/middlewares"
	"github.com/gin-gonic/gin"
)

func CreateRoutes() {
	router := gin.Default()

	authRoutes := router.Group("/auth")
	authRoutes.POST("/signup", handlers.Signup)
	authRoutes.POST("/signin", handlers.Signin)

	bankRoutes := router.Group("/bank")
	bankRoutes.POST("", middlewares.RequireAuth, handlers.CreateBank)
	bankRoutes.POST("/bulk", middlewares.RequireAuth, handlers.BulkCreateBanks)	//create multiple banks in single request
	bankRoutes.GET("",middlewares.RequireAuth, handlers.GetAllBanks)
	bankRoutes.GET("/branch",middlewares.RequireAuth, handlers.GetAllBanksWithBranches)
	bankRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetBankByID)
	bankRoutes.GET("/:id/branch",middlewares.RequireAuth, handlers.GetAllBranchesOfBankByID)
	bankRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateBank)
	bankRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllBanks)
	bankRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteBankByID)

	branchRoutes := router.Group("/branch")
	branchRoutes.POST("",middlewares.RequireAuth, handlers.CreateBranch)
	branchRoutes.POST("/bulk",middlewares.RequireAuth, handlers.BulkCreateBranch)
	branchRoutes.GET("",middlewares.RequireAuth, handlers.GetAllBranches)
	branchRoutes.GET("/bank/account",middlewares.RequireAuth, handlers.GetAllBranchesWithBankAndAccounts)
	branchRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetBranchByID)
	branchRoutes.GET("/:id/account",middlewares.RequireAuth, handlers.GetAllAccountsOfBranchByID)
	branchRoutes.GET("/:id/customer",middlewares.RequireAuth, handlers.GetAllCustomersOfBranchByID)
	branchRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateBranch)
	branchRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllBranches)
	branchRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteBranchByID)

	accountRoutes := router.Group("/account")
	accountRoutes.POST("",middlewares.RequireAuth, handlers.CreateAccount)
	accountRoutes.POST("/bulk",middlewares.RequireAuth, handlers.BulkCreateAccount)
	accountRoutes.GET("",middlewares.RequireAuth, handlers.GetAllAccounts)
	accountRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetAccountByID)
	accountRoutes.GET("/:id/customer",middlewares.RequireAuth, handlers.GetAllCustomersByAccountID)
	accountRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateAccount)
	accountRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllAccounts)
	accountRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteAccountByID)
	

	customerRoutes := router.Group("/customer")
	customerRoutes.POST("",middlewares.RequireAuth, handlers.CreateCustomer)
	customerRoutes.POST("/bulk",middlewares.RequireAuth, handlers.BulkCreateCustomer)
	customerRoutes.GET("",middlewares.RequireAuth, handlers.GetAllCustomers)
	customerRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetCustomerByID)
	customerRoutes.GET("/:id/account",middlewares.RequireAuth, handlers.GetAllAccountsByCustomerID)
	customerRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateCustomer)
	customerRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllCustomers)
	customerRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteCustomerByID)

	mappingRoutes := router.Group("/account_to_customer")
	mappingRoutes.POST("",middlewares.RequireAuth, handlers.CreateMapping)
	mappingRoutes.POST("/bulk",middlewares.RequireAuth, handlers.BulkCreateMapping)
	mappingRoutes.GET("",middlewares.RequireAuth, handlers.GetAllMappings)
	mappingRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetMappingByID)
	mappingRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateMapping)
	mappingRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllMappings)
	mappingRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteMappingByID)

	transactionRoutes := router.Group("/transaction")
	transactionRoutes.POST("",middlewares.RequireAuth, handlers.CreateTransaction)
	transactionRoutes.GET("",middlewares.RequireAuth, handlers.GetAllTransactions)
	transactionRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetTransactionByID)
	transactionRoutes.GET("/account/:id",middlewares.RequireAuth, handlers.GetTransactionByAccountID)
	transactionRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllTransactions)
	transactionRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteTransactionByID)

	router.Run(":8080")
}
