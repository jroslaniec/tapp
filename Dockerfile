FROM golang:alpine AS base

WORKDIR /app

COPY . .

RUN go mod download && go build -o tapp main.go

FROM alpine

RUN apk update \
    && apk add ca-certificates \
    && update-ca-certificates

COPY --from=base /app/tapp tapp

CMD ["./tapp"]
