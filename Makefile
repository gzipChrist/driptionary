BINARY_NAME=drip

build:
	go build -o ${BINARY_NAME} cmd/driptionary/main.go

run:
	go run cmd/driptionary/main.go

clean:
	go clean
	rm ${BINARY_NAME}

tidy:
	go mod tidy
