# Pockets

> A tool to manage multiple [Pocketbase](https://pocketbase.io) instances

## Requirements

### Cli

### Server

- Local Docker installation

## Stack

### Cli

- Cobra
- Charm's Huh

### Server

- Chi Router
- Docker Sdk

## Configuration

### Cli

#### Config Values

- ENDPOINT: the endpoint of the pockets server you are connecting to (defaults to: http://localhost:3000)

You have multiple for ways to configure the cli:
- config file
- env variables


#### Config file

The config file should be a `pockets.{toml,yaml,json}` file located either:
- in the same directory as the executable
- in `$HOME/.config/pockets/pockets.{toml,yaml,json}`

#### Env Variables

You can also use env variables. Each Env Variable should be prefixed with `PKTS_`

#### Example

A **pockets.toml** file could look like this:
```toml
ENDPOINT="http://localhost:3000"
```

The same could be achieved with an environment variable setup like this:
```sh
export PKTS_ENDPOINT="http://localhost:3000"
```
