# gin-cat-backend

This is a Go API for managing a collection of cats. It includes the following endpoints:

## Endpoints

### Cats

```
GET /cats
GET /cats/:id
POST /cats
PUT /cats/:id
DELETE /cats/:id
```

## Running the API

To start the API, run the following command from the root directory:

```bat
go run main.go
```

The API will be served at http://localhost:8080.

## Testing

To run the tests for the API, run the following command from the root directory:

```
go test ./.../tests
```

## Update Swagger

```
swag init --parseDependency --parseInternal
```
