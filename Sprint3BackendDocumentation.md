# Backend CRUD API Documentation

Our API provides create, read, update, and delete (CRUD) functionalities for storing user, tag, and item information in a SQLite database using GoLang and GORM. This database, called OOTD.db, houses a user's first and last name along with a username, password, and a timestamp referring to when each entry was last accessed or modified. It also houses an item's name, category, image, and image ID and a relevant tag's name and tag ID for each image. We also implemented functionality to log a user in, which checks to see if the user attempting to log in exists by checking their username and password. 

## Endpoints

The endpoints created include:

- POST /login - log in a user

- POST /item - create an item
- GET /item - retrieve a list of all tags
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
