# test-xm

## Steps to install application
- using script to run make files
    - `./dockerise-up.sh` (use `chmod +x dockerise-up.sh` in case of permission issues)
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
    - `./dockerise-rm.sh` (use `chmod +x dockerise-rm.sh` in case of permission issues)
- step by step
    - `docker-compose down -v`
    - `docker image rm app-test-xm`

## Notes:
- An `.env` file and a `env.yml` are included in the repo for this test. However in a working environment these files would not be included. The `.env` file is here beinh used to populate the variables in the local docker-compose file. The `env.yml` is an example of what would be in another repository from where Kubernetes would take the env variables to build the containers.