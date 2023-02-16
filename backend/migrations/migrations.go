package migrations

import (
    "github.com/jinzhu/gorm"
    "github.com/megan-shah1/backend/helpers"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
    gorm.Model
    Username string
    Email    string
    Password string
}

type Account struct {
    gorm.Model
    Type   string
    Name   string
    UserID uint
}

func connectDB() *gorm.DB {
    
    db, err := gorm.Open("postgres", "host=127.0.0.1 port=5432 user=postgres dbname=closetapp password=postgres sslmode=disable")
    helpers.HandleErr(err)
    return db
}

func createAccounts() {
    db := connectDB()

    users := [2]User{
        {Username: "Marty", Email: "martin@martin.com"},
        {Username: "Megan", Email: "megan@megan.com"},
    }

    for i := 0; i < len(users); i++ {
        generatedPassword := helpers.HashAndSalt([]byte(users[i].Username))
        user := User{Username: users[i].Username, Email: users[i].Email, Password: generatedPassword}
        db.Create(&user)

        account := Account{Type: "Personal Account", Name: string(users[i].Username) + "'s" + "Closet", UserID: user.ID}
        db.Create(&account)
    }
    defer db.Close()
}

func Migrate() {
    db := connectDB()
    db.AutoMigrate(&User{}, &Account{})
    defer db.Close()

    createAccounts()
}

