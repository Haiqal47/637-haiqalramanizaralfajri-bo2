# MyGram API SPECIFICATION

## Users

### Register Users

- Method: `POST`
- Endpoint : `/users`
- Accept : `multipart/form-data`
- Content-Type: `application/json`
- Request :

```form-data
age:integer
email:string
password:string
username:string
profile_image:file
```

- Response:

```json
{
  "age": 0,
  "email": "string",
  "id": 0,
  "username": "string"
}
```

### Login Users

- Method: `POST`
- Endpoint : `/users`
- Accept : `application/json`
- Content-Type: `application/json`
- Request :

```json
{
  "email": "string",
  "password": "string"
}
```

- Response:

```json
{
  "token": "string"
}
```

### Update Users

- Authorization: Bearer \<Token\>
- Method: `PUT`
- Endpoint : `/users/{userId}`
- Accept : `multipart/form-data`
- Content-Type: `application/json`
- Request :

```form-data
email:string
username:string
profile_image:file
```

- Response:

```json
{
  "age": 0,
  "email": "string",
  "id": 0,
  "profile_image_url": "string",
  "updated_at": "string",
  "username": "string"
}
```

### Delete Users

- Authorization: Bearer \<Token\>
- Method: `DELETE`
- Endpoint : `/users/{userId}`
- Content-Type: `application/json`
- Response:

```json
{
  "message": "string"
}
```

## Photos

### Create Photos

- Authorization: Bearer \<Token\>
- Method: `POST`
- Endpoint : `/photos`
- Accept : `multipart/form-data`
- Content-Type: `application/json`
- Request :

```form-data
title:string
caption:string
photo:file
```

- Response:

```json
{
  "caption": "string",
  "created_at": "string",
  "id": 0,
  "photo_url": "string",
  "title": "string",
  "user_id": 0
}
```

### Get Photos

- Authorization: Bearer \<Token\>
- Method: `GET`
- Endpoint : `/photos`
- Content-Type: `application/json`
- Response:

```json
[
  {
    "User": {
      "email": "string",
      "username": "string"
    },
    "caption": "string",
    "created_at": "string",
    "id": 0,
    "photo_url": "string",
    "title": "string",
    "updated_at": "string",
    "user_id": 0
  }
]
```

### Update Photos

- Authorization: Bearer \<Token\>
- Method: `PUT`
- Endpoint : `/photos/{photoId}`
- Accept : `multipart/form-data`
- Content-Type: `application/json`
- Request :

```form-data
title:string
caption:string
photo:file
```

- Response:

```json
{
  "id": 0,
  "message": "string",
  "photo_id": 0,
  "updated_at": "string",
  "user_id": 0
}
```

### Delete Photos

- Authorization: Bearer \<Token\>
- Method: `DELETE`
- Endpoint : `/photos/{photoId}`
- Content-Type: `application/json`

- Response:

```json
{
  "message": "string"
}
```

## Comments

### Create Comments

- Authorization: Bearer \<Token\>
- Method: `POST`
- Endpoint : `/comments`
- Accept : `application/json`
- Content-Type: `application/json`
- Request :

```json
{
  "message": "string",
  "photo_id": 0
}
```

- Response:

```json
{
  "created_at": "string",
  "id": 0,
  "message": "string",
  "photo_id": 0,
  "user_id": 0
}
```

### Get Comments

- Authorization: Bearer \<Token\>
- Method: `GET`
- Endpoint : `/comments`
- Content-Type: `application/json`
- Response:

```json
[
  {
    "Photo": {
      "caption": "string",
      "id": 0,
      "photo_url": "string",
      "title": "string",
      "user_id": 0
    },
    "User": {
      "email": "string",
      "id": 0,
      "username": "string"
    },
    "created_at": "string",
    "id": 0,
    "message": "string",
    "photo_id": 0,
    "updated_at": "string",
    "user_id": 0
  }
]
```

### Update Comments

- Authorization: Bearer \<Token\>
- Method: `PUT`
- Endpoint : `/comments/{commentId}`
- Accept : `application/json`
- Content-Type: `application/json`
- Request :

```json
{
  "message": "string"
}
```

- Response:

```json
{
  "id": 0,
  "message": "string",
  "photo_id": 0,
  "updated_at": "string",
  "user_id": 0
}
```

### Delete Comments

- Authorization: Bearer \<Token\>
- Method: `DELETE`
- Endpoint : `/comments/{commentId}`
- Content-Type: `application/json`

- Response:

```json
{
  "message": "string"
}
```

## Social Medias

### Create Social Medias

- Authorization: Bearer \<Token\>
- Method: `POST`
- Endpoint : `/socialmedias`
- Accept : `application/json`
- Content-Type: `application/json`
- Request :

```json
{
  "name": "string",
  "social_media_url": "string"
}
```

- Response:

```json
{
  "created_at": "string",
  "id": 0,
  "name": "string",
  "social_media_url": "string",
  "user_id": 0
}
```

### Get Social Medias

- Authorization: Bearer \<Token\>
- Method: `GET`
- Endpoint : `/socialmedias`
- Content-Type: `application/json`
- Response:

```json
[
  {
    "User": {
      "id": 0,
      "profile_image_url": "string",
      "username": "string"
    },
    "created_at": "string",
    "id": 0,
    "name": "string",
    "social_media_url": "string",
    "updated_at": "string",
    "user_id": 0
  }
]
```

### Update Social Medias

- Authorization: Bearer \<Token\>
- Method: `PUT`
- Endpoint : `/socialmedias/{socialMediaId}`
- Accept : `application/json`
- Content-Type: `application/json`
- Request :

```json
{
  "name": "string",
  "social_media_url": "string"
}
```

- Response:

```json
{
  "id": 0,
  "name": "string",
  "social_media_url": "string",
  "updated_at": "string",
  "user_id": 0
}
```

### Delete Social Medias

- Authorization: Bearer \<Token\>
- Method: `DELETE`
- Endpoint : `/socialmedias/{socialMediaId}`
- Content-Type: `application/json`

- Response:

```json
{
  "message": "string"
}
```
