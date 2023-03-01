# Completed Work
## Frontend
- All frontend components are being displayed through a mock database, rather than hard coded.
- All mock data is injected into the needed components using Angular **services**.
- On the "/closet/{clothing-category}" pages, {clothing-category} is generated based on the name of the clothing category the user selects.
- Closet items are filtered according to the activated route the user is visiting (as further described in the "Frontend Unit Tests" section).
- The Cypress framework was set up and tests have been configured. These tests are 1:1 with the added features described above.

## Backend
- Database is fully functional and set up, compatible with front end login page (includes first name, last name, username, and password columns)
- CRUD API is fully functional **using GORM and Gorilla MUX router** and is running at localhost:9000.
  - This can be seen by making requests through Postman. Through Postman, requests are made to create new entries in the database (for new users), read existing information, update entries (if user decides to change login information), and delete entries (user deletes account). 
  - Information is automigrated into a User table in database; later more tables will be created to house a closet containing tags and images of clothing
- The unit tests have been set up and configured. The tests are 1:1 with the CRUD functionality described above.

# Unit Tests
## Frontend
The Cypress testing for the frontend can be found [here](https://github.com/gatorcloset/OOTD/blob/closet/cypress/e2e/closet.cy.ts).

The Cypress framework was used to unit test the "Closet" page. The tests perform the following:
- Checks that http://localhost:4200/closet can be successfully visited; verifies that routing is correct.
- Checks the functionality of each clothing category card; verifies that routing to /closet/jeans is performing properly.
- Checks that closet items are filtering properly according to the activiated router. Specifically, the test checks that if the user clicks the "Jeans" category, they should be redirected to a page that only displays clothing cards under the "Jean" category.

## Backend
The unit tests for the backend can be found [here](https://github.com/gatorcloset/OOTD/blob/main/backend/go/src/github.com/user/user_test.go)

These tests were used to unit test the database CRUD functionality. The tests perform the following:
- Checks that database can be accessed and requests can be made via [http://localhost:9000]
- Checks that new users can be created (POST functionality)
- Checks that entries can be read (GET functionality for one user and entire list of users)
- Checks that user information can be updated (UPDATE functionality)
- Checks that accounts can be deleted (DELETE functionality)

# API Documentation
Find our backend CRUD API Documentation here -> https://github.com/gatorcloset/OOTD/blob/main/backendAPIDocumentation.md
