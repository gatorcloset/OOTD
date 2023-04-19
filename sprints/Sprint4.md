# Completed Work

[Sprint 4 Video Demo]()

## Frontend
- Added more login/signup validation:
  - Sign up and log in buttons are disabled until all validation errors are handled
  - Sign up page has validation relating to each field (improper name/username format, all fields must be filled out, etc.)
- Implemented guards: Unauthenticated users can only access the entrance, login, and sign up pages
- Authenticated users only see clothing items, outfits, and user details pertaining to their own personal user profile
  - With the help of authentication and guards, all clothing items are now sourced from the database rather than hard-coded in
- Added functionality to create, edit, and delete an outfit
- Added functionality to edit and delete a clothing item
- Final UI additions and adjustments were made to clean up the appearance of the application

## Backend
- Database is fully functional and set up, and compatible with front end login page (includes first name, last name, username, and hashed password columns)
- Database is also configured/organized for the purposes of our OOTD application
  - The database contains tables for **users, images added by the user (items), and outfits that a user creates**
- Log in and logout functions have been created with **sessions** fully implemented (this allows user accounts to virtually be locked and restricts access to only that users account)-- we had some problems hooking up sessions between backend and frontend but these issues have been resolved
- CRUD API for images and outfits are fully functional **using GORM and Gorilla MUX router** and is running at localhost:9000
  - Similar to last sprint, this can be seen by making requests through Postman. Through Postman, we are able to access existing users and log them in. In addition we are able to create outfits using items (images) stored in the database. The ids of the items are saved and they are used to create the outfit. This way when users create an outfit, the outfit will be saved in the database and they will be able to access it later. 

# Unit Tests


## Frontend
Cypress unit test can be found [here](https://github.com/gatorcloset/OOTD/tree/main/cypress/e2e)
- All unit tests were revised to the updated application. Since we are now using guards, we had to make sure the user was authenticated on all unit tests before proceeding to the actual functionality testing. In addition to the previous unit tests, we added the following:
- Tests ability to add a new clothing item
- Tests ability to edit image, name, and category of clothing item
- Tests ability to create a new outfit
- Tests ability to edit outfit name and outfit clothing items

## Backend
Unit tests are fully functional for every CRUD function used throughout the program
The unit tests for the backend can be found [here](https://github.com/gatorcloset/OOTD/blob/main/backend/go/src/github.com/user/user_test.go).

These tests were used to verify the functionality of outfit building, the premise of OOTD. We also conducted tests to determine if we are able to retrieve all of a user's clothing items and all of the user's items in a certain category. The tests perform the following on each of the tables:

- Checks that database can be accessed and requests can be made via [http://localhost:9000]
- Checks that new outfits can be created using Create Outfit (POST functionality)
- Checks that a single outfit can be retrieved using GetOutfit (GET functionality)
- Checks that all outfits can be retrieved using GetOutfits (GET functionality)
- Checks that outfits are able to be updated using UpdateOutfit (single/multiple items and name of outfit) (PUT functionality)
- Checks that outfits are able to be softly deleted using DeleteOutfit (DELETE functionality)
- Checks that we are able to retrieve all of a user's items by user id using GetAllUserItems (GET functionality)
- Checks that we are able to retrieve all of a user's items under a certain category by user id using GetAllItemsCategory (GET functionality)

# Detailed Backend Documentation
Find our detailed backend documentation here -> https://github.com/gatorcloset/OOTD/blob/main/Sprint4BackendDocumentation.md
