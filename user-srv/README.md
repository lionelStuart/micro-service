# User Service

This is the User service

Generated with

```
micro new user-srv --namespace=mu.micro.book --alias=user --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.book.srv.user
- Type: srv
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
./user-srv
```

Build a docker image
```
make docker
```

Test run 
```
micro --registry=consul call mu.micro.book.srv.user User.QueryUserByName '{"userName":"micro"}'
{
        "success": true,
        "user": {
                "id": 10001,
                "name": "micro",
                "pwd": "1234"
        }
}

```