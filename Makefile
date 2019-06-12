PROJECT_NAME=boilerplate-go
PROJECT_ORG=NV4RE
VERSION=latest
GOCMD=go
EXPOSED_PORT= "8080"

REPO?=$(PROJECT_ORG)/$(PROJECT_NAME)
BINARY_NAME=$(PROJECT_NAME)
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: run

build:
	$(GOCMD) build -o ./dist/$(BINARY_NAME) -v
test:
	$(GOCMD) test -v ./...
clean:
	$(go clean)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	make build
	./$(BINARY_NAME)
run-watch:
	fresh
install:
	$(GOCMD) mod init 
	$(GOCMD) get github.com/gravityblast/fresh
build-docker:
	docker build --no-cache -t $(REPO):$(VERSION) .
push-docker:
	docker push $(REPO):$(VERSION)
run-docker:
	docker run --rm -p $(EXPOSED_PORT):8080 $(REPO):$(VERSION) --log-level=debug --port=8080
test-docker:
	docker build --no-cache --rm --target=tester .
