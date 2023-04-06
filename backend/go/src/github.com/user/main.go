package main

//import statements
import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
)

func initializeRouter() {
    corsHandler := cors.Default().Handler

    r := mux.NewRouter()
    r.HandleFunc("/users", GetUsers).Methods("GET")
    r.HandleFunc("/users/{id}", GetUser).Methods("GET")
    r.HandleFunc("/users", CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
    r.HandleFunc("/login", LoginUser).Methods("POST")
    r.HandleFunc("/item", CreateItem).Methods("POST")
    r.HandleFunc("/tag", CreateTag).Methods("POST")
    r.HandleFunc("/item_tag", CreateItemTag).Methods("POST")
    r.HandleFunc("/item/{id}", GetItem).Methods("GET")
    r.HandleFunc("/tag/{id}", GetTag).Methods("GET")
	r.HandleFunc("/tag", GetTags).Methods("GET")
    r.HandleFunc("/tag/{id}", DeleteTag).Methods("DELETE")
	r.HandleFunc("/tag/{id}", UpdateTag).Methods("PUT")

    log.Fatal(http.ListenAndServe(":9000", corsHandler(r)))
}

func main() {
    InitialMigration()
    initializeRouter()
}