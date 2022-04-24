# syntax=docker/dockerfile:1

FROM golang:1.18 as build

WORKDIR /app

COPY go.mod ./src/
COPY go.sum ./src/

RUN go mod download

COPY *.go ./src/

RUN go build -o /main

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]