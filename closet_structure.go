// Notice that the User doesn't have a Closet anymore
// That's a Closet struct is really needed. A User just has a collection of Items.



import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"/Users/nataliehodnett/Desktop/Git Repository SWE/OOTD/backend/go/src/git/user/database.go"
)

type Tag struct {
	TagID    uint `gorm:"primaryKey"`
	TagName  string
	Category string
}

// An Item has many Tags (many to many relationship)
type Item struct {
	gorm.Model
	userID    uint   `gorm:"primaryKey"`
	Name      string `json:"name"`
	Category  string `json:"category"`
	ImagePath string `json:"image"`
}

// Notice now there's an extra table called "ItemTag"
// This creates a relationship between an Item and a Tag. Without this,
// we wouldn't have a relational database and thus we couldn't use SQL queries.
type ItemTag struct {
	ItemID uint
	TagID  uint
}

type user struct {
	userTag     Tag
	userItem    Item
	userItemTag ItemTag
}

/*
// ==TO CREATE A USER==//
func createUserStruct() {

		//Create array of tags
		Tag1 := Tag{123456, "boots", "shoes"}
		Tag2 := Tag{123456, "bra", "underwear"}
		Tag3 := Tag{123456, "headband", "accessories"}
		var array []Tag
		array[0] = Tag1
		array[1] = Tag2
		array[2] = Tag3

		//Create objects for tag, item, and item tag
		sampleTag := Tag{123456, "sample Name", "sample Category"}
		sampleItem := Item{12345, 1234, "category", array}
		sampleItemTag := ItemTag{123, 1234}

		//create sampleuser
		sampleUser := User{sampleTag, sampleItem, sampleItemTag}

		reqBody, _ := json.Marshal(sampleUser)

		// Create request to create user
		req, err := http.NewRequest("POST", "/users", bytes.NewBuffer(reqBody))

		// Set up router and execute request
		router := mux.NewRouter()
		router.HandleFunc("/users", CreateUser)
		rr := httptest.NewRecorder()
		router.ServeHTTP(rr, req)
	}
*/
func CreateItemTable() {

	// Migrate the schema
	err = db.AutoMigrate(&Item{})
	if err != nil {
		fmt.Println("Item Table failed to be created")
	}
	if err == nil {
		fmt.Println("Item Table created successfully!")
	}
}

/*
	func CreateItem(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var item Item
		json.NewDecoder(r.Body).Decode(&item)
		db.Create(&item)
		json.NewEncoder(w).Encode(item)
	}
*/
func createItemTagTable() error {
	db, err := gorm.Open(sqlite.Open("OOTD.db"), &gorm.Config{})
	if err != nil {
		return err
	}

	// Migrate the schema
	err = db.AutoMigrate(&ItemTag{})
	if err != nil {
		return err
	}

	fmt.Println("Item Table created successfully!")
	return nil
}

/*
func TagExists(tagID uint, tagName string, category string) (bool, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return false, err
	}

	// Check if the tag exists in the Tag table
	var count int64
	result := db.Model(&Tag{}).Where("tag_id = ?", tagID).Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	//The backend would have to check if the tags currently exist in our
	//Tag table. If the tag doesn’t exist, then add a new row with that tag information.
	if count > 0 {
		return true, nil
	} else {
		// Create a new Tag row with the given tag ID, name, and category
		newTag := &Tag{
			TagID:    tagID,
			TagName:  tagName,
			Category: category,
		}
		result := db.Create(newTag)
		if result.Error != nil {
			return false, result.Error
		}
		return false, nil
	}
}
*/
