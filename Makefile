GOTEST=go test
GOBUILD=go build
GOCLEAN=go clean
SOURCE_FILE=cmd/main.go
BINARY_FILE=build/service

all: test build

test:
	$(GOTEST) -v ./...

build:
	$(GOBUILD) -o $(BINARY_FILE) $(SOURCE_FILE)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_FILE)

run:
	./$(BINARY_FILE)

docker: docker-build docker-run	

docker-build:
	docker build -t rest-api-service .

docker-run:
	docker run -it --rm -p 8000:8000 --name api-server rest-api-service

.PHONY: test build
