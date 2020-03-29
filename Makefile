BINARY_NAME=bin/vscale
BINARY_NAME_LINUX=$(BINARY_NAME)-linux
BINARY_NAME_WINDOWS=$(BINARY_NAME)-windows.exe
BINARY_NAME_DARWIN=$(BINARY_NAME)-darwin
DOCKER_IMAGE=hawk.collector

export GO111MODULE=on

all: bin

build:
	go build -o $(BINARY_NAME) -v ./
	chmod +x $(BINARY_NAME)
clean:
	go clean
	rm -rf $(BINARY_NAME)
	rm -rf $(BINARY_NAME_LINUX)
	rm -rf $(BINARY_NAME_WINDOWS)
	rm -rf $(BINARY_NAME_DARWIN)
run: bin
	cp config.yml ./bin/config.yml
	./bin/monitoring --config ./bin/config.yml

build-all: build-linux build-windows build-darwin

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME_LINUX) -v $(SRC_DIRECTORY)

build-windows:
	GOOS=windows GOARCH=amd64 go build -o $(BINARY_NAME_WINDOWS) -v $(SRC_DIRECTORY)

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME_DARWIN) -v $(SRC_DIRECTORY)
