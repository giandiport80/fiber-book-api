package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/giandiport80/fiber-book-api/domain"
)

type customerRepository struct {
	db *goqu.Database
}

func NewCustomer(con *sql.DB) domain.CustomerRepository {
	return &customerRepository{
		db: goqu.New("default", con),
	}
}

// FindAll implements [domain.CustomerRepository].
func (c *customerRepository) FindAll(ctx context.Context) (result []domain.Customer, err error) {
	dataset := c.db.From("customers").Where(goqu.C("deleted_at").IsNull())
	err = dataset.ScanStructsContext(ctx, &result)
	return
}

// FindById implements [domain.CustomerRepository].
func (c *customerRepository) FindById(ctx context.Context, id string) (*domain.Customer, error) {
	var customer domain.Customer
	dataset := c.db.From("customers").Where(
		goqu.C("deleted_at").IsNull(),
		goqu.C("id").Eq(id),
	)

	found, err := dataset.ScanStructContext(ctx, &customer)
	if err != nil {
		return nil, err
	}

	if !found {
		return nil, nil
	}

	return &customer, nil
}

// Save implements [domain.CustomerRepository].
func (c *customerRepository) Save(ctx context.Context, cr *domain.Customer) error {
	executor := c.db.Insert("customers").Rows(cr).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Update implements [domain.CustomerRepository].
func (c *customerRepository) Update(ctx context.Context, cr *domain.Customer) error {
	executor := c.db.Update("customers").Where(goqu.C("id").Eq(cr.ID)).Set(cr).Executor()
	_, err := executor.ExecContext(ctx)
	return err
}

// Delete implements [domain.CustomerRepository].
func (c *customerRepository) Delete(ctx context.Context, id string) error {
	executor := c.db.Update("customers").
		Where(goqu.C("id").Eq(id)).
		Set(goqu.Record{"deleted_at": sql.NullTime{Valid: true, Time: time.Now()}}).
		Executor()

	_, err := executor.ExecContext(ctx)
	return err
}
