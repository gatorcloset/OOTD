package main

//import statements
import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func initializeRouter() {
	corsOptions := cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		AllowCredentials: true,
	}
	corsHandler := cors.New(corsOptions).Handler

	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
	r.HandleFunc("/login", LoginUser).Methods("POST")
	r.HandleFunc("/logout", Logout).Methods("POST")
	r.HandleFunc("/item", CreateItem).Methods("POST")
	r.HandleFunc("/item/{id}", UpdateItem).Methods("PUT")
	r.HandleFunc("/item/{id}", DeleteItem).Methods("DELETE")
	r.HandleFunc("/tag", CreateTag).Methods("POST")
	r.HandleFunc("/item_tag", CreateItemTag).Methods("POST")
	r.HandleFunc("/item/{id}", GetItem).Methods("GET")
	r.HandleFunc("/tag/{id}", GetTag).Methods("GET")
	r.HandleFunc("/tag", GetTags).Methods("GET")
	r.HandleFunc("/tag/{id}", DeleteTag).Methods("DELETE")
	r.HandleFunc("/tag/{id}", UpdateTag).Methods("PUT")
	r.HandleFunc("/authUser", getUserData).Methods("GET")
	r.HandleFunc("/authUser", getUserData).Methods("GET")
	r.HandleFunc("/item", GetItems).Methods("GET")
	r.HandleFunc("/outfit", CreateOutfit).Methods("POST")
	r.HandleFunc("/outfit/{id}", UpdateOutfit).Methods("PUT")
	r.HandleFunc("/outfit/{id}", DeleteOutfit).Methods("DELETE")
	r.HandleFunc("/outfit/{id}", GetOutfit).Methods("GET")
	r.HandleFunc("/outfit", GetOutfits).Methods("GET")
	r.HandleFunc("/users/{id}/items", GetUserItems).Methods("GET")
	r.HandleFunc("/users/{id}/category/{name}", GetAllItemsCategory).Methods("GET")

	log.Fatal(http.ListenAndServe(":9000", corsHandler(r)))
}

func main() {
	InitialMigration()
	initializeRouter()
}
