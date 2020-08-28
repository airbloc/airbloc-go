Airbloc
==========

[![Build Status](https://img.shields.io/github/workflow/status/airbloc/airbloc-go/Test?style=for-the-badge)](https://github.com/airbloc/airbloc-go/actions)

[Airbloc](https://airbloc.org) is user-oriented data exchange protocol which is designed to be scalable, transparent, and able to handle enterprise-grade massive data.

  * **User-Oriented**: Every data on Airbloc is owned and controlled by user, and can be queried as a user profile.
  * **Privacy:** Airbloc protects privacy through data encryption and privacy shield technology, with GDPR compliances.
  * **Decentralized:** Airbloc does not require any central authorities. Instead, it uses Ethereum and its own cryptoeconomic protocol to handle data exchange in a decentralized manner.

For details, please check our [technical whitepaper](https://abr.ge/2ffuu).

This repository contains an implementation of Airbloc client and node using Go.

**Airbloc is under VERY ACTIVE DEVELOPMENT and should be treated as pre-alpha software.** This means it is not meant to be run in production, its APIs are subject to change without warning.

## Getting Started


### Building from Source

#### Prerequisites

Go and Node is required to compile clients and smart contracts.

 * [Go](http://golang.com) == 1.11.5
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


```
 $ airbloc
```

#### Running Airbloc Server
> You may need to set up configurations before running Airbloc.

```
 $ airbloc server
```

### Running Airbloc from Local Environment

Before launching Airbloc network on local, Make sure that:

 * [MongoDB](https://www.mongodb.com) is running on `localhost:27017`
 * [Ethereum](https://ethereum.org) node is running on `localhost:8545`

We recommend using [Docker Compose](https://docs.docker.com/compose/) to launch local environment.
The setup can done in a command by:

```
$ docker-compose run start_dependencies
```

> If you don't use Docker Compose and use your own Ethereum network, you may have to deploy Airbloc contracts manually.
You can see guides on [Contract Documentations](https://github.com/airbloc/airbloc-go/blob/master/contracts/README.md).

After setting up environments, you can launch Airbloc server via:
```
 $ airbloc server
```

You can customize ports and other parameters by passing parameters.
```
 $ airbloc server --port 8080 --config /path/to/your/config.yml
```


## Testing

Test files are located in `test/` directory.

#### Unit Test

```
 $ make test
```

#### Local End-to-End (E2E) Testing


```
 $ docker-compose run start_dependencies; docker-compose run airbloc
 $ make test-e2e
```

## Development


### Regenerating [Protobuf](https://developers.google.com/protocol-buffers/) Binds

If you modify any Protobuf definitions in `proto/`, you should regenerate the protobuf binds.

```
 $ make generate-proto
```

### Generating Python [Protobuf](https://developers.google.com/protocol-buffers/) Binds

```
 $ make generate-python-pb
```

The generated code will be located in `build/gen`.

## License

The airbloc-go project is licensed under the [Apache License v2.0](https://www.apache.org/licenses/LICENSE-2.0),
also included in our repository in the `LICENSE` file.
