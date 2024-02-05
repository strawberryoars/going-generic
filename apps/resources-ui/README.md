
# resources-ui

Creating a generic resource table using HTMX.
Infinite scroll of resources to lazily load resources as they come into the users viewport.

TODO: column formatting
TODO: cleanup server-side html templating

## devlopment

Install [GVM](https://github.com/moovweb/gvm)

module:
```
go mod init github.com/strawberryoars/going-generic/apps/resources-api
```

Run web server:
```
go run main.go
```