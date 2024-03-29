package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var testDB *gorm.DB

func setupTest() {
	// Connect to the test database
	var err error
	testDB, err = gorm.Open(sqlite.Open("test2.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to test database")
	}

	// Run initial migration to create necessary tables
	InitialMigration()
}

func cleanupTest() {
	// Delete all records from the database after each test
	testDB.Delete(&User{})
}

//==TEST USER==//

func TestCreateUser(t *testing.T) {
	// Set up test database connection and data
	setupTest()
	defer cleanupTest()

	// Create a test user
	user := User{
		First_Name: "John",
		Last_Name:  "Doe",
		Username:   "jdoe",
		Password:   "password",
	}
	reqBody, _ := json.Marshal(user)

	// Create request to create user
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set up router and execute request
	router := mux.NewRouter()
	router.HandleFunc("/users", CreateUser)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	createdUser := User{
		First_Name: "John",
		Last_Name:  "Doe",
		Username:   "jdoe",
		Password:   "password",
	}

	// Check that user was created in the database
	assert.Equal(t, user.First_Name, createdUser.First_Name)
	assert.Equal(t, user.Last_Name, createdUser.Last_Name)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Password, createdUser.Password)

	// Check that the response body contains the created user
	//expectedResBody, _ := json.Marshal(createdUser)
	//assert.Equal(t, string(expectedResBody), rr.Body.String())
}

func TestDeleteUser(t *testing.T) {
	// Set up test database connection and data
	setupTest()
	defer cleanupTest()
	user := User{First_Name: "John", Last_Name: "Doe", Username: "johndoe", Password: "password"}
	db.Create(&user)

	// Create a new request to delete the user
	req, err := http.NewRequest("DELETE", "/users/"+strconv.Itoa(int(user.ID)), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the DeleteUser function with the request and response recorder
	handler := http.HandlerFunc(DeleteUser)
	handler.ServeHTTP(rr, req)

	// Assert that the response status code is OK
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Assert that the response body is correct
	assert.Equal(t, "\"The user has successfully been deleted.\"\n", rr.Body.String(), "handler returned unexpected body")
}

func TestUpdateUser(t *testing.T) {
	setupTest()
	defer cleanupTest()
	// Create a new user
	user := User{First_Name: "John", Last_Name: "Doe", Username: "johndoe", Password: "password"}
	db.Create(&user)

	// Define the update request body
	update := User{First_Name: "Jane", Last_Name: "Doe", Username: "janedoe", Password: "newpassword"}

	// Update the user in the database
	db.Model(&user).Updates(update)

	// Check that the user was updated in the database
	var updatedUser User
	db.First(&updatedUser, user.ID)
	assert.Equal(t, update.First_Name, updatedUser.First_Name)
	assert.Equal(t, update.Last_Name, updatedUser.Last_Name)
	assert.Equal(t, update.Username, updatedUser.Username)
	assert.Equal(t, update.Password, updatedUser.Password)

	// Delete the user from the database
	db.Delete(&updatedUser)
}

type expectedUser struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func TestGetUser(t *testing.T) {
	setupTest()
	defer cleanupTest()

	// Create a new user
	user := User{First_Name: "John", Last_Name: "Doe", Username: "johndoe", Password: "password"}
	db.Create(&user)

	// Create a new request with a URL that includes the user ID parameter
	req, err := http.NewRequest("GET", "/users/"+strconv.Itoa(int(user.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Call the GetUser function with the request and response recorder
	handler := http.HandlerFunc(GetUser)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response Content-Type header
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ctype, "application/json")
	}

	// Check the response body (JSON-encoded user object)
	var responseUser expectedUser
	if err := json.Unmarshal(rr.Body.Bytes(), &responseUser); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	expected := expectedUser{
		Firstname: "John",
		Lastname:  "Doe",
		Username:  "johndoe",
		Password:  "password",
	}
	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(expected)
	if string(actualJson) != string(expectedJson) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(actualJson), string(expectedJson))
	}

	// Delete the user from the database
	db.Delete(&user)
}

func TestGetUsers(t *testing.T) {
	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to record the response
	rr := httptest.NewRecorder()

	// Initialize a new router and define the route for getting all users
	router := mux.NewRouter()
	router.HandleFunc("/users", GetUsers).Methods("GET")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var users []User
	err = json.Unmarshal(rr.Body.Bytes(), &users)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, users)
}

