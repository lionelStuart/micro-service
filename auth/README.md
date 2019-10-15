# Auth Service

This is the Auth service

Generated with

```
micro new auth --namespace=mu.micro.book --alias=auth --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.auth
- Type: srv
- Alias: auth

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./auth-srv
```

Build a docker image
```
make docker
```

Test User-Srv

```
micro --registry=consul --api_namespace=mu.micro.book.web api --handler=web

cd user-srv
go run main.go plugin.go

cd user-web
go run main.go

cd auth
go run main.go

# login
curl --request POST --url http://127.0.0.1:8080/user/login --header "Content-Type: application/x-www-form-urlencoded" --data "userName=micro&pwd=1234"
# logout
curl --request POST --url http://127.0.0.1:8080/user/logout --cookie "remember-me-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzM3NzUyMjQsImp0aSI6IjEwMDAxIiwiaWF0IjoxNTcxMTgzMjI0LCJpc3MiOiJib29rLm1pY3JvLm11IiwibmJmIjoxNTcxMTgzMjI0LCJzdWIiOiIxMDAwMSJ9.rKUDptpHqKwJLqQafgrIXk1AuXz_Dp7aSi3L2ycM1f"

```