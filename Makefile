BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

test:
	go test -v main.go

clean:
	go clean
	rm ${BINARY_NAME}