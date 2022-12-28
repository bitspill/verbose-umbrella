OIP Daemon
===

Process that monitors the FLO Blockchain indexing property formatted OIP messages and provides search and retrieval of
OIP records (aka artifacts).

## Requirements
- Go 1.18+
- Elasticsearch 7+
- Flod

## Build Instructions

1. Clone repository to desired directory `git clone https://github.com/oipwg/oip`
2. `cd oip`
3. `go mod download`
4. `go build ./cmd/oipd`
5. The executable is `oipd` built in the root directory of the project
6. Run tests with `go test -v -race`

