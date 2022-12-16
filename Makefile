all:
	docker-compose   up -d


build:
	go build -o urlShort cmd/urlShort/main.go

inmemory:
	go run cmd/urlShort/main.go -store-flag "inmemory"

test:
	-go test ./internal/server/... -v

status:
	docker ps -a

clean:
	- docker-compose down

.PHONY: all build inmemory test clean status