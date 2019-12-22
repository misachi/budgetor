package main

import (
    "fmt"
    _"reflect"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
    "budgetor/models"
    repo "budgetor/repository"
)

func main() {
    db, err := gorm.Open("sqlite3", "test.db")
    if err != nil {
        fmt.Println(err)
        panic("failed to connect database")
    }
    defer db.Close()

    db.AutoMigrate(&models.User{})

    u_repo := repo.NewUserRepository(db)
    user := models.User{Name: "Carlos Misachi"}
    //data := map[string]interface{}{"Name": "Carlos Alphonse Misachi", "Phone": "01234567"}
    
    u_repo.DeleteUser(&user)
    //u_repo.AddUser(user)
    //u_repo.UpdateUser(&user, data)
    fmt.Println(u_repo.GetUser(user))
    //fmt.Println(u_repo.GetAllUsers([]models.User{user,}))
}