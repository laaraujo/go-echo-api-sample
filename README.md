# Golang Echo API

Sample Rest API project I made while learning some Go and [Echo](https://echo.labstack.com/)

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

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
