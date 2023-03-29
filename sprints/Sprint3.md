# Completed Work

[Sprint 3 Demo]

## Frontend

## Backend
- Database is fully functional and set up, and compatible with front end login page (includes first name, last name, username, and hashed password columns)
- Database is also configured/organized for the purposes of our OOTD application
  - The database contains tables for **users, images added by the user (items), all the tags related to images, and image specific tags**
  - In the user table, we implemented password hashing. This can be seen when viewing our DB user table. The fields under the password category are securely stored as a series of random letters and numbers
- Log in and logout functions have been created with **sessions** fully implemented (this allows user accounts to virtually be locked and restricts access to only that users account)
- CRUD API for images, tags, and item_tags are fully functional **using GORM and Gorilla MUX router** and is running at localhost:9000
  - Similar to last sprint, this can be seen by making requests through Postman. Through Postman, we are able to access existing users and log them in. Their information is returned and sent to the front end, such that users are able to view their closet and personal information when they log in.

# Unit Tests

## Frontend
The Cypress testing for the frontend can be found [here]

## Backend
The unit tests for the backend can be found [here] (https://github.com/gatorcloset/OOTD/blob/main/backend/go/src/github.com/user/user_test.go)
These tests were used to unit test the database CRUD functionality for our tag and item tables. The tests perform the following on each of the tables:

Checks that database can be accessed and requests can be made via [http://localhost:9000]
Checks that new tags and items can be created (POST functionality)
Checks that entries can be read (GET functionality for one user and entire list of users, tags, and items)
Checks that tag and item information can be updated (UPDATE functionality)
Checks that tags and items can be deleted (DELETE functionality)

# Detailed Backend Documentation
Find our detailed backend documentation here -> https://github.com/gatorcloset/OOTD/blob/main/Sprint3BackendDocumentation.md
