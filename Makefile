API_DOCS_PATH = /docs
all: init test build

.PHONY: init
init:
	@echo "> Installing the server dependencies ..."
	@go mod tidy -v
	@go get -v ./...
	@go install github.com/swaggo/swag/cmd/swag

.PHONY: test
test:
	@echo "> Testing the server source code ..."
	@go test -cover -covermode atomic -coverprofile cover.out -race ./...
	@go tool cover -func cover.out

.PHONY: build
build: gen-swagger
	@echo "> Building the server binary ..."
	@go build -o bin/quiz-master .

.PHONY: gen-swagger
gen-swagger:
	@echo "Updating API documentation..."
	@swag init -o ${API_DOCS_PATH}

.PHONY: run
run: build
	@echo "> RUN the server binary ..."
	@./bin/quiz-master

.PHONY: migrate
migrate:
	@echo "create migrate file to /migrations"
	@migrate create -ext sql -dir migrations -seq create_table_

.PHONY: migrate-up
migrate-up:
	@echo "migrate up"
	@migrate -database 'mysql://root:@tcp(localhost:3306)/evoting' -path migrations up