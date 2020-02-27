# graphql-go-todo

Example graphql server project using a custom [graphql-go](https://github.com/guzmanweb/graphql-go) library.

## Usage

The project has two ways of being deployed.

### Locally
To deploy it by hand we must first create an environment variable with our postgres DB string connection.

```console
export PSQL_CONN_STR="user=USER dbname=DBNAME password=PASSWORD host=HOST port=PORT  sslmode=disable"
```
Once the environment variable is exported, we go to the root of the project and execute: 

```console
go run main.go 
```

After that the server should be started on port 3000.

### Docker

To deploy the project using Docker we must build the image, the easiest way would be to go to the root of the project and execute the following command:

```console
docker build -t graphql-go-todo .
```

Once the image is created, it can be started by indicating the environment variable needed for the connection to the postgres as shown in the following example.

```console
docker run --env PSQL_CONN_STR="user=USER dbname=DBNAME password=PASSWORD host=HOST port=PORT  sslmode=disable" -d -p 3000:3000 graphql-go-todo
```
After that the server should be started on port 3000.



