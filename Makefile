all:
	@go build -o bin/main main.go 

install: # install dependencies
	@go mod tidy

run:
	@./bin/main