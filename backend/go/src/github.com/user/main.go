package main

//import statements
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/login", LoginUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":9000", r))

}

func main() {
	InitialMigration()
	initializeRouter()
}
