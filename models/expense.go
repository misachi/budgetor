package models

import (
    "time"
    "errors"
    "github.com/jinzhu/gorm"
)

// Expense Categories
const (
    Food = "Food"
    Entertainment = "Entertainment"
    Airtime = "Airtime"
    Shopping = "Shopping"
    Transport = "Transport"
    Family = "Family"
    Utils = "Utils"  // house utilities(electricity, internet, water etc)
    Rent = "Rent"
    
    Daily = "Daily"
    Weekly = "Weekly"
    Monthly = "Monthly"
    Yearly = "Yearly"
)
var Category = [7]string{Food, Entertainment, Airtime, Shopping, Transport, Family, Utils}
var Period = [4]string{Daily, Weekly, Monthly, Yearly}

type Expense struct {
    gorm.Model
    Name            string      `gorm:"type:varchar(100);"`
    BudgetLines   []BudgetLine  `gorm:"foreignkey:Id"`
}

// Ensure only allowed Category Name is saved. Could not find
// Gorm implementation for choices for fields
func (E *Expense) BeforeCreate() (err error) {
    errored := false
    for _, cat := range Category {
        if E.Name == cat {
            errored = true
        }
    }

    if errored == false {
        err = errors.New("Field Name not allowed")
    }
    return
}

type BudgetPeriod struct {
    gorm.Model
    Start           *time.Time
    End             *time.Time
    Isactive        bool        `gorm:"default:'true'"`
    Periodtype      string      `gorm:"type:varchar(100)"`
    Budgets         Budget      `gorm:"foreignkey:Id"`
    Expenses        []Expense   `gorm:"foreignkey:Id"`
}

func (P *BudgetPeriod) BeforeCreate() (err error) {
    errored := false
    for _, field := range Period {
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
    Amount          float64
}

type BudgetLine struct {
    gorm.Model
    Item            string      `gorm:"type:varchar(20)"`
    Description     *string      `gorm:type:text`
    Amount          float64
}