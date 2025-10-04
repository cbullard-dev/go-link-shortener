BINARY_NAME=main.out
RELATIVE_PATH=./main.out

build:
	go build -o ${BINARY_NAME} main.go

run:
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

check:
ifeq ("$(wildcard $(RELATIVE_PATH))","${RELATIVE_PATH}")
	./${BINARY_NAME} https:\/\/example.com\/test
endif

test:
	go test -v ./...

clean:
	go clean
ifneq ("$(wildcard $(RELATIVE_PATH))","")
	rm ${BINARY_NAME}
endif

bootstrap:
	git fetch && git pull
	cp bootstrap/git-pre-commit.hook .git/hooks/pre-commit
	chmod a+x .git/hooks/pre-commit