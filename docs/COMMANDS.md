# Workshop/Commands

The workshop CLI is the primary interface to manage your cluster

```bash
> workshop
CLI interface for workshop. Create, manage & improve your automated trading strategies
with rich and diverse data sets, a first class local development experience and a progression
story for deployment across clouds.

workshop is a tabletop application using the kappa architecture. Coordinate your microservices to
process and deliver a large scale data driven application.

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

## apps

```bash
> workshop apps
Commands to manage OTP applications

Usage:
  workshop apps [command]

Available Commands:
  list        List all applications
  start       Start applications
  stop        Stop applications

Flags:
  -h, --help   help for apps

Use "workshop apps [command] --help" for more information about a command.
```

## brokers

```bash
> workshop brokers
TODO> brokers
```

## lakes

```bash
> workshop lakes
Commands to manage data lakes available to the cluster

Usage:
  workshop lakes [command]

Available Commands:
  list        List all data lakes configured in the cluster
  remove      Remove a data lake from the cluster
  set         Set the configuration for a data lake

Flags:
  -h, --help   help for lakes

Use "workshop lakes [command] --help" for more information about a command.
```

## marketdata

Full documentation on [marketdata subcommands](./subcommands/MARKETDATA_SUBCOMMANDS.md)

```bash
> workshop marketdata
Commands to manage market data

Usage:
  workshop marketdata [command]

Available Commands:
  products      Commands to manage product market data
  sources       List all market data sources available to the node
  subscriptions Commands to manage subscription market data

Flags:
  -h, --help   help for marketdata

Use "workshop marketdata [command] --help" for more information about a command.
```

## start

```bash
> workshop start
to stop press Ctrl-C
2023/01/26 16:25:13 Application started with pid=<9993D758.0.1011>, app_name="marketdata", node_name="nodename@localhost"
2023/01/26 16:25:13 Application started with pid=<9993D758.0.1013>, app_name="orderbook", node_name="nodename@localhost"
2023/01/26 16:25:13 Application started with pid=<9993D758.0.1014>, app_name="process", node_name="nodename@localhost"
2023/01/26 16:25:13 Application started with pid=<9993D758.0.1015>, app_name="deliver", node_name="nodename@localhost"
```
