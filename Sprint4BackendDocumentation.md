
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
  - Input is a request to JSON:
  ````
  ```
    {
  "Name": "Pink Shoes Outfit",
  "TopID": 1,
  "BottomID": 2,
  "ShoesID": 4
  }
 ```
  ````
    
  Sample Response Format:
 ````
  ```
   {
    "ID": 2,
    "CreatedAt": "2023-04-17T21:33:58.218509-04:00",
    "UpdatedAt": "2023-04-17T21:46:41.734317-04:00",
    "DeletedAt": null,
    "Name": "Pink Shoes Outfit",
    "Tops": {
        "ID": 1,
        "CreatedAt": "2023-04-17T21:29:21.646798-04:00",
        "UpdatedAt": "2023-04-17T21:29:21.646798-04:00",
        "DeletedAt": null,
        "user_id": 1,
        "name": "black tank top",
        "category": "tops",
        "image": "assets/item-images/72dcbe2e-89c2-4d56-a613-37f87d5e3673.jpeg"
    },
    "TopID": 1,
    "Bottoms": {
        "ID": 2,
        "CreatedAt": "2023-04-17T21:32:35.909144-04:00",
        "UpdatedAt": "2023-04-17T21:32:35.909144-04:00",
        "DeletedAt": null,
        "user_id": 2,
        "name": "cargo pants",
        "category": "bottoms",
        "image": "assets/item-images/3df45954-7108-4214-a0c7-318876bb83e4.jpeg"
    },
    "BottomID": 2,
    "OnePieces": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "OnePieceID": 0,
    "Accessories": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "AccessoriesID": 0,
    "Shoes": {
        "ID": 4,
        "CreatedAt": "2023-04-17T21:39:05.439145-04:00",
        "UpdatedAt": "2023-04-17T21:39:05.439145-04:00",
        "DeletedAt": null,
        "user_id": 4,
        "name": "pink jordans",
        "category": "shoes",
        "image": "assets/item-images/84442d2a-b3b3-4cd7-85b0-b01ae1d0381f.jpeg"
    },
    "ShoesID": 4
}
   ```
  ````
  
- `GET /outfit`
No additional parameters, data format is JSON
  - Sample Request Format:
    - Nothing needed
  - Sample Response Format (Item list is truncated for the purposes of this demonstration):
  ````
  ```
  
  {
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "user_id": 0,
    "Name": "New Outfit",
    "Tops": {
        "ID": 1,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "TopID": 0,
    "Bottoms": {
        "ID": 2,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "BottomID": 0,
    "OnePieces": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "OnePieceID": 0,
    "Accessories": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "AccessoriesID": 0,
    "Shoes": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "ShoesID": 4
}

{
    "ID": 3,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "user_id": 0,
    "Name": "",
    "Tops": {
        "ID": 1,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "TopID": 0,
    "Bottoms": {
        "ID": 5,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "BottomID": 0,
    "OnePieces": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "OnePieceID": 0,
    "Accessories": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "AccessoriesID": 0,
    "Shoes": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "ShoesID": 3
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
    "ID": 0,
    "CreatedAt": "0001-01-01T00:00:00Z",
    "UpdatedAt": "0001-01-01T00:00:00Z",
    "DeletedAt": null,
    "user_id": 0,
    "Name": "",
    "Tops": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "TopID": 0,
    "Bottoms": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "BottomID": 0,
    "OnePieces": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "OnePieceID": 0,
    "Accessories": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "AccessoriesID": 0,
    "Shoes": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "ShoesID": 0
}
```
````

- `PUT /outfit{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
   - Input is a request to JSON:
  ```
    {
  "Name": "Pink Shoes Outfit",
  "TopID": 1,
  "BottomID": 2,
  "ShoesID": 5
    }
  ```
  
    - Sample Response Format:
  ```
    {
    "ID": 2,
    "CreatedAt": "2023-04-17T21:33:58.218509-04:00",
    "UpdatedAt": "2023-04-17T21:46:41.734317-04:00",
    "DeletedAt": null,
    "Name": "Pink Shoes Outfit",
    "Tops": {
        "ID": 1,
        "CreatedAt": "2023-04-17T21:29:21.646798-04:00",
        "UpdatedAt": "2023-04-17T21:29:21.646798-04:00",
        "DeletedAt": null,
        "user_id": 1,
        "name": "black tank top",
        "category": "tops",
        "image": "assets/item-images/72dcbe2e-89c2-4d56-a613-37f87d5e3673.jpeg"
    },
    "TopID": 1,
    "Bottoms": {
        "ID": 2,
        "CreatedAt": "2023-04-17T21:32:35.909144-04:00",
        "UpdatedAt": "2023-04-17T21:32:35.909144-04:00",
        "DeletedAt": null,
        "user_id": 2,
        "name": "cargo pants",
        "category": "bottoms",
        "image": "assets/item-images/3df45954-7108-4214-a0c7-318876bb83e4.jpeg"
    },
    "BottomID": 2,
    "OnePieces": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "OnePieceID": 0,
    "Accessories": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "user_id": 0,
        "name": "",
        "category": "",
        "image": ""
    },
    "AccessoriesID": 0,
    "Shoes": {
        "ID": 5,
        "CreatedAt": "2023-04-17T21:39:05.439145-04:00",
        "UpdatedAt": "2023-04-17T21:39:05.439145-04:00",
        "DeletedAt": null,
        "user_id": 4,
        "name": "pink jordans",
        "category": "shoes",
        "image": "assets/item-images/84442d2a-b3b3-4cd7-85b0-b01ae1d0381f.jpeg"
    },
    "ShoesID": 5
  }
  ```

- `DELETE /outfit{id}`
  Requires an id, data format is in JSON
    - Sample Request Format:
      - Nothing needed
    - Sample Response Format:
    "The outfit been deleted"

## Error Handling

If an error is encountered, the API will return the following error codes:
- 400 - Bad Request
- 401 - Unauthorized
- 403 - Forbidden
- 404 - Not Found
- 500 - Internal Server Error

## Backend Unit Tests for CRUD functions

The backend unit tests for CRUD functions can be found [here](https://github.com/gatorcloset/OOTD/blob/main/backend/go/src/github.com/user/user_test.go)
