package services

func SeedAdminData(admin *Admin) {
	admin.AddTenant("Tenant A")
	admin.AddTenant("Tenant B")
	admin.AddTenant("Tenant C")

	admin.RecordTransaction(1, 100.0)
	admin.RecordTransaction(2, 150.0)
	admin.RecordTransaction(3, 200.0)
	admin.RecordTransaction(1, 50.0)
	admin.RecordTransaction(2, 100.0)
}
