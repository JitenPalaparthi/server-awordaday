package models

import "fmt"

// GetFieldValidationError is to return an error message
func GetFieldValidationError(field string) error {
	return fmt.Errorf("Field:%s is a required field", field)
}
