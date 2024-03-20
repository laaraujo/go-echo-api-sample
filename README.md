# Golang Echo API

Sample API project made with Golang and Echo

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

## Overview

This project intends to serve as a template to kickstart Go API development with Echo framework.

## Features

- Database schema and migrations with [Goose]()
- Type-safe SQL interfaces with [sqlc]()
- Local development setup with [Docker](), [Docker Compose](), [Make]()
- Live reload with [Air]() in [Docker]()
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
