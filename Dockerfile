FROM golang:1.21-alpine AS base

WORKDIR /app

FROM base as builder
COPY go.mod go.sum ./
RUN go mod download

COPY main.go  .
RUN go build -o ./build

FROM base as release
COPY --from=builder /app/build ./build

CMD ./build