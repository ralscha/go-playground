package main

import (
	"context"
	"fmt"

	"github.com/samber/do/v2"
)

type Database struct {
	DSN string
}

func (d *Database) HealthCheck() error {
	if d.DSN == "" {
		return fmt.Errorf("database DSN is empty")
	}
	return nil
}

func (d *Database) Shutdown(ctx context.Context) error {
	fmt.Printf("Closing database connections for DSN=%s\n", d.DSN)
	return nil
}

type OrderRepository struct {
	DB *Database
}

func NewOrderRepository(i do.Injector) (*OrderRepository, error) {
	return &OrderRepository{DB: do.MustInvoke[*Database](i)}, nil
}

func (r *OrderRepository) Save(order OrderRequest) string {
	return fmt.Sprintf("order %s saved to %s", order.OrderID, r.DB.DSN)
}
