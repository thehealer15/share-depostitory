package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"share-depository/src/platform/company"
	"share-depository/src/platform/investor"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Company API
// @version 1.0
// @description API for managing companies
// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Initialize Gin router
	r := gin.Default()
	/*

	 */
	// Database connection
	db, err := sql.Open("postgres", "host=localhost port=6969 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Register API routes
	r.POST("/api/company/add", company.AddCompanyHandler(db))
	r.DELETE("/api/company/delete", company.RemoveCompanyHandler(db))
	r.GET("/api/company/", company.GetAllCompaniesHandler(db))
	// Swagger UI setup
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/api/investor/add", investor.AddInvestorHandler(db))

	// stock buy sell APIs

	r.POST("/api/investor/credit_shares", investor.CreditSharesHandler(db))
	r.POST("/api/investor/debit_shares", investor.DebitSharesHandler(db))

	// list all shares
	r.GET("/api/investor/portfolio", investor.GetInvestorPortfolio(db))

	// Start server
	r.Run(":8080")
}
