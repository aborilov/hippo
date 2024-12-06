package model

import "github.com/google/uuid"

type Form uint8

const (
	FormTablet Form = 1 + iota
	FormCapsule
	FormLiquid
)

//go:generate enumer -type=Form -trimprefix=Form -text -json -sql -transform=snake -output=enum_form_gen.go

type Medication struct {
	ID     uuid.UUID
	Name   string
	Dosage int64
	Form   Form
}
