build:
	go build -o ./bin/blockchain-go

run: build
	./bin/blockchain-go

test:
	go test -v ./...
