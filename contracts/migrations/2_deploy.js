const fs = require('fs');

const Accounts = artifacts.require("Accounts");
const AppRegistry = artifacts.require("AppRegistry");
const SchemaRegistry = artifacts.require("SchemaRegistry");
const CollectionRegistry = artifacts.require("CollectionRegistry");
const Exchange = artifacts.require("Exchange");
const SparseMerkleTree = artifacts.require("SparseMerkleTree");
const DataRegistry = artifacts.require("DataRegistry");

const DEPLOYMENT_OUTPUT_PATH = '../deployment.local.json';


module.exports = function (deployer) {
    deployer.then(async () => {
        // contracts without any dependencies will go here:
        await deployer.deploy([
            Accounts,
            AppRegistry,
            SchemaRegistry,
            Exchange,
            SparseMerkleTree,
        ]);
        await deployer.deploy(CollectionRegistry, AppRegistry.address, SchemaRegistry.address);
        await deployer.deploy(DataRegistry, Accounts.address, CollectionRegistry.address, SparseMerkleTree.address);

        const deployments = {
            Accounts: Accounts.address,
            AppRegistry: AppRegistry.address,
            SchemaRegistry: SchemaRegistry.address,
            CollectionRegistry: CollectionRegistry.address,
            Exchange: Exchange.address,
            SparseMerkleTree: SparseMerkleTree.address,
            DataRegistry: DataRegistry.address,
        };

        console.log('Writing deployments to deployment.local.json');
        fs.writeFileSync(DEPLOYMENT_OUTPUT_PATH, JSON.stringify(deployments, null, '  '));

        console.log('You should run "go run generate_adapter.go" if any ABIs have been changed.');
    });
};
