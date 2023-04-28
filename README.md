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
- An `.env` file and a `env.yml` are included in the repo for this test. However in a working environment these files would not be included. The `.env` file is here beinh used to populate the variables in the local docker-compose file. The `env.yml` is an example of what would be in another repository from where Kubernetes would take the env variables to build the containers. In this file the sensitive variable values are not shown and are marked secrets.
- For the secure endpoints a jwt token must be passed as a bearer token in the request header.
- A token can be retrieved by sending a `/companies/token` request with correct email and password in the body.
- New users can be added by sending a `/companies/token` request. However a user has been added to migrations to ease testing. The user credentials are: 
```
{
    "email": "thedoe@gmail.com",
    "password": "ilovegolang"
}
```
