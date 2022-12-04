# SurrealDB Go Example Template

Testing the surrealdb Go library.

### Requirements

Install surrealdb.

### How to run

Start surrealdb database by running the bash script/make file:

```bash
make start-db
```

To run the Go server, run the following:

```bash
make start-server
```


# Commands

```bash
go build
go test
go run ./cmd
```

# Commands

Example server just has three handlers:
 - health-check
 - cafes (post)
 - cafes (get)

Work in progress!

To test /cafes post endpoint use JSON payload with the following structure:

```bash
{
    "cafes": [
        {
            "name": "New cafe",
            "coffee": "warm",
            "chairs": "none"
        },
        {
            "name": "Bob's cafe",
            "coffee": "cold",
            "chairs": "Plentiful"
        }
    ]
}
```

# Note

Using specific version of SurrealDB. To get version of surrealdb use follow command:

```bash
go get github.com/surrealdb/surrealdb.go@{commithash}
```
