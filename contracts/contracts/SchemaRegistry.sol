pragma solidity ^0.4.24;

import "./Utils.sol";

contract SchemaRegistry {

    event Registration(address indexed registrar, bytes8 _id);
    event Unregistration(bytes8 indexed _id);

    struct Schema {
        address owner;
        string name;
    }
    mapping (bytes8 => Schema) public schemas;
    mapping (bytes32 => bool) public nameExists;

    function register(string _name) public {
        bytes32 hashedName = keccak256(abi.encodePacked(_name));
        require(!nameExists[hashedName], "The schema already exists!");

        bytes8 id = Utils.generateId(hashedName, msg.sender);
        Schema storage schema = schemas[id];

        schema.owner = msg.sender;
        schema.name = _name;
        nameExists[hashedName] = true;

        emit Registration(msg.sender, id);
    }
    
    function unregister(bytes8 _id) public {
        Schema storage schema = schemas[_id];
        require(schema.owner == msg.sender, "Only owner can do this");

        bytes32 hashedName = keccak256(abi.encodePacked(schema.name));
        nameExists[hashedName] = false;

        delete schemas[_id];
        emit Unregistration(_id);
    }

    function exists(bytes8 _id) public view returns (bool) {
        return (schemas[_id].owner != address(0));
    }
}
