# I really love 2 REST

## General

This project is a simple CRUD app written in Go, uses PostgreSQL as primary database and Redis as cache (cache coming soon for now, only carcass is ready)

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
