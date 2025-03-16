package investor

import (
	"database/sql"
	"fmt"
	"net/http"
	"share-depository/src/service"
	"share-depository/src/util"

	"github.com/gin-gonic/gin"
)

// Investor represents an investor in the platform.investors table
type Investor struct {
	Name   string `json:"investor_name"`
	GovtID string `json:"govt_id"` // Use GovtID to match JSON and DB naming
}

// for buy sell activity
type StockRequest struct {
	GovtID   string `json:"govt_id"`
	Ticker   string `json:"ticker"`
	Quantity int    `json:"quantity"`
}

// to list all stocks a perticular investor holds
type Holding struct {
	Ticker   string `json:"ticker"`
	Quantity int    `json:"quantity"`
}

// AddInvestorHandler handles onboarding a new investor
// @Summary Onboard a new investor
// @Description Adds an investor to platform.investors, creates investor_{govt-id} schema, and sets up tables
// @Tags investor
// @Accept json
// @Produce json
// @Param investor body Investor true "Investor data"
// @Success 201 {object} util.Response "Investor onboarded successfully"
// @Failure 400 {object} util.Response "Invalid request body"
// @Failure 500 {object} util.Response "Failed to onboard investor"
// @Router /api/investor/add [post]
func AddInvestorHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse request
		var inv Investor
		if err := c.ShouldBindJSON(&inv); err != nil {
			util.JSON(c, http.StatusBadRequest, "", "Invalid request body", err.Error())
			return
		}

		// Validate input
		if inv.Name == "" || inv.GovtID == "" {
			util.JSON(c, http.StatusBadRequest, "", "Investor name and govt_id are required", "")
			return
		}

		// Set search path to platform for initial insert
		if err := util.SetSearchPath(db, "platform"); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to set search path to platform", err.Error())
			return
		}

		// Begin a transaction
		tx, err := db.Begin()
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to start transaction", err.Error())
			return
		}
		defer tx.Rollback() // Rollback if not committed

		// Step 1: Add investor to platform.investors
		query := "INSERT INTO platform.investor  (investor_name, govt_id) VALUES ($1, $2) ON CONFLICT (govt_id) DO NOTHING"
		result, err := tx.Exec(query, inv.Name, inv.GovtID)
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to add investor to platform.investors", err.Error())
			return
		}

		// Check if the investor was actually inserted (0 rows affected means conflict)
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to check rows affected", err.Error())
			return
		}
		if rowsAffected == 0 {
			util.JSON(c, http.StatusConflict, "", "Investor with this govt_id already exists", "")
			return
		}

		// Step 2: Create investor_{govt-id} schema
		schemaName := fmt.Sprintf("investor_%s", inv.GovtID)
		_, err = tx.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", schemaName))
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", fmt.Sprintf("Failed to create schema %s", schemaName), err.Error())
			return
		}

		// Step 3: Create tables in the new schema
		_, err = tx.Exec(fmt.Sprintf(`
            CREATE TABLE IF NOT EXISTS %s.investor_details (
                govt_id text PRIMARY KEY
            );
            CREATE TABLE IF NOT EXISTS %s.holdings (
                ticker text PRIMARY KEY,
                quantity int
            );
        `, schemaName, schemaName))
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", fmt.Sprintf("Failed to create tables in schema %s", schemaName), err.Error())
			return
		}

		// Step 4: Insert initial investor_details entry
		_, err = tx.Exec(fmt.Sprintf("INSERT INTO %s.investor_details (govt_id) VALUES ($1)", schemaName), inv.GovtID)
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", fmt.Sprintf("Failed to initialize investor_details in %s", schemaName), err.Error())
			return
		}

		// Commit the transaction
		if err := tx.Commit(); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to commit transaction", err.Error())
			return
		}

		// Success response
		util.JSON(c, http.StatusCreated, "Investor onboarded successfully", "", "")
	}
}

func CreditSharesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req StockRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			util.JSON(c, http.StatusBadRequest, "", "Invalid request body", err.Error())
			return
		}

		if req.GovtID == "" || req.Ticker == "" || req.Quantity <= 0 {
			util.JSON(c, http.StatusBadRequest, "", "govt_id, ticker, and positive quantity are required", "")
			return
		}

		tx, err := db.Begin()
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to start transaction", err.Error())
			return
		}
		defer tx.Rollback()

		schema := fmt.Sprintf("investor_%s", req.GovtID)
		var exists bool
		err = tx.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.schemata WHERE schema_name = $1)"), schema).Scan(&exists)
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to check investor schema", err.Error())
			return
		}
		if !exists {
			util.JSON(c, http.StatusNotFound, "", "Investor not found", "")
			return
		}

		newQuantity, err := service.UpdateStockQuantity(tx, schema, req.Ticker, req.Quantity)
		if err != nil {
			util.JSON(c, http.StatusBadRequest, "", "Failed to credit shares", err.Error())
			return
		}

		if err := tx.Commit(); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to commit transaction", err.Error())
			return
		}

		util.JSON(c, http.StatusOK, fmt.Sprintf("Shares credited successfully, new quantity: %d", newQuantity), "", "")
	}
}

/*
@debitshareHandler
**/

func DebitSharesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req StockRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			util.JSON(c, http.StatusBadRequest, "", "Invalid request body", err.Error())
			return
		}

		if req.GovtID == "" || req.Ticker == "" || req.Quantity <= 0 {
			util.JSON(c, http.StatusBadRequest, "", "govt_id, ticker, and positive quantity are required", "")
			return
		}

		tx, err := db.Begin()
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to start transaction", err.Error())
			return
		}
		defer tx.Rollback()

		schema := fmt.Sprintf("investor_%s", req.GovtID)
		var exists bool
		err = tx.QueryRow(fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM information_schema.schemata WHERE schema_name = $1)"), schema).Scan(&exists)
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to check investor schema", err.Error())
			return
		}
		if !exists {
			util.JSON(c, http.StatusNotFound, "", "Investor not found", "")
			return
		}

		newQuantity, err := service.UpdateStockQuantity(tx, schema, req.Ticker, -req.Quantity)
		if err != nil {
			util.JSON(c, http.StatusBadRequest, "", "Failed to debit shares", err.Error())
			return
		}

		if err := tx.Commit(); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to commit transaction", err.Error())
			return
		}

		msg := "Shares debited successfully"
		if newQuantity == 0 {
			msg = "Shares debited successfully, stock fully sold"
		}
		util.JSON(c, http.StatusOK, fmt.Sprintf("%s, remaining quantity: %d", msg, newQuantity), "", "")
	}
}

func GetInvestorPortfolio(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		govtID := c.Query("govt_id")
		if govtID == "" {
			util.JSON(c, http.StatusBadRequest, "", "govt_id query parameter is required", "")
			return
		}

		schema := fmt.Sprintf("investor_%s", govtID)
		var exists bool
		err := db.QueryRow("SELECT EXISTS (SELECT 1 FROM information_schema.schemata WHERE schema_name = $1)", schema).Scan(&exists)
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to check investor schema", err.Error())
			return
		}
		if !exists {
			util.JSON(c, http.StatusNotFound, "", "Investor not found", "")
			return
		}

		query := fmt.Sprintf("SELECT ticker, quantity FROM %s.holdings", schema)
		rows, err := db.Query(query)
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to fetch holdings", err.Error())
			return
		}
		defer rows.Close()

		var holdings []Holding
		for rows.Next() {
			var h Holding
			if err := rows.Scan(&h.Ticker, &h.Quantity); err != nil {
				util.JSON(c, http.StatusInternalServerError, "", "Failed to scan holding", err.Error())
				return
			}
			holdings = append(holdings, h)
		}

		if err := rows.Err(); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Error iterating holdings", err.Error())
			return
		}

		c.JSON(http.StatusOK, holdings)
	}
}
