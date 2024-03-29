package main

import (
	//"errors"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/mattn/go-sqlite3"
	"golang.org/x/crypto/bcrypt"
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
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	ImagePath string `json:"image"`
}

type Tag struct {
	gorm.Model
	TagName string `json:"tagname"`
}

type ItemTag struct {
	gorm.Model
	ItemID uint `json:"item_ID"`
	TagID  uint `json:"tag_ID"`
}

type Outfit struct {
	gorm.Model
	UserID        uint  `json:"user_id"`
	Name		  string `json:"Name"`
	Tops          Item `gorm:"foreignKey:TopID"`
	TopID         uint
	Bottoms       Item `gorm:"foreignKey:BottomID"`
	BottomID      uint
	OnePieces     Item `gorm:"foreignKey:OnePieceID"`
	OnePieceID    uint
	Accessories   Item `gorm:"foreignKey:AccessoriesID"`
	AccessoriesID uint
	Shoes         Item `gorm:"foreignKey:ShoesID"`
	ShoesID       uint
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
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&ItemTag{})
	db.AutoMigrate(&Outfit{})
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

	var existingUser User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err == nil {
		http.Error(w, "This username has already been taken. Choose another username", http.StatusConflict)
		return
	}

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

	// Save the session
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func getUserData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, _ := store.Get(r, "session-name")

	authenticated := session.Values["authenticated"]
	username := session.Values["username"]
	userID := session.Values["userID"]

	// Return the session values in the response
	data := map[string]interface{}{
		"authenticated": authenticated,
		"username":      username,
		"userID":        userID,
	}

	json.NewEncoder(w).Encode(data)
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
	// imagePath := "images/" + fileName
	imagePath := "../../../../../src/assets/item-images/" + fileName
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

	// Get the value of "id" from the form
	idStr := r.FormValue("id")
	// Convert the "id" value from string to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		// Handle the error
	}

	// Create a new item object and save it to the database
	item := Item{UserID: uint(id), Name: r.FormValue("name"), Category: r.FormValue("category"), ImagePath: imagePath[19:]}
	result := db.Create(&item)
	if result.Error != nil {
		http.Error(w, "error creating item", http.StatusInternalServerError)
		return
	}

	// Return the created item as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	// Retrieve the item from the database
	var item Item
	db.First(&item, id)

	// Return the retrieved item as JSON
	json.NewEncoder(w).Encode(item)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	// Retrieve all tags from the database
	w.Header().Set("Content-Type", "application/json")
	var items []Item
	result := db.Find(&items)
	if result.Error != nil {
		http.Error(w, "error retrieving tags", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(items)
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// Parse the multipart form data
	if err := r.ParseMultipartForm(32 << 20); err != nil {
		http.Error(w, "error parsing multipart form data", http.StatusBadRequest)
		return
	}

	// Retrieve the image file from the form data
	file, handler, err := r.FormFile("image")
	if err != nil && err != http.ErrMissingFile {
		// Error occurred while retrieving image file
		http.Error(w, "error retrieving image file", http.StatusBadRequest)
		return
	}

	// Get the value of "id" from the form
	idStr := r.FormValue("id")
	// Convert the "id" value from string to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		http.Error(w, "error parsing item ID", http.StatusBadRequest)
		return
	}

	// Find the existing item in the database
	item := Item{}
	if err := db.First(&item, uint(id)).Error; err != nil {
		http.Error(w, "item not found", http.StatusNotFound)
		return
	}

	// Update the item's fields
	item.Name = r.FormValue("name")
	if category := r.FormValue("category"); category != "undefined" {
		item.Category = category
	}

	// If a new image file is uploaded, update the image file
	if file != nil && err != http.ErrMissingFile {
		// Generate a unique file name for the image
		fileName := uuid.New().String() + filepath.Ext(handler.Filename)

		// Save the new image file to a local directory
		// imagePath := "images/" + fileName
		imagePath := "../../../../../src/assets/item-images/" + fileName
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

		// Update the item's image path
		item.ImagePath = imagePath[19:]

		file.Close()
	}

	// Update the item in the database
	result := db.Save(&item)
	if result.Error != nil {
		http.Error(w, "error updating item", http.StatusInternalServerError)
		return
	}

	// Return the updated item as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var item Item
	db.First(&item, params["id"])
	db.Delete(&item, params["id"])
	json.NewEncoder(w).Encode("This item has successfully been deleted.")
}

func CreateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var tag Tag
	err := json.NewDecoder(r.Body).Decode(&tag)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	db.Create(&tag)
	json.NewEncoder(w).Encode(tag)
}

func GetTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]

	// Retrieve the tag from the database
	var tag Tag
	db.First(&tag, id)
	json.NewEncoder(w).Encode(tag)
}

func GetTags(w http.ResponseWriter, r *http.Request) {
	// Retrieve all tags from the database
	w.Header().Set("Content-Type", "application/json")
	var tags []Tag
	result := db.Find(&tags)
	if result.Error != nil {
		http.Error(w, "error retrieving tags", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tags)
}

func UpdateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tag Tag
	db.First(&tag, params["id"])
	json.NewDecoder(r.Body).Decode(&tag)
	db.Save(&tag)
	json.NewEncoder(w).Encode(tag)
}

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var tag Tag
	db.First(&tag, params["id"])
	db.Delete(&tag, params["id"])
	json.NewEncoder(w).Encode("This tag has successfully been deleted.")
}

func CreateItemTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var itemTag ItemTag
	json.NewDecoder(r.Body).Decode(&itemTag)

	// Check if the item and tag exist in the database
	var item Item
	result := db.First(&item, itemTag.ItemID)
	if result.Error != nil {
		http.Error(w, "item not found", http.StatusBadRequest)
		return
	}

	var tag Tag
	result = db.First(&tag, itemTag.TagID)
	if result.Error != nil {
		http.Error(w, "tag not found", http.StatusBadRequest)
		return
	}

	// Create the item tag object and save it to the database
	itemTag.ItemID = item.ID
	itemTag.TagID = tag.ID
	db.Create(&itemTag)

	// Return the created item tag as JSON
	json.NewEncoder(w).Encode(itemTag)
}

// old create outfit code
/*func CreateOutfit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var items []Item
	json.NewDecoder(r.Body).Decode(&items)

	// Create new outfit and assign items
	outfit := Outfit{}

	for i := 0; i < len(items); i++ {
		if items[i].Category == "tops" {
			outfit.Tops = items[i]
		} else if items[i].Category == "bottoms" {
			outfit.Bottoms = items[i]
		}else if items[i].Category == "one-pieces" {
			outfit.OnePieces = items[i]
		} else if items[i].Category == "accessories" {
			outfit.Accessories = items[i]
		} else if items[i].Category == "shoes" {
			outfit.Shoes = items[i]
		}
	}

	db.Create(&outfit)
	json.NewEncoder(w).Encode(outfit)
}
*/

func CreateOutfit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newOutfit Outfit
	err := json.NewDecoder(r.Body).Decode(&newOutfit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db.Create(&newOutfit)

	json.NewEncoder(w).Encode(newOutfit)
}

func UpdateOutfit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	var existingOutfit Outfit
	var updatedOutfit Outfit
	db.Preload("Tops").Preload("Bottoms").Preload("OnePieces").Preload("Accessories").Preload("Shoes").First(&existingOutfit, params["id"])

	if existingOutfit.ID == 0 {
		http.Error(w, "Outfit not found", http.StatusNotFound)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updatedOutfit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedOutfit.Name != "" {
		existingOutfit.Name = updatedOutfit.Name
	}
	if updatedOutfit.TopID != 0 {
		var top Item
		if err := db.First(&top, updatedOutfit.TopID).Error; err != nil {
			http.Error(w, "Top item not found", http.StatusNotFound)
			return
		}
		existingOutfit.Tops = top
		existingOutfit.TopID = top.ID
	}
	if updatedOutfit.BottomID != 0 {
		var bottom Item
		if err := db.First(&bottom, updatedOutfit.BottomID).Error; err != nil {
			http.Error(w, "Bottom item not found", http.StatusNotFound)
			return
		}
		existingOutfit.Bottoms = bottom
		existingOutfit.BottomID = bottom.ID
	}
	if updatedOutfit.OnePieceID != 0 {
		var onePiece Item
		if err := db.First(&onePiece, updatedOutfit.OnePieceID).Error; err != nil {
			http.Error(w, "One piece item not found", http.StatusNotFound)
			return
		}
		existingOutfit.OnePieces = onePiece
		existingOutfit.OnePieceID = onePiece.ID
	}
	if updatedOutfit.AccessoriesID != 0 {
		var accessory Item
		if err := db.First(&accessory, updatedOutfit.AccessoriesID).Error; err != nil {
			http.Error(w, "Accessory item not found", http.StatusNotFound)
			return
		}
		existingOutfit.Accessories = accessory
		existingOutfit.AccessoriesID = accessory.ID
	}
	if updatedOutfit.ShoesID != 0 {
		var shoes Item
		if err := db.First(&shoes, updatedOutfit.ShoesID).Error; err != nil {
			http.Error(w, "Shoes item not found", http.StatusNotFound)
			return
		}
		existingOutfit.Shoes = shoes
		existingOutfit.ShoesID = shoes.ID
	}

	fmt.Println("Outfit has been updated successfully")
	db.Save(&existingOutfit)
	json.NewEncoder(w).Encode(existingOutfit)
}

func DeleteOutfit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var outfit Outfit
	db.First(&outfit, params["id"])
	db.Delete(&outfit, params["id"])
	json.NewEncoder(w).Encode("The outfit has successfully been deleted.")
}

func GetOutfits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var outfits []Outfit
	db.Preload("Tops").Preload("Bottoms").Preload("OnePieces").Preload("Accessories").Preload("Shoes").Find(&outfits)
	json.NewEncoder(w).Encode(outfits)
}

func GetOutfit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var outfit Outfit
	db.Preload("Tops").Preload("Bottoms").Preload("OnePieces").Preload("Accessories").Preload("Shoes").First(&outfit, params["id"])
	json.NewEncoder(w).Encode(outfit)
}

func GetUserItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var items []Item
	db.Where("user_id = ?", params["id"]).Find(&items)
	json.NewEncoder(w).Encode(items)
}

func GetAllItemsCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var items []Item
	db.Where("user_id = ? AND category = ?", params["id"], params["name"]).Find(&items)
	json.NewEncoder(w).Encode(items)
}

func GetUserOutfits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var outfits []Outfit
	db.Preload("Tops").Preload("Bottoms").Preload("OnePieces").Preload("Accessories").Preload("Shoes").Where("user_id = ?", params["id"]).Find(&outfits)
	// db.Where("user_id = ?", params["id"]).Find(&outfits)
	json.NewEncoder(w).Encode(outfits)
}
