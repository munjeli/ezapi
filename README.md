# ezapi
Scaffold generator for APIs and services in Go.

[![Go Report Card](https://goreportcard.com/badge/github.com/munjeli/ezapi)](https://goreportcard.com/report/github.com/munjeli/ezapi) ![Go](https://github.com/munjeli/ezapi/workflows/Go/badge.svg) [![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

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

### generating 
Arguments are positional: API type, name, and target directory where the API or service is generated. API types are `API`, `netsrv` for a standalone API and networked service respectively. A networked service is a service with a single endpoint for communication. This is a design for infrastructure utilities that don't need full CRUD handling. 

`api` by default sets up a server and stubs for basic CRUD and validation. A service is implemented for CRUD handling that just hands back a response.

`netsrv` creates a backend service with a server and single endpoint. 

All stubs include test stubs. `api` and `netsrv` will also generate stubs for integration tests.

Examples:
```
./eza api kitten ./kittensdir
```
Generates stubs for CRUD, a server, tests, and integration tests in a directory at `./kittensdir/apis/kittenapi/`. 

```
./eza netsrv puppy ./puputil
```
Generates a service at `./puputil/service/puppy/`, and a server with a handler at `./puputil/server/`

You can generate multiple APIs, services and networked services in one repository without conflict assuming they're distinctly named. 

`make run` will start a single server if there's just one in `apis` or `server`. Otherwise you need to call the API by name, i.e. `make run kitten` for the kitten API in the examples above. Same for `clean`. `make` will vet, fmt, and build the code to an executable named `eza` in the same dir. 

## Contributing
Fork it, write your code, make a PR. 85% test coverage is requested as minimum. 
## Team

| ![Ele Munjeli](https://avatars3.githubusercontent.com/u/1256674?s=200&u=39055650c504feef4d1c8ee662373932d0ccd074&v=4) |
| :---: |
| <a href="https://github.com/munjeli" target="_blank">**Ele Munjeli**</a> |
## License
This project is developed under an Apache 2.0 license. Read it: [LICENSE](https://github.com/munjeli/ezapi/blob/master/LICENSE)