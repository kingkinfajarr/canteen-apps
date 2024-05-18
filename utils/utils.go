package utils

import (
	"canteen-apps/services"
	"fmt"
)

func ShowMenu(admin *services.Admin) {
	var choice int

	for {
		fmt.Println("\n=============== Welcome to Canteen Tel-U ===============:")
		fmt.Println("1. Add Tenant")
		fmt.Println("2. Update Tenant")
		fmt.Println("3. Delete Tenant")
		fmt.Println("4. Record Transaction")
		fmt.Println("5. List Tenants")
		fmt.Println("6. Exit")
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			fmt.Print("Enter tenant name: ")
			fmt.Scan(&name)
			admin.AddTenant(name)
		case 2:
			var id int
			var name string
			fmt.Print("Enter tenant ID to update: ")
			fmt.Scan(&id)
			fmt.Print("Enter new tenant name: ")
			fmt.Scan(&name)
			admin.UpdateTenant(id, name)
		case 3:
			var id int
			fmt.Print("Enter tenant ID to delete: ")
			fmt.Scan(&id)
			admin.DeleteTenant(id)
		case 4:
			var id int
			var amount float64
			fmt.Print("Enter tenant ID for transaction: ")
			fmt.Scan(&id)
			fmt.Print("Enter transaction amount: ")
			fmt.Scan(&amount)
			admin.RecordTransaction(id, amount)
		case 5:
			admin.ListTenants()
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
}
