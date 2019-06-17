PROJECT_NAME=boilerplate-go
PROJECT_ORG=devit-tel
VERSION=latest
GOCMD=go
EXPOSED_PORT=80

REPO?=$(PROJECT_ORG)/$(PROJECT_NAME)
BINARY_NAME=$(PROJECT_NAME)
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: run

build:
	$(GOCMD) build -o ./dist/$(BINARY_NAME) -tags netgo -a -v
test:
	$(GOCMD) test -v ./...
env:
	godotenv
clean:
	$(go clean)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	make build
	./$(BINARY_NAME)
run-watch:
	godotenv fresh
install:
	$(GOCMD) get github.com/joho/godotenv/cmd/godotenv
	$(GOCMD) get github.com/gravityblast/fresh
build-docker:
	docker build -t $(REPO):$(VERSION) .
push-docker:
	docker push $(REPO):$(VERSION)
run-docker:
	docker run -p $(EXPOSED_PORT):80 $(REPO):$(VERSION)
test-docker:
	docker build --no-cache --rm --target=tester .
