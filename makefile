restart:
	docker-compose down
	docker-compose up --build -d

dev:
	cd ./go
	go run main.go