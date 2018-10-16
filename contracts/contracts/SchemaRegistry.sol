pragma solidity ^0.4.24;

contract SchemaRegistry {
    
    mapping (bytes32 => address) reg;

    function register(bytes32 id) public {
        reg[id] = msg.sender;
    }
    
    function unregister(bytes32 id) public {
        require(reg[id] == msg.sender, "only owner can do this");
        delete reg[id];
    }

    function check(bytes32 id) public view returns (bool) {
        return (reg[id] != address(0));
    }
}