FROM golang:1.22.3-bullseye

WORKDIR /api

RUN go install github.com/air-verse/air@latest

CMD [ "air", "-c", ".air.toml"]