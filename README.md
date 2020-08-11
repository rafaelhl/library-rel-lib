## Sample project to validate the [REL lib](https://fs02.github.io/rel/#/) as database layer

This project implements a book library with four operations:

 - POST /books (Insert a new book on the library)
 - GET /books/{bookID} (Get a book by your ID on the library)
 - PUT /books/{bookID} (Update a book on the library)
 - GET /books (Get all books on the library)
 
## How to run

### Migration

To build the database you have to run the SQL migration on your local MySQL database:
  - [20200720053810_create_library_tables.sql](./migrations/20200720053810_create_library_tables.sql)

### Run

To up the application only use the following Go command:

```shell script
    go run .
```
  
To validate if all up ok send an HTTP GET to `/ping` route and expect a text `pong`:

```shell script
    curl http://localhost:8080/ping   
```
