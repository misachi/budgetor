package main

import (
	"encoding/json"
	"fmt"
	_ "reflect"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"budgetor/budgets"
	"budgetor/users"
	"budgetor/utils"
	_ "budgetor/utils"
)

func main() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	db.AutoMigrate(&users.User{})
	db.AutoMigrate(&budgets.BudgetPeriod{})

	uRepo := users.NewUserRepository(db)
	user := users.User{Name: "Carlos Misachi"}
	user.ID = 99

	bp := budgets.BudgetPeriod{Periodtype: utils.Daily}

	uRepo.AddBudgetPeriod(&user, &bp)

	objs, _ := uRepo.GetUser(user)

	val, err := json.MarshalIndent(objs, "", " ")
	fmt.Println(string(val))
}
