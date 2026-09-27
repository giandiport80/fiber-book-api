package service

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"uuid"

	"github.com/giandiport80/fiber-book-api/domain"
	"github.com/giandiport80/fiber-book-api/dto"
)

type customerService struct {
	customerRepository domain.CustomerRepository
}

// Update implements [domain.CustomerService].
func (c *customerService) Update(ctx context.Context, req dto.UpdateCustomerRequest) error {
	persisted, err := c.customerRepository.FindById(ctx, req.ID)
	if err != nil {
		return err
	}

	if persisted == nil {
		return errors.New("Data customer tidak ditemukan")
	}

	persisted.Code = req.Code
	persisted.Name = req.Name
	persisted.UpdatedAt = sql.NullTime{Valid: true, Time: time.Now()}

	return c.customerRepository.Update(ctx, persisted)
}

// Create implements [domain.CustomerService].
func (c *customerService) Create(ctx context.Context, req dto.CreateCustomerRequest) error {
	customer := domain.Customer{
		ID:        uuid.New().String(),
		Name:      req.Name,
		Code:      req.Code,
		CreatedAt: sql.NullTime{Valid: true, Time: time.Now()},
	}
	return c.customerRepository.Save(ctx, &customer)
}

// Index implements [domain.CustomerService].
func (c *customerService) Index(ctx context.Context) ([]dto.CustomerData, error) {
	customers, err := c.customerRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	var customerData []dto.CustomerData
	for _, v := range customers {
		customerData = append(customerData, dto.CustomerData{
			ID:   v.ID,
			Code: v.Code,
			Name: v.Name,
		})
	}

	return customerData, nil
}

func NewCustomer(customerRepository domain.CustomerRepository) domain.CustomerService {
	return &customerService{
		customerRepository: customerRepository,
	}
}
