Airbloc
==========

This repository contains implementation of Airbloc Protocol and networks using Golang.

## Installation

It is strongly recommended that you use a released version.
Release binaries are available on the [releases](https://github.com/airbloc/airbloc-go/releases) page.

## Building from Source

#### Prerequisites

Go and Node is required to compile clients and smart contracts.

 * [Go](http://golang.com) >= 1.10
 * [dep](https://github.com/golang/dep)
 * [Node.JS](http://nodejs.org) >= 10.0

#### Build

```
 $ make airbloc
```

If you see some errors related to header files on macOS, please install XCode Command Line Tools.

```
xcode-select --install
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
 $ ./generate-proto.sh
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
