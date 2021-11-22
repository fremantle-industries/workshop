# Workshop/Commands/Marketdata

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

## sources

```bash
> workshop marketdata sources
+---------------+--------+
| NAME          | STATUS |
+---------------+--------+
| alpaca        | -      |
| gateio        | -      |
| okx           | -      |
| binance       | -      |
| bitfinex      | -      |
| bybit         | -      |
| kraken        | -      |
| kucoin        | -      |
| ledgerx       | -      |
| bitmex        | -      |
| coinbase      | -      |
| polygonio     | -      |
| deltaexchange | -      |
| deribit       | -      |
+---------------+--------+
```

## marketdata products

```bash
> workshop marketdata products
Commands to manage product market data

Usage:
  workshop marketdata products [command]

Available Commands:
  clear       Remove all products from the data lake for the given source
  list        List all products currently in the data lake
  pull        Pull the latest products into the data lake for the given sources
  remove      Remove the given products from the data lake for the given source

Flags:
  -h, --help   help for products

Use "workshop marketdata products [command] --help" for more information about a command.
```

## marketdata subscriptions

```bash
> workshop marketdata subscriptions
Commands to manage subscription market data

Usage:
  workshop marketdata subscriptions [command]

Available Commands:
  add         Subscribe the cluster to the given market data
  list        List the current subscriptions in the cluster
  remove      Unsubscribe the cluster from the given market data

Flags:
  -h, --help   help for subscriptions

Use "workshop marketdata subscriptions [command] --help" for more information about a command.
```
