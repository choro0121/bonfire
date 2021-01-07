# Docker command
## Build backend image
1. `docker build -t backend .`

## Run backend binary
1. `docker run -e "PORT=3000" -p 3000:3000 -t backend`

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
