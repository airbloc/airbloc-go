const fs = require('fs');

const Accounts = artifacts.require("Accounts");
const AppRegistry = artifacts.require("AppRegistry");
const SchemaRegistry = artifacts.require("SchemaRegistry");
const CollectionRegistry = artifacts.require("CollectionRegistry");
const Exchange = artifacts.require("Exchange");
const SimpleContract = artifacts.require("SimpleContract");
const SparseMerkleTree = artifacts.require("SparseMerkleTree");
const DataRegistry = artifacts.require("DataRegistry");

// test
const ERC20Mintable = artifacts.require("ERC20Mintable");

const DEPLOYMENT_OUTPUT_PATH = '../deployment.local.json';


module.exports = function (deployer, network) {
    deployer.then(async () => {
        // contracts without any dependencies will go here:
        await deployer.deploy([
            Accounts,
            AppRegistry,
            SchemaRegistry,
            Exchange,
            SparseMerkleTree,
        ]);
        await deployer.deploy(CollectionRegistry, Accounts.address, AppRegistry.address, SchemaRegistry.address);
        await deployer.deploy(DataRegistry, Accounts.address, CollectionRegistry.address, SparseMerkleTree.address);
        await deployer.deploy(SimpleContract, Exchange.address);
        // test
        await deployer.deploy(ERC20Mintable);

        const deployments = {
            Accounts: Accounts.address,
            AppRegistry: AppRegistry.address,
            SchemaRegistry: SchemaRegistry.address,
            CollectionRegistry: CollectionRegistry.address,
            Exchange: Exchange.address,
            SimpleContract: SimpleContract.address,
            SparseMerkleTree: SparseMerkleTree.address,
            DataRegistry: DataRegistry.address,
        };

        if (network === 'test' || network === 'local') {
            deployments.ERC20Mintable = ERC20Mintable.address;
        }

        console.log('Writing deployments to deployment.local.json');
        fs.writeFileSync(DEPLOYMENT_OUTPUT_PATH, JSON.stringify(deployments, null, '  '));

        console.log('You should run "make generate-bind" if any ABIs have been changed.');
    });
};
