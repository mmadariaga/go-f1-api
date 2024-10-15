# Go REST API - Learning Project

This project is a personal learning exercise to build a simple REST API using Go and Next.js. The goal is to practice building APIs with the following tools and libraries:

- **[go-chi/chi](https://github.com/go-chi/chi/v5)**: A lightweight, idiomatic HTTP router for Go.
- **[rs/zerolog](https://github.com/rs/zerolog)**: A high-performance, structured logging library.
- **[stretchr/testify](https://github.com/stretchr/testify)**: A testing toolkit for Go with a variety of helpers for assertions and mocking.


## Install & Run
```bash
git clone https://github.com/mmadariaga/go-api.git
cd go-api/api
# install dependencies
go mod tidy 
# Run tests
bin/test
# Run the server
bin/run
```

## Features

Currently, the API only has two routes:
- `GET /ping`: Returns a simple "pong" response.
- `GET /races/{year}`: Returns application_races.Response json object with a summary of all the races of the year 2024 in f1. This endpoint requires Basic Auth (go:api)

Example:
```bash
$ curl http://localhost:8080/ping
pong
```

## TODO
 - [ ] Accept year as an argumento of /races api endpoint (2024 currently)
 - [ ] Create an interface in next.js with tailwind

## Licence
This project is licensed under the MIT License - see the LICENSE file for details.