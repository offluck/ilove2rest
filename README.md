# I really love 2 REST

## General

This project is a simple CRUD app written in Go, uses PostgreSQL as primary database and Redis as cache (cache coming soon for now, only carcass is ready)

I tried implementing some of [12-factor app](https://12factor.net/) and HATEOAS ideas, hope you'll like it;)

## Project Structure

```
ilove2rest
├── Dockerfile
├── Makefile
├── README.md
├── cmd
│   └── app: Application entrypoint (main)
├── config
│   ├── dev.yaml: BackEnd launches locally, DataBase launches inside a container
│   └── prod.yaml: Launches in docker-compose
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── config
│   ├── repository: DataBase and Cache interaction layer
│   ├── entities
│   └── server: HTTP interaction layer
└── migrations
```

## Endpoints

You can find a OpenAPI/Swagger specification in `api` directory (Not present yet)

While it is not present, let me point out URLs here:
- `/health`:
  - GET: Responses with "I am healthy as an ox!" if service is healthy and ready to work
- `/api/v0/user`:
  - GET: Responses with list of all users' info
  - POST: Adds user's info and responses with it
- `/api/v0/user/{username}`:
  - GET: Responses with chosen user's info
  - PUT: Updates chosen user's info and responses with it
  - DELETE: Deletes chosen user's info

## Starting the App

The easiest way to start a production version is to run:
```bash
make prod
```
It will build backend and database containers and start them. The backend will wait till database container is healthy

## TODOs

List of ideas:
- Implement cache
- Add useful info to headers
- Enhance error handling and representation
- Enhance logging
