package main

import (
	//"errors"

	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	_ "github.com/mattn/go-sqlite3"
)

/*
var DB *gorm.DB

// An Item has many Tags (many to many relationship)

	type item struct {
		gorm.Model
		UserID    uint   `gorm:"primaryKey"`
		Name      string `json:"Name"`
		Category  string `json:"Category"`
		ImagePath string `json:"Image"`
	}

	type tag struct {
		TagID    uint `gorm:"primaryKey"`
		TagName  string
		Category string
	}

var migrationMutex sync.Mutex

	func ItemMigration() {
		migrationMutex.Lock()
		defer migrationMutex.Unlock()

		// Connect to database
		var err error
		DB, err = gorm.Open(sqlite.Open("Item.db"), &gorm.Config{})

		// if error display message
		if err != nil {
			panic("failure to connect to database")
		}

		//print connected to display connection
		fmt.Println("item1 db connected")

		//create nested tables within a transaction
		tx := DB.Begin()
		defer func() {
			if r := recover(); r != nil {
				tx.Rollback()
				panic(r)
			}
		}()
		if err := tx.AutoMigrate(&item{}, &tag{}); err != nil {
			tx.Rollback()
			panic(err)
		}
		tx.Commit()
	}

	func GetItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var Item []item
		DB.Find(&Item)
		json.NewEncoder(w).Encode(Item)

}

	func GetItems(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var Item item
		DB.First(&Item, params["id"])
		json.NewEncoder(w).Encode(Item)

}

	func CreateItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var Item item
		json.NewDecoder(r.Body).Decode(&Item)
		DB.Create(&Item)
		json.NewEncoder(w).Encode(Item)
	}

	func UpdateItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var Item item
		DB.First(&Item, params["id"])
		json.NewDecoder(r.Body).Decode(&Item)
		DB.Save(&Item)
		json.NewEncoder(w).Encode(Item)
	}

	func DeleteItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		var Item item
		DB.First(&Item, params["id"])
		DB.Delete(&Item, params["id"])
		json.NewEncoder(w).Encode("The item has successfully been deleted.")
	}
*/
type itemTag struct {
	ItemID uint `gorm:"primaryKey"`
	TagID  uint `gorm:"primaryKey"`
}

type item struct {
	gorm.Model
	UserID    uint   `gorm:"primaryKey"`
	Name      string `json:"Name"`
	Category  string `json:"Category"`
	ImagePath string `json:"Image"`
	//Tags      []tag  `gorm:"many2many:item_tags;"`
}

type tag struct {
	gorm.Model
	TagName  string
	Category string
	//Items    []item `gorm:"many2many:item_tags;"`
}

var dbMutex sync.Mutex

func Migrate() {

	db, err := gorm.Open(sqlite.Open("Item.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("open failed")
	}

	//err = db.AutoMigrate(&item{}, &tag{}, &itemTag{})
	err = db.AutoMigrate(&item{})
	if err != nil {
		fmt.Println("migrate failed")
	}

	err = db.AutoMigrate(&tag{})
	if err != nil {
		fmt.Println(" tag migrate failed")
	}

	fmt.Println("connected to item db")
}
