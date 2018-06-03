RM=rm
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=lms

DEPLOY_NAME=deploy
DEPLOY_PATH=cmd/deploy/deploy.go

PORT=8080
CLEARDB_DATABASE_URL="mysql://root:@/lms"

default: test-oneline build deploy

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

test-oneline:
	$(GOTEST) ./...

test:
	$(GOTEST) -v ./...

clean: deploy-clean
	$(GOCLEAN)
	$(RM) -f $(BINARY_NAME)

run:
	$(GOBUILD) -o $(BINARY_NAME) -v
	PORT=$(PORT) CLEARDB_DATABASE_URL=$(CLEARDB_DATABASE_URL) ./$(BINARY_NAME)

deploy:
	$(GOBUILD) -o $(DEPLOY_NAME) -v $(DEPLOY_PATH)

deploy-run: deploy
	./$(DEPLOY_NAME)

deploy-clean:
	$(RM) -f $(DEPLOY_NAME)
