# docker-compose
## Build
1. `docker-compose build`

## Run
1. `docker-compose run frontend yarn dev`

## Create Docker network
1. `docker network create pgsql-network`

# Heroku deploy
## Push image to Heroku container repository
1. `heroku container:login`
2. `heroku container:push web`

## Release image
1. `heroku container:release web`

# Heroku config
## Push .env
1. `heroku config:push`

## Pull .env
1. `heroku config:pull`
