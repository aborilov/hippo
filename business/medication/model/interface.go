package model

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	Create(context.Context, *Medication) (*Medication, error)
	List(context.Context) ([]*Medication, error)
	Get(context.Context, uuid.UUID) (*Medication, error)
	Update(context.Context, *Medication) (*Medication, error)
	Delete(context.Context, uuid.UUID) error
}

type Repository interface {
	Create(context.Context, *Medication) (*Medication, error)
	List(context.Context) ([]*Medication, error)
	Get(context.Context, uuid.UUID) (*Medication, error)
	Update(context.Context, *Medication) (*Medication, error)
	Delete(context.Context, uuid.UUID) error
}
