pragma solidity ^0.4.24;

import "./SchemaRegistry.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./CollectionRegistry.sol";

contract AppRegistry is Ownable {
    using SafeMath for uint256;

    struct App {
        bytes32 id;
        // TODO: acl
        address owner;
    }

    mapping (bytes32 => App) reg;

    function newApp(bytes32 _appId) internal view returns (App memory) {
        require(check(_appId), "app already exists");
        return App({
            id: _appId,
            owner: msg.sender
        });
    }

    function newOwner(bytes32 _appId, address _newOwner) public {
        require(checkOwner(_appId, msg.sender), "only owner can transfer ownership");
        reg[_appId].owner = _newOwner;
    }

    function checkOwner(bytes32 _appId, address _owner) public view returns (bool) {
        return get(_appId).owner == _owner;
    } 

    function register(bytes32 _appId) public {
        reg[_appId] = newApp(_appId);
    }
    
    function unregister(bytes32 _appId) public onlyOwner {
        delete reg[_appId];
    }

    function get(bytes32 _appId) internal view returns (App storage) {
        return reg[_appId];
    }

    function check(bytes32 _appId) public view returns (bool) {
        return (get(_appId).id != bytes32(0x0));
    }
}