# Backend CRUD API Documentation

Our API provides create, read, update, and delete (CRUD) functionalities for storing user information in a SQLite database using GoLang and GORM. This database, called OOTD.db, houses a user's first and last name along with a username, password, and a timestamp referring to when each entry was last accessed or modified. Later, it will also house images of the user's closet and relevant tags for each image.

## Endpoints

The endpoints created include:
- POST /users - create a new user account
- GET /users - retrieve a list of all user accounts
- GET /users/:id - retrieve information for a specific user account
- PUT /users/:id - update an existing user account
- DELETE /user/:id - delete an existing user account

## Request and Response Formats 
`
