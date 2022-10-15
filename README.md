## Simple Golang Goa Design and Redis Auth


### Quick Start

````
goa gen auth/design
go run auth/cmd/auth
````

### Redis

````
docker pull redis
docker run --name redis -p 6379:6379 -d redis:latest
````

### Curl

````
curl --location --request POST 'http://localhost:8080/auth/authentication/auth?X-Provider=1' \
--header 'accept: application/json' \
--header 'Content-Type: application/json' \
--data-raw '{
  "login": "login1",
  "password": "password",
  "privacyAccepted": true,
  "uuid": "eFd1CEdC-93Fc-38db-dAae-e029817F045F"
}'
````

## Gin rewrite response

### Res
````
{
"brand": "Apple",
"category": "smartphones",
"description": "An apple mobile which is nothing like apple",
"discountPercentage": 12.96,
"id": 1,
"images": [
"https://dummyjson.com/image/i/products/1/1.jpg",
"https://dummyjson.com/image/i/products/1/2.jpg",
"https://dummyjson.com/image/i/products/1/3.jpg",
"https://dummyjson.com/image/i/products/1/4.jpg",
"https://dummyjson.com/image/i/products/1/thumbnail.jpg"
],
"price": 549,
"rating": 4.69,
"stock": 94,
"thumbnail": "https://dummyjson.com/image/i/products/1/thumbnail.jpg",
"title": "iPhone 9",
"uuid": "abe21b5a-83d5-4f13-bd4f-2833216c9099"
}
````