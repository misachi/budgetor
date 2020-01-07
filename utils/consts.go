package utils

// Expense Categories
const (
	Food          = "Food"
	Entertainment = "Entertainment"
	Airtime       = "Airtime"
	Shopping      = "Shopping"
	Transport     = "Transport"
	Family        = "Family"
	Utility       = "Utility" // house utilities(electricity, internet, water etc)
	Rent          = "Rent"

	Daily   = "Daily"
	Weekly  = "Weekly"
	Monthly = "Monthly"
	Yearly  = "Yearly"
)

var Category = [7]string{Food, Entertainment, Airtime, Shopping, Transport, Family, Utility}
var Period = [4]string{Daily, Weekly, Monthly, Yearly}
