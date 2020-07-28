# ezapi
Scaffold generator for APIs and services in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/munjeli/ezapi)](https://goreportcard.com/report/github.com/munjeli/ezapi) ![Go](https://github.com/munjeli/ezapi/workflows/Go/badge.svg)

## Contents
- [Installation](#installation)
- [Usage](#usage)
- [Contributing](#contributing)
- [Team](#team)
- [License](#license)

## Installation
Clone the repo, run `make build`. You can add it to your path if you're going to do a lot of this. I named the executable `eza` because less typing.

## Usage
### make
Use `make` to build and run the code. 
`make build` will build the executable for your os.
`make run` will start servers on `localhost:8080` or, in the case of TLS enabled services, `localhost:443`. `make clean` will kill the server process and delete the executable. 

### generating 
Arguments are positional: API type, name, and target directory where the API or service is generated. API types are `API`, `netsrv` for a standalone API and networked service respectively. A networked service is a service with a single endpoint for communication. This is a design for infrastructure utilities that don't need full CRUD handling. 

`api` by default sets up a server and stubs for basic CRUD and validation.

`netsrv` creates a backend service with a server and single endpoint. 

All stubs include test stubs. `api` and `netsrv` will also generate stubs for integration tests.

Examples:
```
./eza api kitten ./kittensdir
```
Generates stubs for CRUD, a server, tests, and integration tests in a directory at `./kittensdir/apis/kitten/`. 
```
./eza srv kitten ./kittensdir
```
Generates a service for CRUD methods to back the `kitten` API at `./kittensdir/services/kitten/`
```
./eza netsrv puppy ./puputil
```
Generates a service at `./puputil/service/puppy/`, and a server with a handler at `./puputil/server/puppy/`

You can generate multiple APIs, services and networked services in one repository without conflict assuming they're distinctly named. 

`make run` will start a single server if there's just one in `apis` or `server`. Otherwise you need to call the API by name, i.e. `make run kitten` for the kitten API in the examples above. Same for `clean`. `make` will vet, fmt, and build the code to an executable named `eza` in the same dir. 

## Contributing
## Team
## License