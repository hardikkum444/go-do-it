build:
	@go build -o bin/go-do-it

run: build
	@./bin/go-do-it

test:
	@go test -b ./...
