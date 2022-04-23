# syntax=docker/dockerfile:1

FROM golang:1.18 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /main

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /main

EXPOSE 8080

ENTRYPOINT ["/main"]