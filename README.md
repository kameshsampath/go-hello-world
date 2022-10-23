# Hello World

A simple REST API built in `golang` using Labstack's [Echo](https://https://echo.labstack.com/]), to demonstrate how to build multi architecture container images using [Drone CI](https://drone.io) and [Buildx]([https://](https://docs.docker.com/build/buildx/install/)

## Pre-requisites

- [Docker Desktop](https://docs.docker.com/desktop/)
- [Drone CI CLI](https://docs.drone.io/cli/install/)

## Environment Setup

The `.env` file helps configure the following settings,

- `PLUGIN_REGISTRY` - the docker registry to use
- `PLUGIN_TAG`      - the tag to push the image to docker registry
- `PLUGIN_REPO`     - the docker registry repository

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
