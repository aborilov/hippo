package pg

import (
	"context"
	"errors"
	"fmt"

	"github.com/aborilov/hippo/business/medication/model"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

const (
	table = "medication"
)

func NewRepository(db *sqlx.DB) (model.Repository, error) {
	if db == nil {
		return nil, errors.New(`"db" cannot be nil`)
	}

	r := &repository{
		db: db,
		gq: goqu.New("postgres", db),
	}
	return r, nil
}

type repository struct {
	db *sqlx.DB
	gq *goqu.Database
}

func (repo *repository) Create(ctx context.Context, m *model.Medication) (*model.Medication, error) {
	rec := fromServiceMedication(m)
	if _, err := repo.gq.Insert(table).Rows(rec).Executor().ExecContext(ctx); err != nil {
		return nil, err
	}
	return repo.Get(ctx, m.ID)
}
func (repo *repository) List(ctx context.Context) ([]*model.Medication, error) {
	recs := []Medication{}
	err := repo.gq.From(table).ScanStructsContext(ctx, &recs)
	if err != nil {
		return nil, err
	}
	var meds []*model.Medication
	for _, r := range recs {
		s, err := r.toService()
		if err != nil {
			return nil, err
		}
		meds = append(meds, s)
	}

	return meds, nil
}
func (repo *repository) Get(ctx context.Context, id uuid.UUID) (*model.Medication, error) {
	record := &Medication{}
	found, err := repo.gq.From(table).Where(goqu.I("id").Eq(id.String())).ScanStructContext(ctx, record)
	if err != nil {
		return nil, fmt.Errorf("unable to get medication: %w", err)
	}
	if !found {
		return nil, model.ErrNotFound{MedicationID: id.String()}
	}
	t, err := record.toService()
	if err != nil {
		return nil, fmt.Errorf("unable to parse medication from db: %w", err)
	}
	return t, nil
}
func (repo *repository) Update(ctx context.Context, m *model.Medication) (*model.Medication, error) {
	record := fromServiceMedication(m)
	_, err := repo.gq.Update(table).Where(goqu.I("id").Eq(record.ID)).Set(record).Executor().ExecContext(ctx)
	if err != nil {
		return nil, err
	}
	return repo.Get(ctx, m.ID)
}
func (repo *repository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := repo.gq.Delete(table).Where(goqu.I("id").Eq(id.String())).Executor().ExecContext(ctx)
	return err
}
