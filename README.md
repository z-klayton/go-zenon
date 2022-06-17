# Zenon Node

Reference Golang implementation of the Alphanet - Network of Momentum Phase 0.

## Building from source

Building `znnd` requires both a Go (version 1.16 or later) and a C compiler. You can install them using your favourite package manager. Once the dependencies are installed, please run:

```shell
make znnd
```

## Running `znnd`

Since version `0.0.2`, `znnd` is configured with the Alphanet Genesis and default seeders.

Use [znn-controller](https://github.com/zenon-network/znn_controller_dart) to configure your full node. For more information please consult the [Wiki](https://github.com/zenon-network/znn-wiki).


## Enabling and configuring HTTPS/WSS RPC endpoints

Enabling secure communication over HTTPS and WSS for the go-zenon RPC endpoints requires having a domain and a valid SSL certificate for that domain.

Here is an example config.json:

```json
{
  "RPC": {
    "EnableHTTPS": true,
    "HTTPSHost": "node.domain.com",
    "HTTPSVirtualHosts": [
      "node.domain.com"
    ],
    "HTTPSPort": 35991,
    "HTTPCors": [
      "*"
    ],
    "EnableWSS": true,
    "WSSHost": "node.domain.com",
    "WSSPort": 35992,
    "WSOrigins": [
      "*"
    ],
    "TLS": {
      "Key": "/path/to/privkey.pem",
      "Certificate": "/path/to/fullchain.pem"
    }
  }
}
```
