package main

import (
	"bytes"
	"encoding/json"
	"log"
	"mime/multipart"
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

func TestCreateItem(t *testing.T) {
	// Create a new request with a fake image file
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)
	fileWriter, err := writer.CreateFormFile("image", "test.jpg")
	if err != nil {
		t.Fatalf("Failed to create form file: %v", err)
	}
	_, err = fileWriter.Write([]byte("This is a fake image"))
	if err != nil {
		t.Fatalf("Failed to write to form file: %v", err)
	}
	writer.WriteField("name", "Test Item")
	writer.WriteField("category", "Test Category")
	writer.Close()
	request, err := http.NewRequest("POST", "/items", requestBody)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	// Record the HTTP response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateItem)
	handler.ServeHTTP(rr, request)

	// Check the HTTP status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	// Check the response body
	expected := Item{Name: "Test Item", Category: "Test Category"}
	err = json.Unmarshal(rr.Body.Bytes(), &expected)
	if err != nil {
		t.Fatalf("Failed to unmarshal response: %v", err)
	}
	if expected.Name != "Test Item" {
		t.Errorf("Handler returned unexpected item name: got %v, want %v", expected.Name, "Test Item")
	}
	if expected.Category != "Test Category" {
		t.Errorf("Handler returned unexpected item category: got %v, want %v", expected.Category, "Test Category")
	}
}

/*
func TestCreateTag(t *testing.T) {
	// Create a new HTTP request
	requestBody, _ := json.Marshal(Tag{
		TagName:  "Test Tag",
		Category: "Test Category",
	})
	req, err := http.NewRequest("POST", "/tags", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Create a new recorder to record the HTTP response
	rr := httptest.NewRecorder()

	// Call the CreateTag function with the HTTP request and recorder
	CreateTag(rr, req)

	// Check the HTTP status code
	if rr.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, rr.Code)
	}

	// Check the response body
	expectedResponseBody := Tag{
		TagName:  "Test Tag",
		Category: "Test Category",
	}
	var responseBody Tag
	err = json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to parse response body: %v", err)
	}
	if !reflect.DeepEqual(responseBody, expectedResponseBody) {
		t.Errorf("Expected response body %v, but got %v", expectedResponseBody, responseBody)
	}
}
*/

func TestCreateTag(t *testing.T) {
	// Set up a temporary test database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open database: %v", err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("failed to get DB connection: %v", err)
		}
		defer sqlDB.Close()
	}()

	// Migrate the schema
	err = db.AutoMigrate(&Tag{})
	if err != nil {
		t.Fatalf("Failed to migrate database schema: %v", err)
	}

	// Create a new test tag
	newTag := Tag{
		TagName:  "testtag",
		Category: "testcategory",
	}
	requestBody, err := json.Marshal(newTag)
	if err != nil {
		t.Fatalf("Failed to marshal JSON request body: %v", err)
	}

	// Set up the request
	req, err := http.NewRequest("POST", "/tags", bytes.NewBuffer(requestBody))
	if err != nil {
		t.Fatalf("Failed to create HTTP request: %v", err)
	}

	// Set up the response recorder
	rr := httptest.NewRecorder()

	// Call the handler function
	handler := http.HandlerFunc(CreateTag)
	handler.ServeHTTP(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected response status code: got %v, want %v", status, http.StatusOK)
	}

	// Check the response body
	expectedTag := Tag{
		TagName:  "testtag",
		Category: "testcategory",
	}
	expectedResponseBody, err := json.Marshal(expectedTag)
	if err != nil {
		t.Fatalf("Failed to marshal expected response body: %v", err)
	}
	if rr.Body.String() != string(expectedResponseBody) {
		t.Errorf("Unexpected response body: got %v, want %v", rr.Body.String(), string(expectedResponseBody))
	}
}
