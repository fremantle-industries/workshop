# workshop

Create, manage & improve your automated trading strategies with rich and diverse data sets, a
first class local development experience and a progression story for deployment across clouds.

`workshop` is a [tabletop](https://github.com/fremantle-industries/tabletop) application using
the [kappa](https://milinda.pathirage.org/kappa-architecture.com) architecture. Coordinate your
microservices to process and deliver a large scale data driven application.

[Getting Started](./docs/GETTING_STARTED.md) | [Commands](./docs/COMMANDS.md) | [Architecture](./docs/ARCHITECTURE.md) | [Market Data](./docs/MARKET_DATA.md)

## Install

Use the `go install` command to download, build and install the `workshop` binary into the path
set by your `GOBIN` environment variable.

```bash
> go install github.com/fremantle-industries/workshop@latest
```

## Setup

Initialize a project and `go module` in the current working directory. The `new` command
will generate your project with an application by the same name.

```bash
> workshop new github.com/myuser/myworkshop
creating directory /tmp/myworkshop
creating directory /tmp/myworkshop/docker-compose.yml
creating directory /tmp/myworkshop/Makefile
creating directory /tmp/myworkshop/tabletop.yml
creating template /tmp/myworkshop/README.md
...
```

## Usage

A comprehensive overview of available commands and options are covered in the [documentation](./docs/COMMANDS.md).

```bash
> workshop
CLI interface for workshop. Create, manage & improve your automated trading strategies.

Usage:
  workshop [command]

Available Commands:
  apps        Commands to manage OTP applications
  brokers     List cluster brokers
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  lakes       Commands to manage data lakes available to the cluster
  marketdata  Commands to manage market data
  new         Create a new workshop project
  start       Start all applications

Flags:
  -h, --help      help for workshop
  -v, --version   version for workshop

Use "workshop [command] --help" for more information about a command.
```

```bash
> workshop start
to stop press Ctrl-C
2023/01/26 12:52:01 Application started with pid=<9993D758.0.1011>, app_name="marketdata", node_name="nodename@localhost"
2023/01/26 12:52:01 Application started with pid=<9993D758.0.1013>, app_name="orderbook", node_name="nodename@localhost"
2023/01/26 12:52:01 Application started with pid=<9993D758.0.1014>, app_name="process", node_name="nodename@localhost"
2023/01/26 12:52:01 Application started with pid=<9993D758.0.1015>, app_name="deliver", node_name="nodename@localhost"
```

## Endpoints

| Name                       | Endpoint                                              |
| ---------------------------| ------------------------------------------------------|
| Grafana                    | [`grafana.localhost`](http://grafana.localhost)       |
| Prometheus                 | [`prometheus.localhost`](http://prometheus.localhost) |
| Redpanda Kafka Console     | [`redpanda.localhost`](http://redpanda.localhost)     |
| MinIO Console              | [`minio.localhost`](http://minio.localhost)           |

## Contributing

### Development

Run the default `make` target which downloads dependencies to build and run
`docker compose` on the host machine.

```bash
> make
```

### Test

Run the `workshop` test suite

```bash
> make test
```

## Authors

- Alex Kwiatkowski - alex+git@fremantle.io

## License

`workshop` is released under the [MIT license](./LICENSE.md)
