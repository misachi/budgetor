package main

import (
	"log"
	"net/http"

	"github.com/jinzhu/gorm"

	"budgetor/api"
	"budgetor/budgets"
	"budgetor/expenses"
	"budgetor/users"
	"budgetor/utils"
)

// Container for all project apps
var app = []utils.Entity{
	&users.User{},
	&expenses.Expense{},
	&expenses.ExpenseLine{},
	&budgets.Budget{},
	&budgets.BudgetPeriod{},
}

// Make migrations for all project apps
func migrate(db *gorm.DB) {
	for _, elem := range app {
		elem.DBMigrate(db)
	}
}

func main() {
	migrate(utils.Conn)
	log.Fatal(http.ListenAndServe(":8080", api.Handlers()))
}
