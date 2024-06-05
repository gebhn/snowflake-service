# Snowflake Service

A service for creating and managing Uplite Snowflakes

## Quick Start

Must have [Go compiler][1], [GNU Make][2], and [protoc][3]. Additional dependencies required for code generation will be installed via the make targets if they are not present in your $GOPATH. See the [makefile][4] for an exhaustive list of dependencies.

A [Turso][5] database is required as well and it may be convenient for you as a developer to have access to the CLI as well. See their documentation for more information.

#### Environment

Required environment variables can be found inside of the [example env](./env.example)

```sh
    export TURSO_DB_URL=libsql://your-database-here.turso.io
    export TURSO_DB_TOKEN=y0uR.D4T4BAs3_toK3N
    export GRPC_SERVER_PORT=50051
```

#### Code Generation

Code generation is required for both the data access layer as well as the gRPC/PB files, both are conveniently exposed to you as a developer via the `make generate` target. See the [makefile](./makefile) for dependencies and plugins.

#### Compilation

Example compilation

#### Container

Example container

#### Testing

Example testing

[1]: https://go.dev/
[2]: https://www.gnu.org/software/make/
[3]: https://grpc.io/docs/protoc-installation/
[4]: ./makefile
[5]: https://turso.tech/
