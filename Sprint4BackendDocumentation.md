
# Detailed Backend Documentation

Our API provides create, read, update, and delete (CRUD) functionalities for storing user, item and outfit information in a SQLite database using GoLang and GORM. This database, called "OOTD.db", houses a user's first and last name along with a username, hashed (encrypted) password, and a timestamp referring to when each entry was last accessed or modified. It also stores user-uploaded clothing items, the category they fall in, image, and image tags. We also implemented functionality to create outfits by taking items that already exist in the database and allowing users to select one top, bottom, accessory, one-piece, and shoe combo (some of these fields can also be null). When users select a combination of items they like, they are able to title the outfit and save it in the database to refer back to later!

Our vision for the database is to have our three tables house user-specific images and outfits which they will be able to access and modify when they login to OOTD.

## Endpoints

The endpoints created for this sprint include:

- POST /login - authenticates a user and fixed session handling
- POST /outfit - create an outfit
- GET /outfit - retrieve a list of all outfits
- GET /outfit/{id} - retrieve information for a specific outfit
- PUT /outfit/{id} - update an existing outfit
- DELETE /outfit/{id} - delete an existing outfit
- GET /users/{id}/items- get all items for a specific user
- GET /users/{id}/category/{name}- get all items of a specific category for a specific user

## Request and Response Formats



## Error Handling

If an error is encountered, the API will return the following error codes:
- 400 - Bad Request
- 401 - Unauthorized
- 403 - Forbidden
- 404 - Not Found
- 500 - Internal Server Error

## Backend Unit Tests for CRUD functions

The backend unit tests for CRUD functions can be found [here](https://github.com/gatorcloset/OOTD/blob/main/backend/go/src/github.com/user/user_test.go)
