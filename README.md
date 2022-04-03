# Book and Author information GraphQL API
[![Go Reference](https://pkg.go.dev/badge/golang.org/x/example.svg)](https://pkg.go.dev/golang.org/x/example)
## Problem 

How would you construct an event ( message bus ) driven micro-services for the given scenario where a user requested to see his/her order, profile, and invoices but the data is distributed separately among Order, Profile, and Invoice Micro Services each containing its own Database. Kindly, please illustrate with a diagram.

## Setup
```shell
     $ go test ./...  # run from root of the project directory 
     $ docker-compose down --volume # to make sure to remove shared volume
     $ docker-compose up --build --force-recreate # here --force-recreate is optional
```
- ![#f03c15](https://via.placeholder.com/15/f03c15/000000?text=+) `Note`If you interested to run it from locally without Docker please ensure database and .env exist
>Example .env file
>>GIN_PORT=8080 \
GIN_MODE=debug \
DB_HOST=localhost \
DB_PORT=5432 \
DB_USER=postgres \
DB_PASSWORD=mithu1996 \
DB_NAME=test