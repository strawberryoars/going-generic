# resources-api - simple generic resources web server:

Generic resources as they are defined here are json objects.
These resources will have their own MongoDB collection for storage.
Each resource should have a defined JSON schema to enforce validation on creation and updates.

- serves generic resources from MongoDB
    - supports CRUD
    - support for filter, sort, and pagination


## TODO
- json patch history
    - endpoint to list json patch history for a given resource
- argo resource events
    - generate notifications if a resource is failed to be created, updated, or deleted. Dev channel.
    - generate event when created, updated, or deleted. Could be a cloudevent that is consumed to replicate data to other clusters or for consumers who want to be aware of changes
- evaluate archive support
    - could have an archival collection for each resource collection so users can retain data without deleting permanently


## devlopment

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