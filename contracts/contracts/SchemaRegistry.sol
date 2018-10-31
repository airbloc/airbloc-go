pragma solidity ^0.4.24;

contract SchemaRegistry {

    event Registered(bytes32 indexed _id);
    event Unregistered(bytes32 indexed _id);

    mapping (bytes32 => address) reg;

    function register(bytes32 metahash) public {
        bytes32 id = keccak256(
            abi.encodePacked(
                metahash,
                msg.sender,
                block.number
            )
        );
        reg[id] = msg.sender;
        emit Registered(id);
    }
    
    function unregister(bytes32 id) public {
        require(reg[id] == msg.sender, "only owner can do this");
        delete reg[id];
        emit Unregistered(id);
    }

    function check(bytes32 id) public view returns (bool) {
        return (reg[id] != address(0));
    }
}