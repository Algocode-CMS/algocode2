BIN_NAME := algocode

build:
	go build -o ${BIN_NAME} cmd/main.go

test:
	echo "OK"

lint:
	golangci-lint run
