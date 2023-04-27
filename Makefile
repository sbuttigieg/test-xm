.PHONY: lint
lint:
	@golangci-lint run

.PHONY: test
test:
	@go test -race -cover ./xm_app/...

.PHONY: compile
compile:
	@GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/test-xm ./cmd/main.go

.PHONY: compose-build
compose-build:
	@docker-compose up --build

.PHONY: compose-start
compose-start:
	@docker-compose up

.PHONY: compose-start-d
compose-start-d:
	@docker-compose up -d

.PHONY: compose-stop
compose-stop:
	@docker-compose stop

.PHONY: compose-remove
compose-remove:
	@docker-compose down -v

.PHONY: image-remove
image-remove:
	@docker image rm app-test-xm