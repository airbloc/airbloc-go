pragma solidity ^0.4.24;
pragma experimental ABIEncoderV2;

import "./Utils.sol";

contract SchemaRegistry {

    event Registered(bytes8 indexed _id, address owner);
    event Unregistered(bytes8 indexed _id);

    struct Schema {
        address owner;
        string name;
    }
    mapping (bytes8 => Schema) schemas;
    mapping (bytes32 => bool) nameExists;

    function register(string _name) public {
        bytes32 hashedName = keccak256(abi.encodePacked(_name));
        require(!nameExists[hashedName], "The schema already exists!");

        bytes8 id = Utils.generateId(hashedName, msg.sender);
        Schema storage schema = schemas[id];

        schema.owner = msg.sender;
        schema.name = _name;
        nameExists[hashedName] = true;

        emit Registered(id, msg.sender);
    }
    
    function unregister(bytes8 _id) public {
        Schema storage schema = schemas[_id];
        require(schema.owner == msg.sender, "Only owner can do this");

        bytes32 hashedName = keccak256(abi.encodePacked(schema.name));
        nameExists[hashedName] = false;

        delete schemas[_id];
        emit Unregistered(_id);
    }

    function get(bytes8 _id) public view returns (Schema memory) {
        require(exists(_id), "Given schema does not exist.");
        return schemas[_id];
    }

    function exists(bytes8 _id) public view returns (bool) {
        return (schemas[_id].owner != address(0));
    }
}