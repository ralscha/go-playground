package main

import "github.com/samber/do/v2"

var CommercePackage = do.Package(
	do.Lazy(NewOrderService),
	do.Lazy(NewStripeGateway),
	do.LazyNamed("primary-repo", NewOrderRepository),
	do.Transient(NewCheckoutSession),
	do.Bind[*StripeGateway, PaymentGateway](),
)
