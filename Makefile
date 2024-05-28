build:
	go build -o ./bin/ticketBot main.go

run: build
	./bin/ticketBot