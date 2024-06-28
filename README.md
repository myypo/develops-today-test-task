## Description

Simple **Go** CRUD web server. Uses **Gin** and **Jet** for communication with **Postgres** database.

## Requirements

Assumes the following tools are installed and in PATH:

1. Go 1.22.3
2. Docker and docker compose
3. [Goose](https://github.com/pressly/goose)
4. Optional: [Just](https://github.com/casey/just) for running the setup script. Alternatively you can simply run the commands specified in the `justfile`

## Setup

Run `justfile local` or execute the bash script specified inside

The server will run, by default on port: `12499`, Postgres on its default port: `5432`

The relevant **Postman** collection is stored in `postman_collection.json`
