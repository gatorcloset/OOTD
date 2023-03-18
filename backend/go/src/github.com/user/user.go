package main

import (
	//"errors"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
	"github.com/google/uuid"
	"path/filepath"
	"os"
	"io"
)

var db *gorm.DB

type User struct {
	gorm.Model
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
	Username      string `json:"username"`
	Password   string `json:"password"`
}

type Item struct {
	gorm.Model
	ItemID       uint   `gorm:"primaryKey"`
    Name     string `json:"name"`
    Category string `json:"category"`
    ImagePath string `json:"image"`
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
	db.AutoMigrate(&Item{})
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []User
	db.Find(&users)
	json.NewEncoder(w).Encode(users)

}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	db.Save(&user)
	json.NewEncoder(w).Encode(user) 
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var user User
	db.First(&user, params["id"])
	db.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("The user has successfully been deleted.") 
}

func CreateItem(w http.ResponseWriter, r *http.Request) {

    // Parse the multipart form data
    if err := r.ParseMultipartForm(32 << 20); err != nil {
        http.Error(w, "error parsing multipart form data", http.StatusBadRequest)
        return
    }

    // Retrieve the image file from the form data
    file, handler, err := r.FormFile("image")
    if err != nil {
        http.Error(w, "image file not found", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Generate a unique file name for the image
    fileName := uuid.New().String() + filepath.Ext(handler.Filename)

    // Save the image file to a local directory
    imagePath := "images/" + fileName
    dst, err := os.Create(imagePath)
    if err != nil {
        http.Error(w, "error saving image file", http.StatusInternalServerError)
        return
    }
    defer dst.Close()

    if _, err := io.Copy(dst, file); err != nil {
        http.Error(w, "error saving image file", http.StatusInternalServerError)
        return
    }

    // Create a new item object and save it to the database
    item := Item{Name: r.FormValue("name"), Category: r.FormValue("category"), ImagePath: imagePath}
    result := db.Create(&item)
    if result.Error != nil {
        http.Error(w, "error creating item", http.StatusInternalServerError)
        return
    }

    // Return the created item as JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(item)
}

