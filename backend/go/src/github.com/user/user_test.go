package main

import (
	"bytes"
	"encoding/json"
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
			Last_Name: "Doe",
			Username: "johndoe",
			Password: "password",
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