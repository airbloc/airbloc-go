Airbloc
==========

[![CircleCI](https://circleci.com/gh/airbloc/token.svg?style=svg)](https://circleci.com/gh/airbloc/airbloc-go)
[![LICENSE](https://img.shields.io/github/license/airbloc/airbloc-go.svg?style=flat-square)](https://github.com/airbloc/airbloc-go/blob/master/LICENSE)
[![Go Version](https://img.shields.io/badge/golang-%3E%3D1.10.0-green.svg?style=flat-square)](https://golang.org/)

[Airbloc](https://airbloc.org) is decentralized data exchange protocol which is designed to be scalable, transparent, and  able to handle enterprise-grade massive data. For details, please check our [technical whitepaper](https://abr.ge/2ffuu).

This repository contains official implementation of Airbloc Protocol and networks using Golang.

**Airbloc is under VERY ACTIVE DEVELOPMENT and should be treated as pre-alpha software.** This means it is not meant to be run in production, its APIs are subject to change without warning.

## Getting Started

It is strongly recommended that you use a released version.
Release binaries are available on the [releases](https://github.com/airbloc/airbloc-go/releases) page.

### Building from Source

#### Prerequisites

Go and Node is required to compile clients and smart contracts.

 * [Go](http://golang.com) >= 1.10
 * [dep](https://github.com/golang/dep)
 * [Node.JS](http://nodejs.org) >= 10.0

#### Build

```
 $ make all
```

If you see some errors related to header files on macOS, please install XCode Command Line Tools.

```
xcode-select --install
```

### Running Airbloc

> NOTE: After BetaNet release

```
 $ airbloc server start
```

#### Launching Local Network

Before launching Airbloc network on local, Make sure that:

 * [airbloc-bigchaindb-proxy](https://github.com/airbloc/airbloc-bigchaindb-proxy) is running on `localhost:9124`
 * [BigchainDB](https://bigchaindb-server.readthedocs.io/en/latest/simple-deployment-template/index.html) >= 2.0 network is running on `localhost:9984`
    * MongoDB used from BigchainDB is running on `localhost:27017`
 * Ethereum node is running on `localhost:8545`

First, smart contracts should be deployed on the local Ethereum network.
You can see guides on [Deploying Contract](#deploying-contract) section.

Second, you need to launch a bootstrap node to initiate P2P network:
```
 $ bootnode
```

> `bootnode` will be run on TCP port 9100. You can customize it via `--port` flag.

Finally, you can launch Airbloc server via:
```
 $ airbloc server start
```

Or you can customize ports and other parameters by passing a configuration file via:
```
 $ airbloc server start --config /path/to/your/config.yml
```


## Testing

Test files are located in `test/` directory.

#### Unit Test

```
 $ make test
```

#### Local End-to-End (E2E) Testing

Before launching Airbloc network on local, Make sure that

 * [airbloc-bigchaindb-proxy](https://github.com/airbloc/airbloc-bigchaindb-proxy) is running on `localhost:9124`
 * [BigchainDB](https://bigchaindb-server.readthedocs.io/en/latest/simple-deployment-template/index.html) >= 2.0 network is running on `localhost:9984`
    * MongoDB used from BigchainDB is running on `localhost:27017`
 * Ethereum node is running on `localhost:8545`

```
 $ go run test/e2e/main.go
```

## Development


### Regenerating [Protobuf](https://developers.google.com/protocol-buffers/) Binds

If you modify Protobuf definitions in `proto/`, you should regenerate the protobuf binds.

```
 $ make generate-proto
```

### Compiling Contract

If there's some changes in contract, contract bindings should be also re-generated.

```
 $ cd contracts/
 $ npm run compile
 $ go run generate_adapter.go
```

### Deploying Contract

First, you should create `contracts/truffle-config.local.js` and implement your own provider.

```js
const fs = require('fs');

// import your own provider (e.g. truffle-privatekey-provider, truffle-ledger-provider)

module.exports = {
    getProviderOf(network) {
        if (network === 'mainnet') return myMainnetProvider;
        if (network === 'ropsten') return myRopstenProvider;
        if (network === 'test') return myTestnetProvider;
    },
};
```

Then you can deploy contracts:

```
 $ cd contracts/
 $ npm run deploy-NETWORK_NAME
 ...
 Writing deployments to deployment.local.json
```

After deploying, `deployment.local.json` file will be generated in your source root.
You can use the deployment configuration to your Airbloc Network using configurations.

## License

The airbloc-go project is licensed under the [Apache License v2.0](https://www.apache.org/licenses/LICENSE-2.0),
also included in our repository in the `COPYRIGHT` file.
