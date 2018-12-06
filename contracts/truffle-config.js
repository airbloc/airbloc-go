const fs = require('fs');

let localConfig = {};
let getProviderOf = function(network) { return () => `${network} is not implemented!` };

if (fs.existsSync('truffle-config.local.js')) {
    const local = require('./truffle-config.local.js');
    localConfig = local.config;
    getProviderOf = local.getProviderOf;
}

const config = {
    networks: {
        mainnet: {
            provider: getProviderOf('mainnet'),
            network_id: 1,
            gas: 500000,
            gasPrice: 12000000000 // 12 Gwei
        },
        ropsten: {
            provider: getProviderOf('ropsten'),
            network_id: 3,
            gas: 3000000
        },
        rinkeby: {
            provider: getProviderOf('rinkeby'),
            network_id: 4,
            gas: 3000000
        },
        test: {
            provider: getProviderOf('test'),
            network_id: 1337,
            gas: 3000000
        },
        klaytn: {
            host: '127.0.0.1',
            port: 8551,
            network_id: '1000', // Aspen network id
            gas: 20000000, // transaction gas limit
            gasPrice: 25000000000, // gasPrice of Aspen is 25 Gpeb
        },
    },
    solc: {
        optimizer: {
            enabled: true,
            runs: 200,
        }
    }
};

module.exports = Object.assign(config, localConfig);
