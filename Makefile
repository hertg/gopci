all: test lint

test:
	go test -v ./...

lint:
	go vet ./...

benchmark:
	go test -v ./... -bench . -benchmem
