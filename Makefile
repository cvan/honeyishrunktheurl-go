all: build

build: clean
	@go env
	@go get .
	@go build -o server.o

clean:
	@go clean
	@rm -f server.o
	@go fmt

run: build
	./server.o

threaded: build
	GOMAXPROCS=8 ./server.o
