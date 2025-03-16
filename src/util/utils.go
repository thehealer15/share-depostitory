package util

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SetSearchPath sets the database search path to the specified schema.
func SetSearchPath(db *sql.DB, schema string) error {
	query := fmt.Sprintf("SET search_path TO %s", schema)
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to set search path to %s: %w", schema, err)
	}
	return nil
}

// Response represents a standard JSON response structure
type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
	Details string `json:"details,omitempty"`
}

// JSON sends a standardized JSON response using Gin's c.JSON
func JSON(c *gin.Context, status int, message, errMsg, details string) {
	resp := Response{}

	switch {
	case status >= http.StatusBadRequest && status < http.StatusInternalServerError:
		// Client errors (400-499)
		resp.Error = errMsg
		if details != "" {
			resp.Details = details
		}
	case status >= http.StatusInternalServerError:
		// Server errors (500+)
		resp.Error = errMsg
		if details != "" {
			resp.Details = details
		}
	default:
		// Success cases (200-399)
		resp.Message = message
	}

	c.JSON(status, resp)
}
