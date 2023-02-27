# Backend CRUD API Documentation

Our API provides create, read, update, and delete (CRUD) functionalities for storing user information in a SQLite database using GoLang and GORM. This database, called OOTD.db, houses a user's first and last name along with a username, password, and a timestamp referring to when each entry was last accessed or modified. Later, it will also house images of the user's closet and relevant tags for each image.

## Endpoints

The endpoints created include:
- POST /users - create a new user account
- GET /users - retrieve a list of all user accounts
- GET /users/{id} - retrieve information for a specific user account
- PUT /users/{id} - update an existing user account
- DELETE /user/{id} - delete an existing user account

## Request and Response Formats 
- `POST /users` 
No additional parameters, data format is JSON
  - Sample Request Format :
  ````
  ```
  {
    "firstname": "Megan",
    "lastname": "Shah",
    "username": "megan.shah07@gmail.com",
    "Password": "samplePassword!"
  }
  ```
  ````
  
  - Sample Response Format:
  ````
  ```
  {
    "ID": 6,
    "CreatedAt": "2023-02-24T17:20:46.634953-05:00",
    "UpdatedAt": "2023-02-24T17:20:46.634953-05:00",
    "DeletedAt": null,
    "firstname": "Megan",
    "lastname": "Shah",
    "username": "megan.shah07@gmail.com",
    "password": "samplePassword!"
  }
  ```
  ````
  
- `GET /users`
No additional parameters, data format is JSON
  - Sample Request Format:
    - Nothing needed
  - Sample Response Format (User list is truncated for the purposes of this demonstration):
  ````
  ```
  {
        "ID": 1,
        "CreatedAt": "2023-02-23T18:08:16.701245-05:00",
        "UpdatedAt": "2023-02-23T18:08:16.701245-05:00",
        "DeletedAt": null,
        "firstname": "Michelle",
        "lastname": "Taing",
        "username": "michelletaing123",
        "password": "testPassword!"
    },
    {
        "ID": 2,
        "CreatedAt": "2023-02-23T18:09:00.078114-05:00",
        "UpdatedAt": "2023-02-23T18:09:00.078114-05:00",
        "DeletedAt": null,
        "firstname": "Natalie",
        "lastname": "Hodnett",
        "username": "nhodnett1",
        "password": "checkPassword*"
    }
```
````
    
- `GET /users{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
      - Nothing needed
    - Sample Response Format:
 ````
  ```
    {
    "ID": 5,
    "CreatedAt": "2023-02-24T17:19:54.804819-05:00",
    "UpdatedAt": "2023-02-24T17:19:54.804819-05:00",
    "DeletedAt": null,
    "firstname": "Aneesha",
    "lastname": "Acharya",
    "username": "aneesha.acharya",
    "password": "differentPassword!"
    }
```
````
    
- `PUT /users{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
  ````
  ```
    {
    "firstname": "This",
    "lastname": "Is Updated",
    "username": "megan.shah07@gmail.com",
    "Password": "samplePassword!"
    }
```
````
    - Sample Response Format:
  ````
  ```
    {
    "ID": 6,
    "CreatedAt": "2023-02-24T17:20:46.634953-05:00",
    "UpdatedAt": "2023-02-27T14:44:38.119619-05:00",
    "DeletedAt": null,
    "firstname": "This",
    "lastname": "Is Updated",
    "username": "megan.shah07@gmail.com",
    "password": "samplePassword!"
  }
  
```
````
  
- `DELETE /users{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
      - Nothing needed
    - Sample Response Format:
    "The user has successfully been deleted."
    
## Error Handling

If an error is encountered, the API will return the following error codes:
- 400 - Bad Request
- 401 - Unauthorized
- 403 - Forbidden
- 404 - Not Found
- 500 - Internal Server Error
