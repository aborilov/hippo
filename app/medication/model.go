package medication

import "github.com/aborilov/hippo/business/medication/model"

type Medication struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Dosage int64  `json:"dosage"`
	Form   string `json:"form"`
}

func (m *Medication) ToService() (*model.Medication, error) {
	f, err := model.FormString(m.Form)
	if err != nil {
		return nil, err
	}
	return &model.Medication{
		Name:   m.Name,
		Dosage: m.Dosage,
		Form:   f,
	}, nil
}

func serviceToMedication(m *model.Medication) *Medication {
	return &Medication{
		ID:     m.ID.String(),
		Name:   m.Name,
		Dosage: m.Dosage,
		Form:   m.Form.String(),
	}
}
