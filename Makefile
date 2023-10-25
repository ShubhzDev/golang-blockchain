build:
	go build -o ./bin/golang-blockchain

run: build
	./bin/golang-blockchain

test:
	go test -v ./...