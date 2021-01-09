FROM golang:alpine as builder

RUN apk update \
  && apk add --no-cache git curl

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /app
COPY ./backend .
RUN go build main.go

FROM alpine:latest as runner
COPY --from=builder /app /app

WORKDIR /app
CMD ./main
