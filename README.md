# Bastet
Go service developed following clean architecture. Centered around the concept of events, provides a REST layer to perform CRUD operations on them.

For a deeper look into the service architecture, including app structure, design decisions, library choices and more take look at [ARCHITECTURE.md](ARCHITECTURE.md).

## How to run
Docker and docker-compose are needed to start up the server.

`make up` will build the containers from scratch, starting the DB and, if successful, the server on port 3333.

A `.env` file is needed, provided upon request. That said, for testing purposes, something basic like this would work:

```
SERVER_HOST=localhost
SERVER_PORT=3333
DB_NAME=bastet
DB_USERNAME=postgres
DB_PASSWORD=postgres
DB_HOST=postgres
DB_PORT=5432
```

## Test the server
Some basic curl commands are provided for convenience to actually test the server is operating right.

1. Ping: `make ping`
2. Create an event: `make create-event`
3. Get an event: `make get-event`
4. Update an event: `make update-event`
5. Delete an event: `make delete-event`
6. Run CRUD: `make all-requests`


## TBD
1. Deployment github workflow.