//==TEST TAG==//

func TestCreateTag(t *testing.T) {
	//Set up test database connection and data
	setupTest()
	defer cleanupTest()

	// Create a new tag
	newTag := Tag{TagName: "Test Tag"}

	reqBody, _ := json.Marshal(newTag)

	//Create request to create tag
	req, err := http.NewRequest("POST", "/tag", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	//Set up router and execute request
	router := mux.NewRouter()
	router.HandleFunc("/tag", CreateTag)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	//Check response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	createdTag := Tag{
		TagName: "Test Tag",
	}

	assert.Equal(t, newTag.TagName, createdTag.TagName)
}

type expectedTag struct {
	TagName string `json:"tagname"`
}

func TestGetTag(t *testing.T) {
	setupTest()
	defer cleanupTest()

	// Create a new tag
	tag := Tag{TagName: "Test Tag"}
	db.Create(&tag)

	// Create a new request with a URL that includes the user ID parameter
	req, err := http.NewRequest("GET", "/tag/"+strconv.Itoa(int(tag.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Call the GetTag function with the request and response recorder
	handler := http.HandlerFunc(GetTag)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response Content-Type header
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ctype, "application/json")
	}

	// Check the response body (JSON-encoded user object)
	var responseTag expectedTag
	if err := json.Unmarshal(rr.Body.Bytes(), &responseTag); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	expected := expectedTag{
		TagName: "Test Tag",
	}
	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(expected)
	if string(actualJson) != string(expectedJson) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(actualJson), string(expectedJson))
	}

	// Delete the tag from the database
	db.Delete(&tag)
}

func TestGetTags(t *testing.T) {
	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", "/tag", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to record the response
	rr := httptest.NewRecorder()

	// Initialize a new router and define the route for getting all users
	router := mux.NewRouter()
	router.HandleFunc("/tag", GetUsers).Methods("GET")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var tag []Tag
	err = json.Unmarshal(rr.Body.Bytes(), &tag)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, tag)
}

func TestUpdateTags(t *testing.T) {
	setupTest()
	defer cleanupTest()
	// Create a new user
	tag := Tag{TagName: ""}
	db.Create(&tag)

	// Define the update request body
	update := Tag{TagName: ""}

	// Update the user in the database
	db.Model(&tag).Updates(update)

	// Check that the user was updated in the database
	var updatedTag Tag
	db.First(&updatedTag, tag.ID)
	assert.Equal(t, update.TagName, updatedTag.TagName)

	// Delete the user from the database
	db.Delete(&updatedTag)
}

func TestDeleteTag(t *testing.T) {
	// Set up test database connection and data
	setupTest()
	defer cleanupTest()
	tag := Tag{TagName: "Test Tag"}
	db.Create(&tag)

	// Create a new request to delete the user
	req, err := http.NewRequest("DELETE", "/tag/"+strconv.Itoa(int(tag.ID)), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the DeleteUser function with the request and response recorder
	handler := http.HandlerFunc(DeleteTag)
	handler.ServeHTTP(rr, req)

	// Assert that the response status code is OK
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Assert that the response body is correct
	assert.Equal(t, "\"This tag has successfully been deleted.\"\n", rr.Body.String(), "handler returned unexpected body")
}

// ==TEST ITEM==//
func TestCreateItem(t *testing.T) {
	//Set up test database connection and data
	setupTest()
	defer cleanupTest()

	// Create a new item
	item := Item{Name: "Name", Category: "Category"}

	reqBody, _ := json.Marshal(item)

	//Create request to create item
	req, err := http.NewRequest("POST", "/item", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	//Set up router and execute request
	router := mux.NewRouter()
	router.HandleFunc("/item", CreateTag)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	//Check response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	createdItem := Item{
		Name:     "Name",
		Category: "Category",
	}

	assert.Equal(t, item.Name, createdItem.Name)
	assert.Equal(t, item.Category, createdItem.Category)
}

type expectedItem struct {
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	ImagePath string `json:"image"`
}

func TestGetItem(t *testing.T) {
	setupTest()
	defer cleanupTest()

	// Create a new item
	item := Item{Name: "Name", Category: "Category"}
	db.Create(&item)

	// Create a new request with a URL that includes the user ID parameter
	req, err := http.NewRequest("GET", "/item/"+strconv.Itoa(int(item.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Call the GetItem function with the request and response recorder
	handler := http.HandlerFunc(GetItem)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response Content-Type header
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ctype, "application/json")
	}

	// Check the response body (JSON-encoded user object)
	var responseTag expectedItem
	if err := json.Unmarshal(rr.Body.Bytes(), &responseTag); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	expected := expectedItem{
		Name:     "Name",
		Category: "Category",
	}
	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(expected)
	if string(actualJson) != string(expectedJson) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(actualJson), string(expectedJson))
	}

	// Delete the item from the database
	db.Delete(&item)
}

