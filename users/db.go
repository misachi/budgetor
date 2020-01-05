package users

import (
	"github.com/jinzhu/gorm"

	"budgetor/budgets"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewUserRepository(c *gorm.DB) *UserRepository {
	return &UserRepository{conn: c}
}

func (R *UserRepository) GetAllUsers(u []User) []User {
	R.conn.Preload("BudgetPeriods").Find(&u)
	return u
}

func (R *UserRepository) GetUser(u User) (User, bool) {
	h := R.conn.Preload("BudgetPeriods").First(&u, u.ID)
	// Check if user exists, if not return empty value
	if h.RowsAffected < 1 {
		return User{}, false
	}
	return u, true
}

func (R *UserRepository) AddUser(u User) {
	R.conn.Create(&u)
}

func (R *UserRepository) UpdateUser(u *User, data map[string]interface{}) {
	R.conn.Model(&u).Updates(data)
}

func (R *UserRepository) DeleteUser(u *User) {
	R.conn.Delete(&u)
}

// AddBudgetPeriod creates a new BudgetPeriod and updates the User
// data
func (R *UserRepository) AddBudgetPeriod(user *User, item *budgets.BudgetPeriod) {
	// Association seems to create a new user record if it does not exist in the
	// database, hence, this check. Also we don't want to create a BudgetPeriod if user
	// does not exist
	if _, boold := R.GetUser(*user); boold {
		R.conn.Create(&item)
		R.conn.Model(&user).Association("BudgetPeriods").Append([]budgets.BudgetPeriod{*item})
	}
}
