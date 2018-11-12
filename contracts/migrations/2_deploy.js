const Accounts = artifacts.require("./Accounts.sol");
const AppRegistry = artifacts.require("./AppRegistry.sol");
const CollectionRegistry = artifacts.require("./CollectionRegistry.sol");
const DataRegistry = artifacts.require("./DataRegistry.sol");
const Exchange = artifacts.require("./Exchange.sol");
const SchemaRegistry = artifacts.require("./SchemaRegistry.sol");

module.exports = async function(deployer) {
    await deployer.deploy(Accounts);
    await deployer.deploy(AppRegistry);
    await deployer.deploy(SchemaRegistry);

    await deployer.deploy(CollectionRegistry, AppRegistry.deployed(), SchemaRegistry.deployed());
};
