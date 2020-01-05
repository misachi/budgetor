package budgets

import (
	"github.com/jinzhu/gorm"
)

type ExpenseRepository struct {
	conn *gorm.DB
}

func NewExpenseRepo(c *gorm.DB) *ExpenseRepository {
	return &ExpenseRepository{conn: c}
}

type BudgetPeriodRepository struct {
	conn *gorm.DB
}

func NewBudgetPeriodRepo(c *gorm.DB) *BudgetPeriodRepository {
	return &BudgetPeriodRepository{conn: c}
}

func (B *BudgetPeriodRepository) Items(items []BudgetPeriod) []BudgetPeriod {
	B.conn.Find(&items)
	return items
}

func (B *BudgetPeriodRepository) Item(item BudgetPeriod) interface{} {
	B.conn.First(&item)
	// Ensure user value is returned only if the ID exists i.e there is
	// at least one user object in the db
	if item.ID < 1 {
		return []int{}
	}
	return item
}
