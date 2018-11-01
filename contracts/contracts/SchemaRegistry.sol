pragma solidity ^0.4.24;

contract SchemaRegistry {

    event Registered(bytes32 indexed _id);
    event Unregistered(bytes32 indexed _id);

    mapping (bytes32 => address) reg;

    function register() public {
        bytes32 id = keccak256(
            abi.encodePacked(
                msg.sender,
                block.number
            )
        );
        reg[id] = msg.sender;
        emit Registered(id);
    }
    
    function unregister(bytes32 _id) public {
        require(reg[_id] == msg.sender, "only owner can do this");
        delete reg[_id];
        emit Unregistered(_id);
    }

    function check(bytes32 _id) public view returns (bool) {
        return (reg[_id] == address(0));
    }
}