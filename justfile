default: 
    just --list --unsorted

@local:
    #!/usr/bin/env bash

    docker compose --file local.docker-compose.yaml up -d --build --force-recreate
    until docker exec sca-database pg_isready -p 5432 ; do sleep 0.25 ; done

    pushd ./migration
    export GOOSE_DRIVER=postgres
    export GOOSE_DBSTRING="postgresql://postgres:postgres@localhost:5432/sca?sslmode=disable"
    sleep 1
    goose up
