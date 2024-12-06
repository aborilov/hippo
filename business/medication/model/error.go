package model

import "fmt"

type ErrNotFound struct {
	MedicationID string
}

func (e ErrNotFound) Error() string {
	return fmt.Sprintf("medication not found (ID: %s)", e.MedicationID)
}
