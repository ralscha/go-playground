package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/samber/do/v2"
)

func main() {
	injector := do.New(
		CommercePackage,
		do.Eager(&struct{ BootAt time.Time }{BootAt: time.Now()}),
	)

	do.ProvideValue(injector, &AppConfig{
		Environment:    "development",
		Region:         "us-east-1",
		Currency:       "USD",
		PaymentGateway: "Stripe",
	})
	do.Provide(injector, func(i do.Injector) (*Database, error) {
		cfg := do.MustInvoke[*AppConfig](i)
		dsn := fmt.Sprintf("postgres://orders.%s.internal:5432/commerce", cfg.Region)
		return &Database{DSN: dsn}, nil
	})
	do.ProvideNamed(injector, "operator", func(do.Injector) (*User, error) {
		return &User{Name: "Aisha", Title: "operations manager"}, nil
	})
	do.ProvideNamed(injector, "auditor", func(do.Injector) (*User, error) {
		return &User{Name: "Marco", Title: "finance auditor"}, nil
	})
	do.ProvideNamedValue(injector, "region", "us-east-1")
	do.ProvideNamedValue(injector, "approver", &ApprovalPolicy{MinAmountForManualReview: 10000})
	do.ProvideNamedTransient(injector, "request-id", func(do.Injector) (string, error) {
		id := checkoutCounter.Add(1)
		return fmt.Sprintf("req-%06d", id), nil
	})

	do.OverrideValue(injector, &AppConfig{
		Environment:    "staging",
		Region:         "eu-west-1",
		Currency:       "EUR",
		PaymentGateway: "Stripe",
	})
	do.OverrideNamedValue(injector, "region", "eu-west-1")
	do.OverrideNamed(injector, "operator", func(do.Injector) (*User, error) {
		return &User{Name: "Ravi", Title: "regional operations lead"}, nil
	})

	service, err := do.Invoke[*OrderService](injector)
	if err != nil {
		panic(err)
	}

	requestID1 := do.MustInvokeNamed[string](injector, "request-id")
	requestID2 := do.MustInvokeNamed[string](injector, "request-id")
	checkout1 := do.MustInvoke[string](injector)
	checkout2 := do.MustInvoke[string](injector)

	order := OrderRequest{OrderID: "SO-2026-00421", Amount: 12999}
	result := service.PlaceOrder(order)

	fmt.Println("=== ORDER PIPELINE ===")
	fmt.Printf("Operator: %s (%s)\n", service.Operator.Name, service.Operator.Title)
	fmt.Printf("Persist: %s\n", result[0])
	fmt.Printf("Payment: %s\n", result[1])
	fmt.Printf("Named transient IDs: %s vs %s\n", requestID1, requestID2)
	fmt.Printf("Package transient sessions: %s vs %s\n", checkout1, checkout2)

	gateway := do.MustInvokeAs[PaymentGateway](injector)
	fmt.Println(gateway.Charge("SO-2026-00422", 4500, service.Currency))

	tenantScope := injector.Scope("tenant-acme-retail")
	do.ProvideNamed(tenantScope, "auditor", func(do.Injector) (*User, error) {
		return &User{Name: "Helena", Title: "tenant finance controller"}, nil
	})

	api := do.MustInvokeStruct[OrderAPI](tenantScope)
	fmt.Printf(
		"API scope region=%s, manual-review-threshold=%d, auditor=%s\n",
		api.Region,
		api.Approver.MinAmountForManualReview,
		api.Auditor.Name,
	)

	if err := do.HealthCheck[*Database](injector); err != nil {
		fmt.Printf("Database healthcheck: %v\n", err)
	} else {
		fmt.Println("Database healthcheck passed")
	}

	fmt.Println("\n=== INJECTOR MAP ===")
	injectorMap := do.ExplainInjector(injector)
	fmt.Println(injectorMap.String())
	if desc, ok := do.ExplainService[*OrderService](injector); ok {
		lines := strings.Split(desc.String(), "\n")
		if len(lines) > 8 {
			lines = lines[:8]
		}
		fmt.Println("=== ORDER SERVICE (snippet) ===")
		fmt.Println(strings.Join(lines, "\n"))
	}

	_ = do.ShutdownWithContext[*Database](context.Background(), injector)
	report := injector.Shutdown()
	fmt.Printf("\nShutdown succeed=%v, services=%d\n", report.Succeed, len(report.Services))
}
