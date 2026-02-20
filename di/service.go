package main

import (
	"fmt"
	"sync/atomic"

	"github.com/samber/do/v2"
)

type OrderService struct {
	Repo     *OrderRepository
	Gateway  PaymentGateway
	Currency string
	Operator *User
}

func NewOrderService(i do.Injector) (*OrderService, error) {
	repo := do.MustInvokeNamed[*OrderRepository](i, "primary-repo")
	gateway := do.MustInvokeAs[PaymentGateway](i)
	cfg := do.MustInvoke[*AppConfig](i)
	operator := do.MustInvokeNamed[*User](i, "operator")

	return &OrderService{
		Repo:     repo,
		Gateway:  gateway,
		Currency: cfg.Currency,
		Operator: operator,
	}, nil
}

func (s *OrderService) PlaceOrder(order OrderRequest) []string {
	results := make([]string, 0, 2)
	results = append(results, s.Repo.Save(order))
	results = append(results, s.Gateway.Charge(order.OrderID, order.Amount, s.Currency))
	return results
}

var checkoutCounter atomic.Int64

func NewCheckoutSession(do.Injector) (string, error) {
	id := checkoutCounter.Add(1)
	return fmt.Sprintf("checkout-session-%06d", id), nil
}
