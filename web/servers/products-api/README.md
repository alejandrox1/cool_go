# [Building and Testing a REST API in Go with Gorilla Mux and PostgreSQL](https://semaphoreci.com/community/tutorials/building-and-testing-a-rest-api-in-go-with-gorilla-mux-and-postgresql)

## API specification

- Create a new product in response to a valid POST request at `/product`
- Fetch a list of products in response to a valid GET request at `/products`
- Fetch a product in response to a valid GET request at `/product/{id}`
- Update a product in response to a valid PUT request at `/product/{id}`
- Delete a product in response to a valid DELETE request at `/product/{id}`

### Creating a product
```
$ curl -d '{"name":"test product", "price":11.22}' http://localhost:8000/products

{"id":1,"name":"test product","price":11.22}
```

## Listing all products
```
$ curl http://localhost:8000/products

[{"id":1,"name":"test product","price":11.22}]
```

### Updating a product
```
$ curl -d '{"name":"test product - NEW", "price":111.22}' -X PUT http://localhost:8000/products/1

{"id":1,"name":"test product - NEW","price":111.22}
```
