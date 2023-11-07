.PHONY: sqlc migrate tidy check test up down

sqlc:
	sqlc generate --experimental --file ./sql/sqlc.yaml

migrate:
	migrate create -ext sql -dir ./internal/repository/migrations -seq create_events_table

tidy:
	go fmt ./...
	go mod tidy -v

check:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

test:
	go test -v -race -buildvcs ./...

up:
	docker-compose up --build

down:
	docker-compose down

ping:
	curl http://localhost:3333/ping

create-event:
	curl -X POST \
		-H "Content-Type: application/json" \
		-d '{"name": "Random Event", "description": "a random test event", "category": "random", "location": "Buenos Aires", "publisher": "Nico", "lineup": ["Random Artits"], "start_ts": "2023-01-01T18:00:00Z","end_ts": "2023-01-01T21:00:00Z"}' \
		http://localhost:3333/bastet

get-event:
	curl http://localhost:3333/bastet/1

update-event:
	curl -X PUT \
		-H "Content-Type: application/json" \
		-d '{"description": "Changed Description"}' \
		http://localhost:3333/bastet/1

delete-event:
	curl -X DELETE http://localhost:3333/bastet/1

all-requests:
	make create-event && make get-event && make update-event && make delete-event
