/*

package main

import (
	"database/sql"
	"log"
	"os"
	"testing"
)

const (
	dbDriver = "sqlite"
	dbSource = "OOTD.db"
)

func TestMain(m *testing.M) {
    // Connect to your database here and initialize your query object
	db, err := sql.Open("sqlite", "OOTD.db")
    if err != nil {
        log.Fatal(err)
  }
	 defer db.Close()

    testQueries = NewQueryObject(db)

    os.Exit(m.Run())
}
*/

package main

import (
    "log"
    "os"
    "testing"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestinitDB(m *testing.M) {
    dataSourceName := "OOTD.db"
    db, err := gorm.Open("sqlite", dataSourceName)

    if err != nil {
        log.Fatal("failed to connect database")
    }
    //db.Exec("CREATE DATABASE test")
    db.LogMode(true)
    db.Exec("USE test111")
    os.Exit(m.Run())
}

