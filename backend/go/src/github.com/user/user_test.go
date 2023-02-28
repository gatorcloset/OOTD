/*
package main

import (

	"testing"

	"src/go/src/github.com/nataliehodnett/util"


	"github.com/stretchr/testify/require"

	)

		f	unc CreateRandomAccount(t *testing.T) User {
					arg := CreateUserParams{
						First_Name: util.RandomName(),
						Last_Name:  util.RandomName(),
						Username:   util.RandomName(),
					Password:   util.RandomName(),
					}

				account, err := testQueries.CreateUser(context.Background(), arg)
				require.NoError(t, err)
					require.NotEmpty(t, account)

				require.Equal(t, arg.First_Name, account.First_Name)
				require.Equal(t, arg.Password, account.Password)

				return account
			}

			func TestCreateUser(t *testing.T) {
			CreateRandomAccount(t)
		}

		func TestGetUser(t *testing.T) {
			// Create a random user account
			account1 := CreateRandomAccount(t)

			// Call the GetUser function with the ID of the created account
			account2, err := testQueries.GetUser(context.Background(), account1.ID)

			// Ensure there were no errors in retrieving the account
			require.NoError(t, err)

			// Ensure the retrieved account is not empty
			require.NotEmpty(t, account2)

			// Ensure the retrieved account's fields match those of the created account
			require.Equal(t, account1.ID, account2.ID)
			require.Equal(t, account1.First_Name, account2.First_Name)
			require.Equal(t, account1.Last_Name, account2.Last_Name)
			require.Equal(t, account1.Username, account2.Username)
			require.Equal(t, account1.Password, account2.Password)
		}

		func TestUpdateAccount(t *testing.T) {
			account1 := CreateRandomAccount(t)

			arg := UpdateAccountParams{
				First_Name: account1.First_Name,
				Last_Name:  account1.Last_Name,
				Username:   account1.Username,
				Password:   account1.Password,
			}

			account2, err := testQueries.UpdateUser(context.Background(), arg)
			require.NoError(t, err)
				require.NotEmpty(t, account2)

				require.Equal(t, account1.First_Name, account2.First_Name)
				require.Equal(t, account1.Last_Name, account2.Last_Name)
				require.Equal(t, account1.Username, account2.Username)
				require.Equal(t, account1.Password, account2.Password)
			}

			func TestDeleteAccount(t *testing.T) {
				account1 := CreateRandomAccount(t)
				err := testQueries.DeleteUser(context.Background(), account1.First_Name)
				require.NoError(t, err)

				account2, err := testQueries.GetUser(context.Background(), account1.First_Name)
				require.Error(t, err)
				require.Empty(t, account2)
			}
*/
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	// Create a new router and recorder
	router := mux.NewRouter()
	recorder := httptest.NewRecorder()

	// Add the GetUsers endpoint to the router and send a GET request
	router.HandleFunc("/users", GetUsers)
	request, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(recorder, request)

	// Check that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check that the response body contains an empty array of users
	expected := "[]\n"
	assert.Equal(t, expected, recorder.Body.String())
}

func TestCreateUser(t *testing.T) {
	// Create a new router and recorder
	router := mux.NewRouter()
	recorder := httptest.NewRecorder()

	// Define a new user to create
	user := User{
		First_Name: "Megan",
		Last_Name:  "Shah",
		Username:   "meganshah",
		Password:   "password",
	}

	// Convert the user to JSON and create a new request
	body, _ := json.Marshal(user)
	request, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	// Add the CreateUser endpoint to the router and send the request
	router.HandleFunc("/users", CreateUser)
	router.ServeHTTP(recorder, request)

	// Check that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check that the response body contains the created user
	var createdUser User
	err := json.Unmarshal(recorder.Body.Bytes(), &createdUser)
	assert.Nil(t, err)
	assert.Equal(t, user.First_Name, createdUser.First_Name)
	assert.Equal(t, user.Last_Name, createdUser.Last_Name)
	assert.Equal(t, user.Username, createdUser.Username)
	assert.Equal(t, user.Password, createdUser.Password)
}

func TestGetUser(t *testing.T) {
	// Create a new router and recorder
	router := mux.NewRouter()
	recorder := httptest.NewRecorder()

	// Create a new user and add it to the database
	user := User{
		First_Name: "Natalie",
		Last_Name:  "Hodnett",
		Username:   "nataliehodnett",
		Password:   "password",
	}
	db.Create(&user)

	// Add the GetUser endpoint to the router and send a GET request
	router.HandleFunc("/users/{id}", GetUser)
	request, _ := http.NewRequest("GET", "/users/1", nil)
	router.ServeHTTP(recorder, request)

	// Check that the response status code is 200 OK
	assert.Equal(t, http.StatusOK, recorder.Code)

	// Check that the response body contains the expected user
	var retrievedUser User
	err := json.Unmarshal(recorder.Body.Bytes(), &retrievedUser)
	assert.Nil(t, err)
	assert.Equal(t, user.ID, retrievedUser.ID)
	assert.Equal(t, user.First_Name, retrievedUser.First_Name)
	assert.Equal(t, user.Last_Name, retrievedUser.Last_Name)
	assert.Equal(t, user.Username, retrievedUser.Username)
	assert.Equal(t, user.Password, retrievedUser.Password)
}