func TestDeleteItem(t *testing.T) {
	// Set up test database connection and data
	setupTest()
	defer cleanupTest()
	item := Item{Name: "Name", Category: "Category"}

	db.Create(&item)

	// Create a new request to delete the user
	req, err := http.NewRequest("DELETE", "/users/"+strconv.Itoa(int(item.ID)), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the DeleteUser function with the request and response recorder
	handler := http.HandlerFunc(DeleteItem)
	handler.ServeHTTP(rr, req)

	// Assert that the response status code is OK
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Assert that the response body is correct
	assert.Equal(t, "\"This item has successfully been deleted.\"\n", rr.Body.String(), "handler returned unexpected body")
}

func TestUpdateItem(t *testing.T) {
	setupTest()
	defer cleanupTest()
	// Create a new user
	item := Item{Name: "Name", Category: "Category"}
	db.Create(&item)

	// Define the update request body
	update := Item{Name: "Name", Category: "Category"}

	// Update the user in the database
	db.Model(&item).Updates(update)

	// Check that the user was updated in the database
	var updatedItem Item
	db.First(&updatedItem, item.ID)
	assert.Equal(t, update.Name, updatedItem.Name)
	assert.Equal(t, update.Category, updatedItem.Category)

	// Delete the user from the database
	db.Delete(&updatedItem)
}

func TestGetItems(t *testing.T) {
	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", "/item", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new HTTP recorder to record the response
	rr := httptest.NewRecorder()

	// Initialize a new router and define the route for getting all users
	router := mux.NewRouter()
	router.HandleFunc("/item", GetUsers).Methods("GET")

	// Serve the HTTP request
	router.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var item []Item
	err = json.Unmarshal(rr.Body.Bytes(), &item)
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(t, item)
}

func TestLoginUser(t *testing.T) {
	// Set up a test user
	user := User{
		First_Name: "John",
		Last_Name:  "Doe",
		Username:   "johndoe@example.com",
	}
	user.Password, _ = HashPassword("password")
	db.Create(&user)
	defer db.Delete(&user)

	// Set up the request body
	requestBody := map[string]string{
		"username": "johndoe@example.com",
		"password": "password",
	}
	requestBodyBytes, _ := json.Marshal(requestBody)

	// Set up the request
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(requestBodyBytes))

	// Set up the response recorder
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(LoginUser)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if rr.Code != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v",
			rr.Code, http.StatusOK)
	}

	// Check the response body
	var responseBody User
	if err := json.NewDecoder(rr.Body).Decode(&responseBody); err != nil {
		t.Errorf("Error unmarshaling response body: %v", err)
	}
	if responseBody.Username != "johndoe@example.com" {
		t.Errorf("Handler returned wrong response body: got %v, want %v",
			responseBody.Username, "johndoe@example.com")
	}
}

func TestPasswordHashing(t *testing.T) {
	// Create a test user
	setupTest()
	defer cleanupTest()

	user := User{
		First_Name: "John",
		Last_Name:  "Doe",
		Username:   "johndoe",
		Password:   "password",
	}

	// Encode the user as JSON
	jsonUser, err := json.Marshal(user)
	if err != nil {
		t.Fatal(err)
	}

	// Make a POST request to create the user
	req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(rr, req)

	// Check that the status code is 200
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Decode the response body
	var createdUser User
	if err := json.NewDecoder(rr.Body).Decode(&createdUser); err != nil {
		t.Errorf("error unmarshaling response body: %v", err)
	}

	// Check that the password in the database is not what the user entered
	if createdUser.Password == user.Password {
		t.Errorf("passwords match: %v", createdUser.Password)
	}

	// Delete the test user
	db.Delete(&createdUser)
}

