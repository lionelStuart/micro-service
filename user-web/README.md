# User Service

This is the User service

Generated with

```
micro new user-web --namespace=mu.micro.book --alias=user --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.web.user
- Type: web
- Alias: user

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
./user-web
```

Build a docker image
```
make docker
```

Test user-web

```
# run api
micro --registry=consul --api_namespace=mu.micro.book.web api --handler=web

# run user-srv
cd user-srv
go run main.go plugin.go

# run uer-web
go run main.go

# request test
curl --request POST --url http://127.0.0.1:8080/user/login --header 'Content-Type: application/x-www-form-urlencoded' --data 'userName=micro&pwd=1234'

{"data":{"id":10001,"name":"micro"},"ref":1570839637332482044,"success":true}

```
