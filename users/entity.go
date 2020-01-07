package users

import (
	"time"

	"github.com/jinzhu/gorm"

	bg "budgetor/budgets"
)

type User struct {
	gorm.Model
	Name          string            `gorm:"type:varchar(100)" json:"name"`
	Dob           time.Time         `json:"dob"`
	Email         *string           `gorm:"type:varchar(250)" json:"email"`
	Phone         string            `gorm:"type:varchar(15)" json:"phone"`
	BudgetPeriods []bg.BudgetPeriod `gorm:"foreignkey:FK"`
}

func (u *User) DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&User{})
}
