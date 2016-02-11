all: test

test:
	go test ./models

r: run

run:
	go run main.go
