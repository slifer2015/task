# Project
HTTP server for a service that would make http requests to 3rd-party services.

## Work algorithm :

The client sends a task to the service to perform an http request to a 3rd-party services.
The task is described in json format, the generated task id is returned
in response and its execution starts in the background.

### Enhancement
I added the ability to configure the number of workers along with other dynamic parts

## Quick Start
- for running project you can simply run
```shell
docker compose up -d
```
- or if you have Go installed you can run with below command
```shell
make start
```
- for generating swagger files
```shell
make generate
```
- for running tests
```shell
make test
```

- for running linter
```shell
make lint
```

after running you can simply check swagger path
http://127.0.0.1:[port]/swagger/index.html

enjoy!