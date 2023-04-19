
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

**For outfit table (which stores collections of items as outfits)**

- `POST /outfit` 

 Sample Request Format:
  - Input is a multipart form with the following information:
    - Name
    - Tops 
    - Bottoms
    - OnePieces
    - Accessories:
    - Shoes
    
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
  
- `GET /outfits`
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
  
  - `GET /outfit{id}`
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

- `PUT /outfit{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
   - Input is an updated multipart form with at least one change to the following information:
    - Name
    - Tops 
    - Bottoms
    - OnePieces
    - Accessories:
    - Shoes

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

## Error Handling

If an error is encountered, the API will return the following error codes:
- 400 - Bad Request
- 401 - Unauthorized
- 403 - Forbidden
- 404 - Not Found
- 500 - Internal Server Error

## Backend Unit Tests for CRUD functions

The backend unit tests for CRUD functions can be found [here](https://github.com/gatorcloset/OOTD/blob/main/backend/go/src/github.com/user/user_test.go)
