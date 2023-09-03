build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r '\.go$$' make run

d.up:
	docker-compose up 

d.down:
	docker-compose down

d.up.build:
	docker-compose --build up
