var Migrations = artifacts.require("./Migrations.sol");

module.exports = function(deployer, network, accounts) {
    console.log(accounts[0]);
    deployer.deploy(Migrations);
};
