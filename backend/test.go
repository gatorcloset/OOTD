package main

//import statements
import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/dixonwille/wmenu/v5"
)

func initializeRouter () {
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users{id}", GetUser).Methods("GET")
	r.HandleFunc("/users", CreateUser).Methods("POST")
	r.HandleFunc("/users{id}", UpdateUsers).Methods("PUT")
	r.HandleFunc("/users{id}", DeleteUsers).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))

}

func main() {

	InitialMigration()
	initializeRouter()


	//displays menu and asks for user input
	menu := wmenu.NewMenu("What would you like to do?")

	menu.Action(func(opts []wmenu.Opt) error { handleFunc(db, opts); return nil })

	menu.Option("Add a new User", 0, true, nil)
	menu.Option("Find an Existing User", 1, false, nil)
	menu.Option("Update a User's Sign In Information", 2, false, nil)
	menu.Option("Delete a user", 3, false, nil)
	menu.Option("Quit Application", 4, false, nil)
	menuerr := menu.Run()

	if menuerr != nil {
		log.Fatal(menuerr)
	}
}

func handleFunc(db *gorm.DB, opts []wmenu.Opt) {

	//if user enters 0, get first name, last name, and email and add to the database
	switch opts[0].Value {

	case 0:
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Enter a first name: ")
		firstName, _ := reader.ReadString('\n')
		if firstName != "\n" {
			firstName = strings.TrimSuffix(firstName, "\n")
		}

		fmt.Print("Enter a last name: ")
		lastName, _ := reader.ReadString('\n')
		if lastName != "\n" {
			lastName = strings.TrimSuffix(lastName, "\n")
		}

		fmt.Print("Enter an email address: ")
		email, _ := reader.ReadString('\n')
		if email != "\n" {
			email = strings.TrimSuffix(email, "\n")
		}


		user := User{
			first_name: firstName,
 			last_name: lastName,
 			email: email,
		}

		createUser(db, user)

		fmt.Println("Account created", user.first_name, user.last_name)

		db.Save(&user)

		break
	}
}

	/*case 1:
		fmt.Println("Finding a Person")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a name to search for : ")
		searchString, _ := reader.ReadString('\n')
		searchString = strings.TrimSuffix(searchString, "\n")
		people := searchForPerson(db, searchString)

		fmt.Printf("Found %v results\n", len(people))

		for _, ourPerson := range people {
			fmt.Printf("\n----\nFirst Name: %s\nLast Name: %s\nEmail: %s\nIP Address: %s\n", ourPerson.first_name, ourPerson.last_name, ourPerson.email, ourPerson.ip_address)
		}
		break

	case 2:
		fmt.Println("Update a User's information")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter an id to update: ")
		updateid, _ := reader.ReadString('\n')

		currentPerson := getPersonById(db, updateid)

		fmt.Printf("First Name (Currently %s):", currentPerson.first_name)
		firstName, _ := reader.ReadString('\n')
		if firstName != "\n" {
			currentPerson.first_name = strings.TrimSuffix(firstName, "\n")
		}

		fmt.Printf("Last Name (Currently %s):", currentPerson.last_name)
		lastName, _ := reader.ReadString('\n')
		if lastName != "\n" {
			currentPerson.last_name = strings.TrimSuffix(lastName, "\n")
		}

		fmt.Printf("Email (Currently %s):", currentPerson.email)
		email, _ := reader.ReadString('\n')
		if email != "\n" {
			currentPerson.email = strings.TrimSuffix(email, "\n")
		}

		fmt.Printf("IP Address (Currently %s):", currentPerson.ip_address)
		ipAddress, _ := reader.ReadString('\n')
		if ipAddress != "\n" {
			currentPerson.ip_address = strings.TrimSuffix(ipAddress, "\n")
		}

		affected := updatePerson(db, currentPerson)

		if affected == 1 {
			fmt.Println("One row affected")
		}

		break

	case 3:
		fmt.Println("Deleting a person by ID")

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter the ID you want to delete : ")
		searchString, _ := reader.ReadString('\n')

		idToDelete := strings.TrimSuffix(searchString, "\n")

		affected := deletePerson(db, idToDelete)

		if affected == 1 {
			fmt.Println("Deleted person from database")
		}

		break

	case 4:
		fmt.Println("Quitting application")
		fmt.Println("Goodbye!")
		os.Exit(3)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
	
}
*/

