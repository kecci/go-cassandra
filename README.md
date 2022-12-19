# go-cassandra

This is experimental golang project with database cassandra.

## Feature Specification

- cassandra:4.1.0
- golang:1.17
- gocql:v1.3.1
- go-chi:v5.0.8

## Folder Structure

```
.
├── database
├── handler
├── model
├── repository
└── router
```

Functionality:

- `databases`: database connector for cluster and session, also wrapper function for execution query.
- `handler`: a handler for function between routes and repositories.
- `model`: a model for struct of database model and request model.
- `repository`: a repository for function query to the database.
- `router`: a router is list of routes HTTP API.

## Preparation

### Cassandra with Docker compose

You can following installation with this link:
https://github.com/kecci/docker-kit/tree/master/cassandra

## API

### 1. POST /user - create a new user

Request:

```sh
curl --location --request POST 'http://localhost:3000/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email" : "kecci1@mail.com",
    "first_name" : "kecci",
    "last_name" : "1"
}'
```

Response:

```sh
"User added"
```

Database:

```sh
cqlsh:first_app> select * from users;

 email           | first_name | last_name
-----------------+------------+-----------
 kecci1@mail.com |      kecci |         1

(1 rows)
```

### 2. GET /user - get a user

Request:
Response:
Database:
