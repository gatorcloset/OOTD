# Completed Work

[Sprint 3 Demo]

## Frontend

## Backend
- Database is fully functional and set up, compatible with front end login page (includes first name, last name, username, and hashed password columns)
- Database is also configured/organized for the purposes of the OOTD application
  - The database contains tables for users, images added by the user (items), all tags related to images, and image specific tags
  - In the user table, the passwords have been hashed. This can be seen because the fields under the password category are stored as a series of random letters and numbers
- Log in and logout functions have fully been created with **sessions** fully implemented (this allows user accounts to virtually be locked and restricts access to only that users account)
- CRUD API is fully functional **using GORM and Gorilla MUX router** and is running at localhost:9000
  - Similar to last sprint, this can be seen by making requests through Postman. Through Postman, we are able to access existing users and log them in. Their information is returned and sent to the front end, such that users are able to view their closet and personal information when they log in.

# Unit Tests


## Frontend
The Cypress testing for the frontend can be found [here]

## Backend
The unit tests for the backend can be found [here]

# Detailed Backend Documentation
Find our detailed backend documentation here -> 
