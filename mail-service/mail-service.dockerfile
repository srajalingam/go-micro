#base go image

FROM golang:1.25-alpine AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o mailerApp ./cmd/api

RUN chmod +x /app/mailerApp

#build a tiny docker image

FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app/mailerApp /app
COPY ./templates /app/templates

CMD ["/app/mailerApp"]

