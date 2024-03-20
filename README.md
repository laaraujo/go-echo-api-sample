# Golang Echo API

Sample Rest API project I made while learning some Go and [Echo](https://echo.labstack.com/)

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [Production Docker image](#production-docker-image)

## Overview

This project intends to serve as a template to kickstart Go API development with Echo framework.

## Features

- Database schema and migrations with [Goose](https://pressly.github.io/goose/)
- Type-safe SQL interfaces with [sqlc](https://sqlc.dev/)
- Local development setup with [Docker](https://www.docker.com/), [Docker Compose](https://docs.docker.com/compose/), [Make](https://www.gnu.org/software/make/manual/make.html)
- Live reload with [Air](https://github.com/cosmtrek/air) in [Docker](https://www.docker.com/)
- Very basic placeholder auth middleware with API_KEY

## Installation

Setup your local environment with the following commands
```sh
# Install goose and sqlc locally
$ make setup
# Build your containers
$ make build
```

## Usage

### Run locally
```sh
# Start your containers
$ make run
```

Then in another shell:
```sh
# Update your database schema by applying the migrations
$ make goose/migrate
```

### Create a new migration file
```sh
$ make goose/create
```
* Note: this will create a file with the default name `*_rename_this_file.sql` so you may want to rename it.

### Generate SQLC types
```sh
$ make sqlc/generate
```

### Stop local containers
```sh
$ make stop
```

## Production Docker image

If you want to test your Docker image:

```sh
# Build your image
$ docker build -f docker/Dockerfile -t golang-echo-api .
```

```sh
# Replace the corresponding env vars and run it!
$ docker run -p 8000:<PORT> -e PORT=<PORT> -e DATABASE_URL=<database_url> -e API_KEY=<api_key> golang-echo-api
```

* If you want to use your local DB to test this build you can use [host.docker.internal](https://docs.docker.com/desktop/networking/#use-cases-and-workarounds) in your database url env var parameter (i.e. `-e DATABASE_URL=postgresql://postgres:postgres@host.docker.internal:5432/postgres`)

