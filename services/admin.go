package services

import (
	"fmt"
	"sort"

	"canteen-apps/models"
)

type Admin struct {
	Tenants      []models.Tenant
	NextTenantID int
	TotalRevenue float64
}

func NewAdmin() *Admin {
	return &Admin{
		Tenants:      []models.Tenant{},
		NextTenantID: 1,
	}
}

func (a *Admin) AddTenant(name string) {
	tenant := models.Tenant{
		ID:   a.NextTenantID,
		Name: name,
	}
	a.Tenants = append(a.Tenants, tenant)
	a.NextTenantID++
	fmt.Println("Tenant added:", tenant)
}

func (a *Admin) UpdateTenant(id int, name string) {
	for i, tenant := range a.Tenants {
		if tenant.ID == id {
			a.Tenants[i].Name = name
			fmt.Println("Tenant updated:", a.Tenants[i])
			return
		}
	}
	fmt.Println("Tenant not found")
}

func (a *Admin) DeleteTenant(id int) {
	for i, tenant := range a.Tenants {
		if tenant.ID == id {
			a.Tenants = append(a.Tenants[:i], a.Tenants[i+1:]...)
			fmt.Println("Tenant deleted:", tenant)
			return
		}
	}
	fmt.Println("Tenant not found")
}

func (a *Admin) RecordTransaction(id int, amount float64) {
	for i, tenant := range a.Tenants {
		if tenant.ID == id {
			tenant.TotalAmount += amount
			tenant.TotalTransactions++
			a.Tenants[i] = tenant
			adminShare := amount * 0.25
			a.TotalRevenue += adminShare
			fmt.Printf("Transaction recorded: Tenant %d receives %.2f, Admin receives %.2f\n", tenant.ID, amount-adminShare, adminShare)
			return
		}
	}
	fmt.Println("Tenant not found")
}

func (a *Admin) ListTenants() {
	sort.Slice(a.Tenants, func(i, j int) bool {
		return a.Tenants[i].TotalTransactions > a.Tenants[j].TotalTransactions
	})

	fmt.Println("Tenant list sorted by number of transactions:")
	for _, tenant := range a.Tenants {
		fmt.Printf("ID: %d, Name: %s, Transactions: %d, Total Amount: %.2f\n",
			tenant.ID, tenant.Name, tenant.TotalTransactions, tenant.TotalAmount)
	}
}
