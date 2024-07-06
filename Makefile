BIN_PATH=bin/kinolove/main.out

all: test build

build: 
	@go build -o ${BIN_PATH} cmd/kinolove/main.go

test:
	@go test -v ./...
	
run: build
	${BIN_PATH}

clean:
	@go clean 
	@if [ -e ${BIN_PATH} ]; then \
	    rm ${BIN_PATH}; \
	fi
