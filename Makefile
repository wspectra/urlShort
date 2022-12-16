all:
	docker-compose   up -d


build:
	go build -o urlShort cmd/urlShort/main.go

inmemory:
	go run cmd/urlShort/main.go -store-flag "inmemory"


clean:
	- docker-compose down
	- docker volume rm $$(docker volume ls -q)