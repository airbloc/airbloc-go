pragma solidity ^0.4.24;

library Utils {

    function generateId(address creator, bytes32 uniqueData) returns (bytes8) {
        bytes memory seed = abi.encodePacked(creator, block.number, uniqueData);
        return bytes8(keccak256(seed));
    }
}
