# go-api-server

#### GO "minimalistic" API Server

To run the server, 
```
go run server.go
```

API Routes,
```
* /post (GET) -> All posts in the phonebook document (database)
* /post/{id} (GET) -> Get a single post
* /post/{id} (POST) -> Create a new post
* /post/{id} (DELETE) -> Delete a post
```