package repository

import (
    "github.com/jinzhu/gorm"
    "budgetor/models"
)

type UserInterface interface {
    GetAllUsers(u []models.User) ([]models.User)
    GetUser(u models.User) interface{}
    AddUser(u models.User)
    UpdateUser(u *models.User, data map[string]interface{})
    DeleteUser(u *models.User)
}

type UserRepository struct {
    conn *gorm.DB
}

func NewUserRepository(c *gorm.DB) *UserRepository {
    return &UserRepository{conn: c}
}

func (R *UserRepository) GetAllUsers(u []models.User) ([]models.User) {
    //users := []models.User{}
    R.conn.Find(&u)
    return u
}

func (R *UserRepository) GetUser(u models.User) interface{} {
    R.conn.First(&u)

    // Ensure user value is returned only if the ID exists i.e there is
    // at least one user object in the db
    if u.ID < 1 {
        return []int{}
    }
    return u
}

func (R *UserRepository) AddUser(u models.User) {
    R.conn.Create(&u)
}

func (R *UserRepository) UpdateUser(u *models.User, data map[string]interface{}) {
    R.conn.Model(&u).Update(data)
}

func (R *UserRepository) DeleteUser(u *models.User) {
    R.conn.Delete(&u)
}