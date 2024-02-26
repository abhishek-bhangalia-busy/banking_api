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
	bankRoutes.GET("",middlewares.RequireAuth, handlers.GetAllBanks)
	bankRoutes.GET("/branch",middlewares.RequireAuth, handlers.GetAllBanksWithBranches)
	bankRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetBankByID)
	bankRoutes.GET("/:id/branch",middlewares.RequireAuth, handlers.GetAllBranchesOfBankByID)
	bankRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateBank)
	bankRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllBanks)
	bankRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteBankByID)

	branchRoutes := router.Group("/branch")
	branchRoutes.POST("",middlewares.RequireAuth, handlers.CreateBranch)
	branchRoutes.GET("",middlewares.RequireAuth, handlers.GetAllBranches)
	branchRoutes.GET("/bank/account",middlewares.RequireAuth, handlers.GetAllBranchesWithBankAndAccounts)
	branchRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetBranchByID)
	branchRoutes.GET("/:id/account",middlewares.RequireAuth, handlers.GetAllAccountsOfBranchByID)
	branchRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateBranch)
	branchRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllBranches)
	branchRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteBranchByID)
	//get all customers and accounts

	accountRoutes := router.Group("/account")
	accountRoutes.POST("",middlewares.RequireAuth, handlers.CreateAccount)
	accountRoutes.GET("",middlewares.RequireAuth, handlers.GetAllAccounts)
	accountRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetAccountByID)
	accountRoutes.GET("/:id/customer",middlewares.RequireAuth, handlers.GetAllCustomersByAccountID)
	accountRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateAccount)
	accountRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllAccounts)
	accountRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteAccountByID)
	//add joint account

	customerRoutes := router.Group("/customer")
	customerRoutes.POST("",middlewares.RequireAuth, handlers.CreateCustomer)
	customerRoutes.GET("",middlewares.RequireAuth, handlers.GetAllCustomers)
	customerRoutes.GET("/:id",middlewares.RequireAuth, handlers.GetCustomerByID)
	customerRoutes.GET("/:id/account",middlewares.RequireAuth, handlers.GetAllAccountsByCustomerID)
	customerRoutes.PATCH("",middlewares.RequireAuth, handlers.UpdateCustomer)
	customerRoutes.DELETE("",middlewares.RequireAuth, handlers.DeleteAllCustomers)
	customerRoutes.DELETE("/:id",middlewares.RequireAuth, handlers.DeleteCustomerByID)

	mappingRoutes := router.Group("/account_to_customer")
	mappingRoutes.POST("",middlewares.RequireAuth, handlers.CreateMapping)
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
