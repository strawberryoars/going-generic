# going-generic


## Project Plan

This repository will manage multiple microservices for generic json schema / resource architecure
- json schema server API
    - serves json schemas to clients
    - generator support (fake data that conforms to the schema)
- generic golang resources API
    - serves generic resources from DB
    - CRUD support w/ json patch history
    - endpoint to list json patch history for a given resource
    - argo events for notification of changes to generic resources
- storage layer(s)
    - mongoDB, minio, etc.
- UI to CRUD generic resources
    - pages to view json patch history
- UI to run visualizations on common probability and statistics
    - this will leverage the generic resources UI

Infrastructure will run on kubernetes (k3s)