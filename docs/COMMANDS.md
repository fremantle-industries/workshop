# Workshop/Commands

```bash
> workshop
CLI interface to manage the workshop platform.

It allows users to execute actions on insights extracted from data that has been streamed,
recorded & analyzed via a single binary categorized by subcommands.

Usage:
  workshop [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  connectors  Commands that manage inbound & outbound connectors
  help        Help about any command
  transforms  Commands that manage transforms from inbound and intermediate topics
  web         Control plane UI & API

Flags:
      --config string   config file (default is $HOME/.workshop.yaml)
  -h, --help            help for workshop
  -v, --version         version for workshop
      --viper           use Viper for configuration (default true)

Use "workshop [command] --help" for more information about a command.
```

## Connectors

```bash
> workshop connectors
Commands that manage inbound & outbound connectors

Usage:
  workshop connectors [command]

Available Commands:
  binance       Commands that manage Binance inbound & outbound connectors
  bit           Commands that manage Bit inbound & outbound connectors
  bitfinex      Commands that manage Bitfinex inbound & outbound connectors
  bitmex        Commands that manage Bitmex inbound & outbound connectors
  bitstamp      Commands that manage Bitstamp inbound & outbound connectors
  bybit         Commands that manage Bybit inbound & outbound connectors
  coinbase      Commands that manage Coinbase inbound & outbound connectors
  deltaexchange Commands that manage Delta Exchange inbound & outbound connectors
  deribit       Commands that manage Deribit inbound & outbound connectors
  gate          Commands that manage Gate inbound & outbound connectors
  gemini        Commands that manage Gemini inbound & outbound connectors
  huobi         Commands that manage Huobi inbound & outbound connectors
  kraken        Commands that manage Kraken inbound & outbound connectors
  kucoin        Commands that manage Kucoin inbound & outbound connectors
  looksrare     Commands that manage LooksRare inbound & outbound connectors
  okx           Commands that manage Okx inbound & outbound connectors

Flags:
  -h, --help   help for connectors

Global Flags:
      --config string   config file (default is $HOME/.workshop.yaml)
      --viper           use Viper for configuration (default true)

Use "workshop connectors [command] --help" for more information about a command.
```

## Transforms

```bash
> workshop transforms
Commands that manage transforms from inbound and intermediate topics

Usage:
  workshop transforms [command]

Available Commands:
  activities  Commands that manage activity transforms

Flags:
  -h, --help   help for transforms

Global Flags:
      --config string   config file (default is $HOME/.workshop.yaml)
      --viper           use Viper for configuration (default true)

Use "workshop transforms [command] --help" for more information about a command.
```

## Web

```bash
> workshop web
Workshop control plane UI & API for system administration

Usage:
  workshop web [command]

Available Commands:
  start       Run the control plane server

Flags:
  -h, --help   help for web

Global Flags:
      --config string   config file (default is $HOME/.workshop.yaml)
      --viper           use Viper for configuration (default true)

Use "workshop web [command] --help" for more information about a command.
```