// ==TEST CREATE OUTFIT==
func TestCreateOutfit(t *testing.T) {
	// Set up test database connection and data
	setupTest()
	defer cleanupTest()

	// Create a new outfit
	outfit := Outfit{
		Name:        "Test Outfit",
		Tops:        Item{Name: "Tops", Category: "Category"},
		Bottoms:     Item{Name: "Bottoms", Category: "Category"},
		OnePieces:   Item{Name: "OnePieces", Category: "Category"},
		Accessories: Item{Name: "Accessories", Category: "Category"},
		Shoes:       Item{Name: "Shoes", Category: "Category"},
	}

	reqBody, _ := json.Marshal(outfit)

	// Create request to create outfit
	req, err := http.NewRequest("POST", "/outfit", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatal(err)
	}

	// Set up router and execute request
	router := mux.NewRouter()
	router.HandleFunc("/outfit", CreateOutfit)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	// Check response status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check that the created outfit has the correct name
	var createdOutfit Outfit
	json.NewDecoder(rr.Body).Decode(&createdOutfit)
	assert.Equal(t, outfit.Name, createdOutfit.Name)
}

func TestDeleteOutfit(t *testing.T) {
	// Set up test database connection and data
	setupTest()
	defer cleanupTest()
	outfit := Outfit{
		Tops: Item{
			Name:      "Test Top",
			Category:  "tops",
			ImagePath: "test_top.jpg",
		},
		Bottoms: Item{
			Name:      "Test Bottom",
			Category:  "bottoms",
			ImagePath: "test_bottom.jpg",
		},
		OnePieces: Item{
			Name:      "Test One Piece",
			Category:  "one-pieces",
			ImagePath: "test_one_piece.jpg",
		},
		Accessories: Item{
			Name:      "Test Accessory",
			Category:  "accessories",
			ImagePath: "test_accessory.jpg",
		},
		Shoes: Item{
			Name:      "Test Shoes",
			Category:  "shoes",
			ImagePath: "test_shoes.jpg",
		},
	}

	db.Create(&outfit)

	// Create a new request to delete the user
	req, err := http.NewRequest("DELETE", "/outfit/"+strconv.Itoa(int(outfit.ID)), nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	// Create a new response recorder
	rr := httptest.NewRecorder()

	// Call the DeleteUser function with the request and response recorder
	handler := http.HandlerFunc(DeleteItem)
	handler.ServeHTTP(rr, req)

	// Assert that the response status code is OK
	assert.Equal(t, http.StatusOK, rr.Code, "handler returned wrong status code")

	// Assert that the response body is correct
	assert.Equal(t, "\"This item has successfully been deleted.\"\n", rr.Body.String(), "handler returned unexpected body")
}

type expectedOutfit struct {
	gorm.Model
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

func TestGetOutfit(t *testing.T) {
	setupTest()
	defer cleanupTest()

	// Create a new item
	outfit := Outfit{
		Tops: Item{
			Name:      "Test Top",
			Category:  "tops",
			ImagePath: "test_top.jpg",
		},
		Bottoms: Item{
			Name:      "Test Bottom",
			Category:  "bottoms",
			ImagePath: "test_bottom.jpg",
		},
		OnePieces: Item{
			Name:      "Test One Piece",
			Category:  "one-pieces",
			ImagePath: "test_one_piece.jpg",
		},
		Accessories: Item{
			Name:      "Test Accessory",
			Category:  "accessories",
			ImagePath: "test_accessory.jpg",
		},
		Shoes: Item{
			Name:      "Test Shoes",
			Category:  "shoes",
			ImagePath: "test_shoes.jpg",
		},
	}
	db.Create(&outfit)

	// Create a new request with a URL that includes the user ID parameter
	req, err := http.NewRequest("GET", "/outfit/"+strconv.Itoa(int(outfit.ID)), nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Call the GetItem function with the request and response recorder
	handler := http.HandlerFunc(GetItem)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response Content-Type header
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v",
			ctype, "application/json")
	}

	// Check the response body (JSON-encoded user object)
	var responseTag expectedItem
	if err := json.Unmarshal(rr.Body.Bytes(), &responseTag); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}

	expected := expectedOutfit{
		Tops: Item{
			Name:      "Test Top",
			Category:  "tops",
			ImagePath: "test_top.jpg",
		},
		Bottoms: Item{
			Name:      "Test Bottom",
			Category:  "bottoms",
			ImagePath: "test_bottom.jpg",
		},
		OnePieces: Item{
			Name:      "Test One Piece",
			Category:  "one-pieces",
			ImagePath: "test_one_piece.jpg",
		},
		Accessories: Item{
			Name:      "Test Accessory",
			Category:  "accessories",
			ImagePath: "test_accessory.jpg",
		},
		Shoes: Item{
			Name:      "Test Shoes",
			Category:  "shoes",
			ImagePath: "test_shoes.jpg",
		},
	}
	expectedJson, _ := json.Marshal(expected)
	actualJson, _ := json.Marshal(expected)
	if string(actualJson) != string(expectedJson) {
		t.Errorf("handler returned unexpected body: got %v want %v",
			string(actualJson), string(expectedJson))
	}

	// Delete the item from the database
	db.Delete(&outfit)
}

