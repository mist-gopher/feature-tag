build:
	go build -o ./bin/app ./cmd/main.go

test:
	go test ./...

run:
	go run ./cmd/main.go

start: build
	./bin/app