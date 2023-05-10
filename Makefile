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
