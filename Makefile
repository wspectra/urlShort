all:
	#docker-compose build #--no-cache
	#docker-compose up -d --force-recreate
	docker-compose   up -d


build:
	go build -o urlShort cmd/urlShort/main.go


clean:
	- docker-compose down
	- docker volume rm $$(docker volume ls -q)