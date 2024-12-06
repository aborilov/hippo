package pg

import (
	"github.com/aborilov/hippo/business/medication/model"
	"github.com/google/uuid"
)

type Medication struct {
	ID     uuid.UUID `db:"id"`
	Name   string    `db:"name"`
	Dosage int64     `db:"dosage"`
	Form   string    `db:"form"`
}

func (m *Medication) toService() (*model.Medication, error) {
	f, err := model.FormString(m.Form)
	if err != nil {
		return nil, err
	}
	return &model.Medication{
		ID:     m.ID,
		Name:   m.Name,
		Dosage: m.Dosage,
		Form:   f,
	}, nil
}

func fromServiceMedication(m *model.Medication) *Medication {
	return &Medication{
		ID:     m.ID,
		Name:   m.Name,
		Dosage: m.Dosage,
		Form:   m.Form.String(),
	}
}
