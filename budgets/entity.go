package budgets

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"

	"budgetor/expenses"
	"budgetor/utils"
)

type BudgetPeriod struct {
	gorm.Model
	Start      *time.Time
	End        *time.Time
	Isactive   bool               `gorm:"default:'true'"`
	Periodtype string             `gorm:"type:varchar(100)"`
	Budgets    []Budget           `gorm:"foreignkey:Id"`
	Expenses   []expenses.Expense `gorm:"foreignkey:Id"`
}

func (P *BudgetPeriod) BeforeCreate() (err error) {
	errored := false
	for _, field := range utils.Period {
		if P.Periodtype == field {
			errored = true
		}
	}

	if errored == false {
		err = errors.New("Field Name not allowed")
	}
	return
}

type Budget struct {
	gorm.Model
	Amount float64
}
