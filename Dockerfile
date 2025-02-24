# Arguments
ARG GOLANG_URI
ARG ALPINE_URI

FROM ${GOLANG_URI:-golang}:1.22.1-alpine as go-builder
RUN apk add build-base

WORKDIR /app

# Copy go mod and sum files.
COPY go.mod go.sum ./

# Download all dependencies.
RUN go mod download

# Copy the service source code.
COPY ./railway-signal-service ./railway-signal-service

RUN GOOS=linux GOARCH=amd64 go build -o main ./railway-signal-service/cmd/api/main.go

FROM ${ALPINE_URI:-alpine}:latest

WORKDIR /app

COPY --from=go-builder /app/main .
ENTRYPOINT ["/app/main"]