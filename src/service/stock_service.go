package service

import (
	"database/sql"
	"errors"
	"fmt"
)

func ValidateStockUpdateInputs(ticker string, quantityChange int) error {
	if ticker == "" {
		return errors.New("ticker cannot be empty")
	}
	if quantityChange == 0 {
		return errors.New("quantity change cannot be zero")
	}
	return nil
}

func UpdateStockQuantity(tx *sql.Tx, schema, ticker string, quantityChange int) (int, error) {
	if err := ValidateStockUpdateInputs(ticker, quantityChange); err != nil {
		return 0, err
	}

	query := fmt.Sprintf("SELECT quantity FROM %s.holdings WHERE ticker = $1 FOR UPDATE", schema)
	var currentQuantity sql.NullInt64
	err := tx.QueryRow(query, ticker).Scan(&currentQuantity)
	if err != nil && err != sql.ErrNoRows {
		return 0, fmt.Errorf("failed to lock and read current quantity: %w", err)
	}

	var newQuantity int
	if err == sql.ErrNoRows {
		// Stock doesn’t exist
		if quantityChange < 0 {
			return 0, fmt.Errorf("cannot debit %d shares of %s: stock not held", -quantityChange, ticker)
		}
		newQuantity = quantityChange // New stock, set initial quantity
	} else {
		// Stock exists
		current := int(currentQuantity.Int64)
		newQuantity = current + quantityChange
		if newQuantity < 0 {
			return 0, fmt.Errorf("cannot debit %d shares of %s: only %d shares held", -quantityChange, ticker, current)
		}
	}

	// Update the holdings table
	if newQuantity == 0 {
		// Complete sell: delete the row
		query = fmt.Sprintf("DELETE FROM %s.holdings WHERE ticker = $1", schema)
		_, err = tx.Exec(query, ticker)
		if err != nil {
			return 0, fmt.Errorf("failed to delete holdings: %w", err)
		}
	} else {
		// Insert or update
		query = fmt.Sprintf(`
            INSERT INTO %s.holdings (ticker, quantity)
            VALUES ($1, $2)
            ON CONFLICT (ticker)
            DO UPDATE SET quantity = $2`, schema)
		_, err = tx.Exec(query, ticker, newQuantity)
		if err != nil {
			return 0, fmt.Errorf("failed to update holdings: %w", err)
		}
	}

	return newQuantity, nil
}
