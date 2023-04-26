# test-xm

## Steps to install application
- using script to run make files
    - `./dockerise-up.sh`
- step by step
    - `GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o bin/test-xm ./cmd/main.go`
    - `docker-compose up --build`

## Steps to stop application
- using make files
    - `make compose-stop`
- using docker-compose
    - `docker-compose stop`

## Steps to remove application (including persisted data)
- using script to run make files
    - `./dockerise-rm.sh`
- step by step
    - `docker-compose down -v`
    - `docker image rm app-test-xm`
