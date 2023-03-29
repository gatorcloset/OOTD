# Deatiled Backend Documentation

Our API provides create, read, update, and delete (CRUD) functionalities for storing user, tag, and item information in a SQLite database using GoLang and GORM. This database, called "OOTD.db", houses a user's first and last name along with a username, hashed (encrypted) password, and a timestamp referring to when each entry was last accessed or modified. It also stores user-uploaded clothing items, the category they fall in, image, and image tags. We also implemented functionality to log a user in, which checks to see if the user attempting to log in exists by verifying their username and password against entries in the database, ensuring that users would only be able to access their own information.

Our vision for the database is to have our four tables house user-specific images and tags which they will be able to access when they login to OOTD.

## Endpoints

The endpoints created for this sprint include:

- POST /login - log in a user
- POST /item - create an item
- GET /item - retrieve a list of all items
- GET /item/{id} - retrieve information for a specific item
- PUT /item/{id} - update an existing item
- DELETE /item/{id} - delete an existing item
- POST /tag - create a tag
- GET /tag - retrieve a list of all tags
- GET /tag/{id} - retrieve information for a specific tag
- PUT /tag/{id} - update an existing tag
- DELETE /tag/{id} - delete an existing tag

## Request and Response Formats 

## Error Handling

If an error is encountered, the API will return the following error codes:
- 400 - Bad Request
- 401 - Unauthorized
- 403 - Forbidden
- 404 - Not Found
- 500 - Internal Server Error

## Backend Unit Tests for CRUD functions

The backend unit tests for CRUD functions can be found here
