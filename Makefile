build:
	go build -o calculator-api .

test:
	go test ./...

fmt:
	go fmt ./...

run:
	go run .