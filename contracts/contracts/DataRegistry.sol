pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";

contract DataRegistry is Ownable {
    mapping (bytes32 => uint256) dataReg;
    // do nothing at this time (stage 1)
}