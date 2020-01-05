package users

import "budgetor/budgets"

/*
UserRepo shoud be
**/
type UserRepo interface {
	GetAllUsers(u []User) []User
	GetUser(u User) interface{}
	AddUser(u User)
	UpdateUser(u *User, data map[string]interface{})
	DeleteUser(u *User)
	AddBudgetPeriod(user *User, item *budgets.BudgetPeriod)
}
