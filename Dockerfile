FROM alpine

RUN apk update \
    && apk add ca-certificates \
    && update-ca-certificates \
    && wget https://github.com/jroslaniec/tapp/releases/download/v0.1.0/tapp \
    && chmod +x tapp

CMD ["./tapp"]
