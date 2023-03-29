# Deatiled Backend Documentation

Our API provides create, read, update, and delete (CRUD) functionalities for storing user, tag, and item information in a SQLite database using GoLang and GORM. This database, called "OOTD.db", houses a user's first and last name along with a username, hashed (encrypted) password, and a timestamp referring to when each entry was last accessed or modified. It also stores user-uploaded clothing items, the category they fall in, image, and image tags. We also implemented functionality to log a user in, which checks to see if the user attempting to log in exists by verifying their username and password against entries in the database, ensuring that users would only be able to access their own information.

Our vision for the database is to have our four tables house user-specific images and tags which they will be able to access when they login to OOTD.

## Endpoints

The endpoints created for this sprint include:

- POST /login - authenticates a user and creates a session
- POST /logout - logs out a user and clears session information
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

**For item table (which stores user-uploaded images of clothing)**

- `POST /item` 

 Sample Request Format:
  - Input is a multipart form with the following information:
    - name
    - category
    - image file
    
  Sample Response Format:
 ````
  ```
   {
   "ID" : 9,
   "CreatedAt" : "2023-03-29T12:47:36.074678-04:00"
   "DeletedAt" : null,
   "user_id" : 0,
   "name" : "\"name\"",
   "category" : "\"category\"",
   "image" : "images/6fbc7ac-60eb-4c8f-81c3-425e2fe37ce4.jpg" 
   }
   ```
  ````
  
- `GET /items`
No additional parameters, data format is JSON
  - Sample Request Format:
    - Nothing needed
  - Sample Response Format (Item list is truncated for the purposes of this demonstration):
  
  ````
  ```
   {
   "ID" : 9,
   "CreatedAt" : "2023-03-27T12:47:10.074678-04:00"
   "DeletedAt" : null,
   "user_id" : 0,
   "name" : "\"name\"",
   "category" : "\"category\"",
   "image" : "images/6abd7ac-60eb-4c8f-81c3-425e2fe37ce4.jpg" 
   }
   
    {
   "ID" : 10,
   "CreatedAt" : "2023-03-29T12:47:35.074678-04:00"
   "DeletedAt" : null,
   "user_id" : 0,
   "name" : "\"name\"",
   "category" : "\"category\"",
   "image" : "images/6fbc7ac-60eb-4c8f-81c3-425e2fwkjee37ce4.jpg" 
   }
   
    {
   "ID" : 11,
   "CreatedAt" : "2023-03-29T12:47:36.074678-04:00"
   "DeletedAt" : null,
   "user_id" : 0,
   "name" : "\"name\"",
   "category" : "\"category\"",
   "image" : "images/6fbac-60eb-4c8f-81djec3-425e2fe37ce4.jpg" 
   }
   ```
  ````
  
  - `GET /item{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
      - Nothing needed
    - Sample Response Format:
 ````
  ```
    {
   "ID" : 9,
   "CreatedAt" : "2023-03-27T12:47:10.074678-04:00"
   "DeletedAt" : null,
   "user_id" : 0,
   "name" : "\"name\"",
   "category" : "\"category\"",
   "image" : "images/6abd7ac-60eb-4c8f-81c3-425e2fe37ce4.jpg" 
   }
```
````

- `PUT /item{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
   - Input is an updated multipart form with at least one change to the following information:
    - name
    - category
    - image file

    - Sample Response Format:
  ````
  ```
    {
   "ID" : 9,
   "CreatedAt" : "2023-03-27T12:47:10.074678-04:00"
   "DeletedAt" : null,
   "user_id" : 0,
   "name" : "\"name\"",
   "category" : "\"category\"",
   "image" : "images/6abd7ac-60eb-4c8f-81c3-425e2fe37ce41111111.jpg" 
   }
  
```
````

- `DELETE /item{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
      - Nothing needed
    - Sample Response Format:
    "The item has successfully been deleted."
    
**For tag table (which stores tags associated to images)**



## Error Handling

If an error is encountered, the API will return the following error codes:
- 400 - Bad Request
- 401 - Unauthorized
- 403 - Forbidden
- 404 - Not Found
- 500 - Internal Server Error

## Backend Unit Tests for CRUD functions

The backend unit tests for CRUD functions can be found [here](
