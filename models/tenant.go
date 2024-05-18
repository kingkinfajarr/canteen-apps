package models

type Tenant struct {
	ID                int
	Name              string
	TotalAmount       float64
	TotalTransactions int
}
