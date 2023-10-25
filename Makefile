build: 
	@go build -o bin/golang-apis

run: build
    @./bin/golang-apis

test:
	@go test -v ./...

migrate:
	go run ./scripts/migrate.go