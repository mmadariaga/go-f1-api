# Go REST API - Learning Project

This project is a personal learning exercise to build a simple REST API using Go and Next.js. The goal is to practice building APIs with the following tools and libraries:

- **[go-chi/chi](https://github.com/go-chi/chi/v5)**: A lightweight, idiomatic HTTP router for Go.
- **[rs/zerolog](https://github.com/rs/zerolog)**: A high-performance, structured logging library.
- **[stretchr/testify](https://github.com/stretchr/testify)**: A testing toolkit for Go with a variety of helpers for assertions and mocking.


## Install & Run
```bash
git clone https://github.com/mmadariaga/go-api.git
cd go-api
docker-compose up
```

## Features

Currently, the API only has two routes:
- `GET /ping`: Returns a simple "pong" response.
- `GET /races/{year}`: Returns application_races.RacesByYearResponse json object with a summary of all the races of the year 2024 in f1. This endpoint requires Basic Auth (go:api)
- `GET /races/drivers`: Returns domain_model.Driver json object array with the drivers that participated in the last race. This endpoint requires Basic Auth (go:api)

Example:
```bash
$ curl http://localhost:8080/ping
pong
```

## WIP
 - Create a web interface with next.js and tailwind

## Licence
This project is licensed under the MIT License - see the LICENSE file for details.