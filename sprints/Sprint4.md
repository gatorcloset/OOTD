# Completed Work

[Sprint 4 Video Demo]()

## Frontend

## Backend
- Database is fully functional and set up, and compatible with front end login page (includes first name, last name, username, and hashed password columns)
- Database is also configured/organized for the purposes of our OOTD application
  - The database contains tables for **users, images added by the user (items), and outfits that a user creates**
- Log in and logout functions have been created with **sessions** fully implemented (this allows user accounts to virtually be locked and restricts access to only that users account)-- we had some problems hooking up sessions between backend and frontend but these issues have been resolved
- CRUD API for images and outfits are fully functional **using GORM and Gorilla MUX router** and is running at localhost:9000
  - Similar to last sprint, this can be seen by making requests through Postman. Through Postman, we are able to access existing users and log them in. In addition we are able to create outfits using items (images) stored in the database. The ids of the items are saved and they are used to create the outfit. This way when users create an outfit, the outfit will be saved in the database and they will be able to access it later. 

# Unit Tests


## Frontend


## Backend
Unit tests are fully functional for every CRUD function used throughout the program
The unit tests for the backend can be found [here](https://github.com/gatorcloset/OOTD/blob/main/backend/go/src/github.com/user/user_test.go).

These tests were used to verify the functionality of outfit building, the premise of OOTD. We also conducted tests to determine if we are able to retrieve all of a user's clothing items and all of the user's items in a certain category. The tests perform the following on each of the tables:

- Checks that database can be accessed and requests can be made via [http://localhost:9000]
- Checks that new outfits can be created using Create Outfit (POST functionality)
- Checks that a single outfit can be retrieved using GetOutfit (GET functionality)
- Checks that all outfits can be retrieved using GetOutfits (GET functionality)
- Checks that outfits are able to be updated using UpdateOutfit (single/multiple items and name of outfit) (PUT functionality)
- Checks that outfits are able to be softly deleted (DELETE functionality)
- Checks that we are able to retrieve all of a user's items by user id
- Checks that we are able to retrieve all of a user's items under a certain category by user id

# Detailed Backend Documentation
Find our detailed backend documentation here -> https://github.com/gatorcloset/OOTD/blob/main/Sprint4BackendDocumentation.md
