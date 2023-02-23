package main

import (
	//"errors"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	gorm.Model
	first_name string `json: "firstname"`
	last_name  string `json: "lastname"`
	email      string `json: "email"`
}

func InitialMigration() {
	// Connect to database
	db, err := gorm.Open(sqlite.Open("OOTD.db"), &gorm.Config{})

	// if error display message
	if err != nil {
		panic("failute to connect to database")
	}

	//print connected to display connection
	fmt.Println("connected")

	//create nested tables
	db.AutoMigrate(&User{})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User 
	json.NewDecoder(r.Body).Decode(&user)
	db.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func 

/*func searchForPerson(db *gorm.DB, searchString string) []user {

	rows, err := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE first_name like '%" + searchString + "%' OR last_name like '%" + searchString + "%'")

	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	people := make([]user, 0)

	for rows.Next() {
		ourPerson := user{}
		err = rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
		if err != nil {
			log.Fatal(err)
		}

		people = append(people, ourPerson)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return people
}

func getPersonById(db *sql.DB, ourID string) user {

	rows, _ := db.Query("SELECT id, first_name, last_name, email, ip_address FROM people WHERE id = '" + ourID + "'")
	defer rows.Close()

	ourPerson := user{}

	for rows.Next() {
		rows.Scan(&ourPerson.id, &ourPerson.first_name, &ourPerson.last_name, &ourPerson.email, &ourPerson.ip_address)
	}

	return ourPerson
}

func updatePerson(db *sql.DB, ourPerson user) int64 {

	stmt, err := db.Prepare("UPDATE people set first_name = ?, last_name = ?, email = ?, ip_address = ? where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(ourPerson.first_name, ourPerson.last_name, ourPerson.email, ourPerson.ip_address, ourPerson.id)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	return affected
}

func deletePerson(db *sql.DB, idToDelete string) int64 {

	stmt, err := db.Prepare("DELETE FROM people where id = ?")
	checkErr(err)
	defer stmt.Close()

	res, err := stmt.Exec(idToDelete)
	checkErr(err)

	affected, err := res.RowsAffected()
	checkErr(err)

	return affected
}
*/
