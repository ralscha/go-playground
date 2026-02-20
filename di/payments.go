package main

import (
	"fmt"

	"github.com/samber/do/v2"
)

type PaymentGateway interface {
	Charge(orderID string, amount int, currency string) string
}

type StripeGateway struct {
	Provider string
}

func NewStripeGateway(i do.Injector) (*StripeGateway, error) {
	cfg := do.MustInvoke[*AppConfig](i)
	return &StripeGateway{Provider: cfg.PaymentGateway}, nil
}

func (g *StripeGateway) Charge(orderID string, amount int, currency string) string {
	return fmt.Sprintf("payment authorized via %s for order=%s amount=%d %s", g.Provider, orderID, amount, currency)
}
