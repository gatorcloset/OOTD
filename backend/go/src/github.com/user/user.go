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
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"

)

var db *gorm.DB
var store = sessions.NewCookieStore([]byte("whoa-its-a-secret-key"))

type User struct {
	gorm.Model
	First_Name string `json:"firstname"`
	Last_Name  string `json:"lastname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type Item struct {
	gorm.Model
	UserID uint `json:"user_id"`
    Name string `json:"name"`
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
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

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

func LoginUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var loginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		http.Error(w, "Error processing login request", http.StatusBadRequest)
		return
	}

	// Find the user with the given username
	var user User
	if err := db.Where("username = ?", loginRequest.Username).First(&user).Error; err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	if !CheckPasswordHash(loginRequest.Password, user.Password) {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	session, err := store.Get(r, "session-name")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Set session values
    session.Values["authenticated"] = true
    session.Values["username"] = loginRequest.Username
	session.Values["userID"] = user.ID

    // Save the session
    err = session.Save(r, w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	json.NewEncoder(w).Encode(user)
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

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	// Get the session
    session, err := store.Get(r, "session-name")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Clear session values
    session.Values["authenticated"] = false
    session.Values["username"] = ""

    // Save the session
    err = session.Save(r, w)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
