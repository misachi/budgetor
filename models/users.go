package models

import (
    "time"

    "github.com/jinzhu/gorm"
)


type User struct {
    gorm.Model
    Name            string         `gorm:"type:varchar(100)"`
    Dob             time.Time
    Email           *string         `gorm:"type:varchar(250)"`
    Phone           string         `gorm:"type:varchar(15)"`
    BudgetPeriods   []BudgetPeriod    `gorm:"foreignkey:Id"`
}