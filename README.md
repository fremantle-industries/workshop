# Workshop

Improve your automated trading strategies with rich and diverse data sets
that you can run locally or in the cloud.

## About

`workshop` is a high performance, developer first trading toolset executable
as a single CLI binary using the [Kappa](https://milinda.pathirage.org/kappa-architecture.com)
architecture on any Kafka API compliant streaming platform such as
[Kafka](https://kafka.apache.org) itself, [Redpanda](https://redpanda.com) or
[Microsoft Event Hubs](https://azure.microsoft.com/en-us/products/event-hubs/#overview).

```bash
> workshop
CLI interface to manage the workshop platform.

It allows users to execute actions on insights extracted from data that has been streamed,
recorded & analyzed via a single binary categorized by subcommands.

Usage:
  workshop [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  compose     Commands that manage local development through docker compose
  connectors  Commands that manage inbound & outbound connectors
  help        Help about any command
  topics      Commands that manage topics required to operate workshop
  transforms  Commands that manage transforms from inbound and intermediate topics
  web         Control plane UI & API

Flags:
      --config string   config file (default is $HOME/.workshop.yaml)
  -h, --help            help for workshop
  -v, --version         version for workshop
      --viper           use Viper for configuration (default true)

Use "workshop [command] --help" for more information about a command.
```

For a full overview of available commands you can read our in depth
[command](./docs/COMMANDS.md) documentation.

## Development

```bash
> make
```

## Test

```bash
> make test
```

## Endpoints

| Name                       | Endpoint                                                  |
| ---------------------------| ----------------------------------------------------------|
| Workshop Control Plane UI  | [`workshop.localhost`](http://workshop.localhost)         |
| Workshop Control Plane API | [`api.workshop.localhost`](http://api.workshop.localhost) |
| Tabletop Control Plane UI  | [`tabletop.localhost`](http://workshop.localhost)         |
| Tabletop Control Plane API | [`api.tabletop.localhost`](http://api..tabletoplocalhost) |
| Grafana                    | [`grafana.localhost`](http://grafana.localhost)           |
| Prometheus                 | [`prometheus.localhost`](http://prometheus.localhost)     |
| Redpanda Kafka Console     | [`redpanda.localhost`](http://redpanda.localhost)         |
| MinIO Console              | [`minio.localhost`](http://minio.localhost)               |

## Authors

- Alex Kwiatkowski - alex+git@fremantle.io

## License

`workshop` is released under the [MIT license](./LICENSE.md)