func TestGetOutfits(t *testing.T) {
	setupTest()
	defer cleanupTest()

	// Create some test outfits
	outfits := []Outfit{
		{
			Tops: Item{
				Name:      "Test Top 1",
				Category:  "tops",
				ImagePath: "test_top_1.jpg",
			},
			Bottoms: Item{
				Name:      "Test Bottom 1",
				Category:  "bottoms",
				ImagePath: "test_bottom_1.jpg",
			},
			OnePieces: Item{
				Name:      "Test One Piece 1",
				Category:  "one-pieces",
				ImagePath: "test_one_piece_1.jpg",
			},
			Accessories: Item{
				Name:      "Test Accessory 1",
				Category:  "accessories",
				ImagePath: "test_accessory_1.jpg",
			},
			Shoes: Item{
				Name:      "Test Shoes 1",
				Category:  "shoes",
				ImagePath: "test_shoes_1.jpg",
			},
		},
		{
			Tops: Item{
				Name:      "Test Top 2",
				Category:  "tops",
				ImagePath: "test_top_2.jpg",
			},
			Bottoms: Item{
				Name:      "Test Bottom 2",
				Category:  "bottoms",
				ImagePath: "test_bottom_2.jpg",
			},
			OnePieces: Item{
				Name:      "Test One Piece 2",
				Category:  "one-pieces",
				ImagePath: "test_one_piece_2.jpg",
			},
			Accessories: Item{
				Name:      "Test Accessory 2",
				Category:  "accessories",
				ImagePath: "test_accessory_2.jpg",
			},
			Shoes: Item{
				Name:      "Test Shoes 2",
				Category:  "shoes",
				ImagePath: "test_shoes_2.jpg",
			},
		},
	}

	// Add the test outfits to the database
	for i := range outfits {
		db.Create(&outfits[i])
	}

	// Create a new request
	req, err := http.NewRequest("GET", "/outfit", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a new ResponseRecorder (which satisfies http.ResponseWriter) to record the response
	rr := httptest.NewRecorder()

	// Call the GetOutfits function with the request and response recorder
	handler := http.HandlerFunc(GetOutfits)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response Content-Type header
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("handler returned wrong content type: got %v want %v", ctype, "application/json")
	}

	// Check the response body (JSON-encoded list of outfit objects)
	var responseOutfits []Outfit
	if err := json.Unmarshal(rr.Body.Bytes(), &responseOutfits); err != nil {
		t.Errorf("failed to unmarshal response body: %v", err)
	}
}

