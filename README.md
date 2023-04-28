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
- On Github, the `master` branch contains the latest code version. It has been pulled from `dev` just before submitting the test.
- An `.env` file and a `env.yml` are included in the repo for this test. However in a working environment these files would not be included. The `.env` file is here beinh used to populate the variables in the local docker-compose file. The `env.yml` is an example of what would be in another repository from where Kubernetes would take the env variables to build the containers. In this file the sensitive variable values are not shown and are marked secrets.
- For the secure endpoints a jwt token must be passed as a bearer token in the request header.
- A token can be retrieved by sending a `/companies/token` request with correct email and password in the body.
- New users can be added by sending a `/companies/users` request. However a user has been added to migrations to ease testing. The user credentials are: 
```
{
    "email": "thedoe@gmail.com",
    "password": "ilovegolang"
}
```

- Passwords are hashed before being saved to database.
- Application runs on localhost:7711, example: localhost:7711/companies/users
- A postman collection was added to the repo to make testing easier.
- Postgres was used as the database, which contains 3 tables: `companies`, `users` and `migrations`. It can be accessed from outside of the docker environment using localhost:5700. The database name, username and password can be found in the environment files as explained above.
- Though unit tests were not in the requirements, I believe that they are essential for good code maintainability and given more time I would have include them. I usually use mocks for service and store functions and aim for >90% coverage. An example unit test file has been included. `create_test.go`
- Linter has been used to keep the code standardised and clean.
- Due to time constraints kafka events and integration test were not done. I focused instead on making  a good code structure.
