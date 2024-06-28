FROM golang:1.22.3-alpine3.19 AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod tidy

COPY . .

RUN go build -o main ./cmd/main.go

FROM alpine:3.19

WORKDIR /

COPY --from=build app/main /main
COPY --from=build app/migration /migration

CMD ["./main"]
