pragma solidity ^0.4.24;

import "./SchemaRegistry.sol";
import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./CollectionRegistry.sol";
import "./Utils.sol";

contract AppRegistry is Ownable {
    using SafeMath for uint256;

    event Registered(string indexed name, bytes8 appId);

    struct App {
        string name;
        address owner;
    }

    mapping (bytes8 => App) public apps;
    mapping (bytes32 => bool) appNameExists;

    function newOwner(bytes8 _appId, address _newOwner) public {
        require(checkOwner(_appId, msg.sender), "only owner can transfer ownership");
        apps[_appId].owner = _newOwner;
    }

    function checkOwner(bytes8 _appId, address _owner) public view returns (bool) {
        return get(_appId).owner == _owner;
    } 

    function register(string _name) public {
        bytes32 hashOfName = keccak256(_name);
        require(!appNameExists[hashOfName], "App name already exists.");
        appNameExists[hashOfName] = true;

        bytes8 appId = Utils.generateId(msg.sender, hashOfName);
        apps[appId].name = _name;
        apps[appId].owner = msg.sender;

        emit Registered(_name, appId);
    }

    function unregister(bytes8 _appId) public {
        require(check(_appId), "App does not exist.");

        bytes32 hashOfName = keccak256(apps[_appId].name);
        appNameExists[hashOfName] = false;
        delete apps[_appId];
    }

    function get(bytes8 _appId) internal view returns (App storage) {
        return apps[_appId];
    }

    function check(bytes8 _appId) public view returns (bool) {
        return (get(_appId).owner != address(0x0));
    }
}