func TestUpdateOutfit(t *testing.T) {
	setupTest()
	defer cleanupTest()

	// Create test items for outfit
	top := Item{Name: "Test Top", Category: "tops", ImagePath: "test_one_piece_1.jpg"}
	bottom := Item{Name: "Test Bottom", Category: "bottoms", ImagePath: "test_one_piece_1.jpg"}
	onePiece := Item{Name: "Test One Piece", Category: "one-pieces", ImagePath: "test_one_piece_1.jpg"}
	accessory := Item{Name: "Test Accessory", Category: "accessories", ImagePath: "test_one_piece_1.jpg"}
	shoes := Item{Name: "Test Shoes", Category: "shoes", ImagePath: "test_one_piece_1.jpg"}

	// Create test outfit
	newOutfit := Outfit{
		Name:          "",
		Tops:          top,
		TopID:         top.ID,
		Bottoms:       bottom,
		BottomID:      bottom.ID,
		OnePieces:     onePiece,
		OnePieceID:    onePiece.ID,
		Accessories:   accessory,
		AccessoriesID: accessory.ID,
		Shoes:         shoes,
		ShoesID:       shoes.ID,
	}

	// Add outfit to database
	db.Create(&newOutfit)

	// Update outfit
	newName := ""
	newBottom := Item{Name: "", Category: "bottoms", ImagePath: "test_one_piece_1.jpg"}
	newOnePiece := Item{Name: "", Category: "one-pieces", ImagePath: "test_one_piece_1.jpg"}

	updatedOutfit := Outfit{
		Name:          newName,
		Bottoms:       newBottom,
		BottomID:      newBottom.ID,
		OnePieces:     newOnePiece,
		OnePieceID:    newOnePiece.ID,
		AccessoriesID: 0,
		ShoesID:       0,
	}
	updatedOutfit.ID = newOutfit.ID

	// Update the outfit object before creating the JSON payload
	newOutfit.Name = newName
	newOutfit.Bottoms = newBottom
	newOutfit.BottomID = newBottom.ID
	newOutfit.OnePieces = newOnePiece
	newOutfit.OnePieceID = newOnePiece.ID
	newOutfit.AccessoriesID = 0
	newOutfit.ShoesID = 0

	outfitJSON, _ := json.Marshal(newOutfit)
	req, err := http.NewRequest("PUT", "/outfit/"+strconv.Itoa(int(newOutfit.ID)), bytes.NewBuffer(outfitJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(UpdateOutfit)
	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check that the outfit has been updated in the database
	var updatedOutfitDB Outfit
	db.First(&updatedOutfitDB, newOutfit.ID)

	assert.Equal(t, updatedOutfit.Name, updatedOutfitDB.Name)
	assert.Equal(t, newBottom.Name, updatedOutfitDB.Bottoms.Name)
	assert.Equal(t, newOnePiece.Name, updatedOutfitDB.OnePieces.Name)
}

func TestGetAllItemsCategory(t *testing.T) {
	setupTest()
	defer cleanupTest()
	// Set up test server and client
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}/category/{name}", GetAllItemsCategory).Methods("GET")
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Make request to test server
	url := fmt.Sprintf("%s/users/%d/category/%s", ts.URL, 1, "tops")
	res, err := http.Get(url)
	if err != nil {
		t.Fatalf("Error making GET request: %v", err)
	}
	defer res.Body.Close()
	fmt.Println(res.StatusCode)

	// Check response code
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Decode response body
	var items []Item
	err = json.NewDecoder(res.Body).Decode(&items)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Check that returned items belong to the correct category and user
	for _, item := range items {
		assert.Equal(t, "tops", item.Category)
		assert.Equal(t, uint(1), item.UserID)
	}
}

func TestGetUserItems(t *testing.T) {
	setupTest()
	defer cleanupTest()

	// Set up test server and client
	router := mux.NewRouter()
	router.HandleFunc("/users/{id}/items", GetUserItems).Methods("GET")
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Add some items for the user to the test database
	user := User{
		First_Name: "test",
		Last_Name:  "user",
		Username:   "testuser",
		Password:   "password",
	}
	db.Create(&user)
	items := []Item{
		{Name: "item1", Category: "books", UserID: user.ID},
		{Name: "item2", Category: "clothes", UserID: user.ID},
		{Name: "item3", Category: "books", UserID: user.ID},
	}
	db.Create(&items)

	// Make request to test server
	url := fmt.Sprintf("%s/users/%d/items", ts.URL, user.ID)
	res, err := http.Get(url)
	if err != nil {
		t.Fatalf("Error making GET request: %v", err)
	}
	defer res.Body.Close()

	// Check response code
	assert.Equal(t, http.StatusOK, res.StatusCode)

	// Decode response body
	var itemsReturned []Item
	err = json.NewDecoder(res.Body).Decode(&itemsReturned)
	if err != nil {
		t.Fatalf("Error decoding response body: %v", err)
	}

	// Check that returned items belong to the correct user
	for _, item := range itemsReturned {
		assert.Equal(t, user.ID, item.UserID)
	}
}
