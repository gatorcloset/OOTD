package main

import (
	//"errors"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB

type User struct {
	gorm.Model
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
	Email      string `json:"email"`
}

func InitialMigration() {
	// Connect to database
	var err error
	db, err = gorm.Open(sqlite.Open("OOTD.db"), &gorm.Config{})

	// if error display message
	if err != nil {
		panic("failure to connect to database")
	}

	//print connected to display connection
	fmt.Println("connected")

	//create nested tables
	db.AutoMigrate(&User{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
}
