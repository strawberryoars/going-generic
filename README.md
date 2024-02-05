# going-generic

This repository will manage multiple microservices for generic json schema / resource architecure.
Infrastructure will run on kubernetes [K3S](https://k3s.io/).

Learn about [JSON Schemas](https://json-schema.org/)

## Project

Architecture was split into multiple miroservices for fun.
I can explore and tinker with each different component of the arhcitecture with out it impacting the rest of the project.

### schema-api
JSON schema web server
    - built off [ElysiaJS](https://elysiajs.com/) and [Bun](https://bun.sh/)
    - serves JSON schemas to clients
    - generator support (fake data that conforms to your schema by utilzing the [JSON Schema Faker](https://github.com/json-schema-faker/json-schema-faker) package)

### resources-api
Golang webserver for CRUD support on generic resources which are backed by a JSON schema on the schema-api
    - serves generic resources from mongo
    - CRUD support
    - schema validation on CRUD operations

### resouces-ui
Simple generic resources UI built w/ [HTMX](https://htmx.org/)
    - generic resources table w/ infinite scroll
