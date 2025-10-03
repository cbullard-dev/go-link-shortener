BINARY_NAME=main.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

test:
	go test -v ./...

clean:
	go clean
	rm ${BINARY_NAME}

bootstrap:
	git fetch && git pull
	cp bootstrap/git-pre-commit.hook .git/hooks/pre-commit
	chmod a+x .git/hooks/pre-commit