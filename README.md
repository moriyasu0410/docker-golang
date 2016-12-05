# docker-golang

* Golang Sample Application on Docker

## Requirement
* Docker >= 1.12

## Usage
* Build
```
$ cd /path/to/docker-golang
$ docker-compose build
```

* Up
```
$ docker-compose up -d
```

* Run tests
```
$ docker-compose exec golang go test -v ./...
```