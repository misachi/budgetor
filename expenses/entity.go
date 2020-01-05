package expenses

import (
	"errors"

	"github.com/jinzhu/gorm"

	"budgetor/utils"
)

type Expense struct {
	gorm.Model
	Name        string        `gorm:"type:varchar(100);"`
	BudgetLines []ExpenseLine `gorm:"foreignkey:Id"`
}

// Ensure only allowed Category Name is saved. Could not find
// Gorm implementation for choices for fields
func (E *Expense) BeforeCreate() (err error) {
	errored := false
	for _, cat := range utils.Category {
		if E.Name == cat {
			errored = true
		}
	}

	if errored == false {
		err = errors.New("Field Name not allowed")
	}
	return
}

type ExpenseLine struct {
	gorm.Model
	Item        string  `gorm:"type:varchar(20)"`
	Description *string `gorm:type:text`
	Amount      float64
}
