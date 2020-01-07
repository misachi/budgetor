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
var app = []interface{}{
	&users.User{},
	&expenses.Expense{},
	&expenses.ExpenseLine{},
	&budgets.Budget{},
	&budgets.BudgetPeriod{},
}

// Make migrations for all project apps
func migrate(db *gorm.DB) {
	for _, elem := range app {
		db.AutoMigrate(elem)
	}
}

func main() {
	db := utils.Conn
	defer db.Close()
	migrate(db)
	log.Fatal(http.ListenAndServe(":8080", api.Handlers(db)))
}
