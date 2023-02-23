package main

import (
	// "fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	First_Name string
	Last_Name string
	Email string
	Closet Closet `gorm:"foreignKey:UserID"`
}

type Closet struct {
	gorm.Model
	Items []Item `gorm:"foreignKey:ClosetID"`
	UserID uint
}

type Item struct {
	gorm.Model
	Image string // string is temp, not sure how we will get images
	Name string
	Category Category `gorm:"foreignKey:CategoryID"`
	CategoryID uint
	Closet     Closet   `gorm:"foreignKey:ClosetID"`
	ClosetID   uint
}

type Category struct {
	gorm.Model
	Image string
	Name string
	Items []Item
}

func main() {	
	// gorm.Open() returns a database object
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}

	// Create the schema
	db.AutoMigrate(&User{}, &Closet{}, &Item{}, &Category{})

	// Create entries
	// Create a user
	user := User {
		First_Name: "Michelle",
		Last_Name: "Taing",
		Email: "mleetaing@gmail.com",
	}
	db.Create(&user)

	// Create a closet
	closet := Closet {
		UserID: user.ID,
	}
	db.Create(&closet)

	// Associate the closet with the user
	user.Closet = closet
	db.Save(&user)

}