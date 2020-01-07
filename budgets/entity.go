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
	FK         uint
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

func (u *BudgetPeriod) DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&BudgetPeriod{})
}

type Budget struct {
	gorm.Model
	Amount float64
}

func (u *Budget) DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&Budget{})
}
