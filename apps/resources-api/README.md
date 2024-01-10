# resources-api

Simple web server:
- serves generic resources from DB
- CRUD support w/ json patch history
- endpoint to list json patch history for a given resource
- argo events for notification of changes to generic resources


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