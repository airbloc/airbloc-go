Airbloc Contracts
=================
`airbloc/contract` contains on-chain backend contract codes
written in [Solidity](https://solidity.readthedocs.io) v0.5.0.

### Build

To build contracts, use `npm run compile`.

```
$ npm run compile
```

You can see compiled contract bytecodes and output ABIs in `build/` directory.

### Generating Go Binds

You need to regenerate contract binds (placed in `github.com/airbloc/airbloc-go/adapter`) if the contract ABI is changed.
For details, please refer `contracts/generate_adapter.go`.

```
$ cd ..                  # must be run on root directory!
$ make generate-bind
```


### Deploying Contracts

#### Setting Up Local Custom Providers

Before deploying contracts on actual network, you need to create `truffle-config.local.js` to set up your own Web3 provider with your wallets
(e.g. Private Key, Ledger Nano S). The example of the file is:

```js
const fs = require('fs');

const PrivateKeyProvider = require("truffle-privatekey-provider");
const privateKey = fs.readFileSync('../private.key').toString();

module.exports = {
    getProviderOf(network) {
        if (network === 'mainnet') return new PrivateKeyProvider(privateKey, 'https://mainnet.infura.io/v3/SOME_API_KEY');
        else if (network === 'ropsten') return new PrivateKeyProvider(privateKey, 'https://ropsten.infura.io/v3/SOME_API_KEY');
        return new PrivateKeyProvider(privateKey, 'MY_DEFAULT_ENDPOINT');
    },
};
```

#### Deploying

To deploy contracts, you can use `npm run deploy <network_name>` command. The network name can be one of those:

* `mainnet`: Uses Ethereum Mainnet with your provider defined in `truffle-config.local.js`
* `kovan`, `ropsten`, `rinkeby`: Uses Ethereum Testnet with your provider defined in `truffle-config.local.js`
* `local`: Uses local network on http://localhost:8545 with default unlocked account

For example, to deploy contract on ropsten, you can type:

```
$ npm run deploy ropsten
```

After deployment, `deployment.local.json` is generated. The generated JSON file contains all deployed contract address
information, and it is essential to run Airbloc server. You can provide the file path to Airbloc server as an argument.

```
 $ airbloc server --deployment ./contracts/deployment.local.json
```

### Test Using Docker

First, you need to build a new image or pull from `airbloc/contracts`.

```
 $ docker build -t airbloc/contracts .
```

Then, you need to run container. This uses [ganache-cli](https://truffleframework.com/ganache), which is a light Ethereum chain client suitable for testing.

```
 $ docker run -it -p 8545:8545 -p 8500:8500 airbloc/contracts
```

To access to `deployment.json` (Contract Deployment Addresses), you can use `http://localhost:8500` endpoint.

```
$ curl http://localhost:8500 # For Example
{
  "Accounts": "0xd95b1f49c581251f9b62cae463e34120acdddd76",
  "AppRegistry": "0x093735b0603de9e655f876589c8577361b074e1b",
  "SchemaRegistry": "0x7c8a3e4a4513514cf7d83921dbfb2700372c9294",
  "CollectionRegistry": "0x558a7882df946fbfb631654db52b0a25fc6def8e",
  "Exchange": "0x4e6385774f651a9b3317be8433c029cb5f58388f",
  "SimpleContract": "0x9a084041d379cf551539a29f8f431d3936e28532",
  "SparseMerkleTree": "0x49da8fd3c2fd575663b848488dadbe73f7452cd4",
  "DataRegistry": "0x0e101b525652ce6aa43c1f9f71dc0c29b1cb1f37",
  "ERC20Mintable": "0x009fb3ad2a28ea072ecc61c9176feab31cae6c68"
}
```

The endpoint URL can be also provided as an argument to airbloc server.
For detailed use, please refer root `docker-compose.yml`.

```
$ airbloc server --deployment http://localhost:8500/
```
