package main

type AppConfig struct {
	Environment    string
	Region         string
	Currency       string
	PaymentGateway string
}

type User struct {
	Name  string
	Title string
}

type ApprovalPolicy struct {
	MinAmountForManualReview int
}

type OrderRequest struct {
	OrderID string
	Amount  int
}

type OrderAPI struct {
	Service  *OrderService   `do:""`
	Approver *ApprovalPolicy `do:"approver"`
	Region   string          `do:"region"`
	Auditor  *User           `do:"auditor"`
}
