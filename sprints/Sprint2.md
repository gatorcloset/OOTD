# Completed Work
## Frontend
- All frontend components are being displayed through a mock database, rather than hard coded.
- All mock data is injected into the needed components using Angular **services**.
- On the "/closet/{clothing-category}" pages, {clothing-category} is generated based on the name of the clothing category the user selects.
- Closet items are filtered according to the activated route the user is visiting (as further described in the "Frontend Unit Tests" section).
- The Cypress framework was set up and tests have been configured. These tests are 1:1 with the added features described above.

## Backend

# Unit Tests
## Frontend
The Cypress testing for the frontend can be found [here](https://github.com/gatorcloset/OOTD/blob/closet/cypress/e2e/closet.cy.ts).

The Cypress framework was used to unit test the "Closet" page. The tests perform the following:
- Checks that http://localhost:4200/closet can be successfully visited; verifies that routing is correct.
- Checks the functionality of each clothing category card; verifies that routing to /closet/jeans is performing properly.
- Checks that closet items are filtering properly according to the activiated router. Specifically, the test checks that if the user clicks the "Jeans" category, they should be redirected to a page that only displays clothing cards under the "Jean" category.

## Backend Unit Tests
(write in a separate folder/file and link here)

# API Documentation
(write in a separate folder/file and link here)