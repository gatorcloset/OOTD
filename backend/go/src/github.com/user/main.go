package main

//import statements
import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func initializeRouter() {
	// Create a new CORS handler
	corsHandler := cors.Default().Handler

	// Create a Mux router
	r := mux.NewRouter()

	// API endpoints
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	// Wrap the router with the CORS middleware
	log.Fatal(http.ListenAndServe(":9000", corsHandler(r)))

}

func main() {

	InitialMigration()
	initializeRouter()

}
