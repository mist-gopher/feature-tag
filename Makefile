build:
	go build -o ./bin/app ./cmd/main.go

test:
	go test -v ./...

run:
	go run ./cmd/main.go

start: build
	./bin/app