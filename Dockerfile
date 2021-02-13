# frontend container
FROM node:alpine as frontend

WORKDIR /app

RUN apk update \
  && apk add --no-cache python3 make g++

COPY ./frontend .
RUN yarn install \
  && yarn generate


# backend container
FROM golang:alpine as backend

WORKDIR /app

RUN apk update \
  && apk add --no-cache git curl \
  && go get -u github.com/cosmtrek/air

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY ./backend .
RUN go build main.go


# runner container
FROM alpine:latest as runner
COPY --from=backend /app /app

WORKDIR /app
CMD ./main
