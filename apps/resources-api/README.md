# resources-api

Simple generic resources web server:
- serves generic resources from MongoDB
    - support for filter, sort, and pagination
    - resources should have a defined schema in schemas-api app
- CRUD support w/ json patch history
    - resources correspond to a collection in MongoDB. API will support CRUD.
    - endpoint to list json patch history for a given resource
    - argo events for notification of changes to generic resources
- archive support w/ arhival collections for resources


# devlopment

Install [GVM](https://github.com/moovweb/gvm)

module:
```
go mod init github.com/strawberryoars/going-generic/apps/resources-api
```


Define mongodb uri in .env file


Run web server:
```
go run main.go
```


Client Request:
```
curl http://localhost:8080/query?collection=blogs
```

## MongoDB Driver

https://raw.githubusercontent.com/mongodb/docs-golang/master/source/includes/usage-examples/code-snippets/command.go