# Hello World

A simple REST API built in `golang` using Labstack's [Echo](https://https://echo.labstack.com/]), to demonstrate how to build multi architecture container images using [Drone CI](https://drone.io) and [Buildx](https://docs.docker.com/build/buildx/install/)

## Pre-requisites

- [Docker Desktop](https://docs.docker.com/desktop/)
- [Drone CI CLI](https://docs.drone.io/cli/install/)

## Download Sources

Clone the sources and CD into it,

```shell
git clone https://github.com/kameshsampath/go-hello-world.git && cd "$(basename "$_" .git)"
```

## Environment Setup

The demo will use a free container registry service [ttl.sh](https://ttl.sh/)for pushing the application image.

```shell
# a unique uid as image identifier, it needs to be in the lowercase
export IMAGE_NAME=$(uuidgen | tr '[:upper:]' '[:lower:]')
# short lived image for 10 mins
export IMAGE_TAG=10m
```

Let's setup `.env` that will be used by the build and docker-compose,

```shell
envsubst < .env.example | tee .env
```

>**NOTE:** the example above uses [envsubst](https://www.man7.org/linux/man-pages/man1/envsubst.1.html) to update the file, if you dont have envsubst installed you can manually update the .env file

## Build the Application

```shell
drone exec --trusted --env-file=.env
```

The command will test, build and push the container image to the `$PLUGIN_REPO:$PLUGIN_TAG`.

## Run Application

```shell
docker-compose up
```

## Testing

```shell
curl http://localhost:8080/
```

The command should return `Hello World!`.
