# Todos CLI App

## Description

The goal of this project is to create a web api which perform math calculation with differents versions in [Go](https://go.dev) to learn the basics of [Go](https://go.dev). If you have any suggestion to add more math endpoints create an issue or a pull request if you want to.

## Tech Stack

- [Cobra](https://cobra.dev)
- [mux](https://github.com/gorilla/mux)
- [prometheus](https://github.com/prometheus/client_golang)
- [UUID](https://github.com/google/uuid)

## Tasks

- [x] Install all dependencies.
- [x] Create a cli
  - [x] Flag which give the ability to choose the log format (`JSON` or `text`)
  - [x] Flag which give the ability to choose the port
  - [x] Flag which give the ability to choose the graceful timeout
  - [x] Port environment variable
- [x] Create a middleware stacker
- [x] Create middlewares
  - [x] Rate limiting
  - [x] Logging
  - [x] Enforce the `Content-Type` header to be `application/json`
- [x] Create the `/v1` endpoints
  - [x] Create a `/add` endpoint (1+1)
  - [x] Create a `/divide` endpoint (1/1)
  - [x] Create a `/multiply` endpoint (1*1)
  - [x] Create a `/substract` endpoint (1-1)
- [x] Create a new middleware which check if the format is a JSON or not
- [x] Create the `/v2` endpoints
  - [x] Update the `/add`, `/divide`, `/multiply`, and `/substract` endpoints to have the new middleware
- [x] Add `/health` endpoint
- [x] Add `/ping` endpoint
- [x] Add prometheus
- [ ] Create tests
- [ ] Github actions
