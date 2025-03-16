package company

import (
	"database/sql"
	"net/http"
	"share-depository/src/util"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type company struct {
	Ticker     string `json:"ticker"`
	Face_value int    `json:"face_value"`
	Name       string `json:"company_name"`
}

// AddCompanyHandler handles adding a new company
// @Summary Add a new company in depository
// @Description Adds a new company to the platform schema
// @Tags company
// @Accept json
// @Produce json
// @Param company body company true "company Data"
// @Success 201 {object} gin.H{"message": "Company added successfully"}
// @Failure 400 {object} gin.H{"error": "Invalid request body"}
// @Failure 500 {object} gin.H{"error": "Failed to insert company"}
// @Router /api/company/add [post]
func AddCompanyHandler(db *sql.DB) gin.HandlerFunc {

	return func(c *gin.Context) {
		var comp company
		if err := c.ShouldBindJSON(&comp); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}
		er := util.SetSearchPath(db, "platform")
		if er != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to set search path to platform", "details": er.Error()})
			return
		}
		query := "INSERT INTO platform.companies (ticker, face_value, company_name) VALUES ($1, $2, $3)"
		_, err := db.Exec(query, comp.Ticker, comp.Face_value, comp.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert company", "details": err.Error()})
			return
		}

		// Return success response
		c.JSON(http.StatusCreated, gin.H{"message": "Company added successfully"})

	}
}
func GetAllCompaniesHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := util.SetSearchPath(db, "platform"); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to set search path", err.Error())
			return
		}
		rows, err := db.Query("SELECT ticker, face_value, company_name FROM companies")
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to fetch companies", err.Error())
			return
		}
		defer rows.Close()

		var companies []company
		for rows.Next() {
			var comp company
			err := rows.Scan(&comp.Ticker, &comp.Face_value, &comp.Name)
			if err != nil {
				util.JSON(c, http.StatusInternalServerError, "", "Malformed JSON : Failed to scan company data", err.Error())
				return
			}
			companies = append(companies, comp)
		}

		if err = rows.Err(); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Mlaformed JSON", err.Error())
			return
		}

		c.JSON(http.StatusOK, companies)

	}
}

// RemoveCompanyHandler deletes a company by ticker
// @Summary Delete a company
// @Description Deletes a company by ticker from the platform schema
// @Tags company
// @Accept json
// @Produce json
// @Param request body struct{Ticker string `json:"ticker"`} true "Ticker of the company to delete"
// @Success 200 {object} util.Response "Successful deletion"
// @Failure 400 {object} util.Response "Invalid request"
// @Failure 404 {object} util.Response "Company not found"
// @Failure 500 {object} util.Response "Server error"
// @Router /company/delete [post]
func RemoveCompanyHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check if db is nil
		if db == nil {
			util.JSON(c, http.StatusInternalServerError, "", "Database connection is nil", "")
			return
		}

		// Parse request
		var request struct {
			Ticker string `json:"ticker"` // Fixed JSON tag
		}
		if err := c.ShouldBindJSON(&request); err != nil {
			util.JSON(c, http.StatusBadRequest, "", "Invalid request body details", err.Error())
			return
		}

		// Set search path
		if err := util.SetSearchPath(db, "platform"); err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to set search path", err.Error())
			return
		}

		// Execute delete
		query := "DELETE FROM platform.companies WHERE ticker = $1"
		result, err := db.Exec(query, request.Ticker)
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to delete company", err.Error())
			return
		}

		// Check rows affected
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			util.JSON(c, http.StatusInternalServerError, "", "Failed to get rows affected", err.Error())
			return
		}
		if rowsAffected == 0 {
			util.JSON(c, http.StatusNotFound, "", "Company not found", "")
			return
		}

		util.JSON(c, http.StatusOK, "Company deleted successfully", "", "")
	}
}
