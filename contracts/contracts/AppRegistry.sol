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

    /**
     * @dev Creates a new application.
     */
    function register(string _name) public {
        bytes32 hashOfName = keccak256(abi.encodePacked(_name));
        require(!appNameExists[hashOfName], "App name already exists.");
        appNameExists[hashOfName] = true;

        bytes8 appId = Utils.generateId(hashOfName, msg.sender);
        apps[appId].name = _name;
        apps[appId].owner = msg.sender;

        emit Registered(_name, appId);
    }

    function transferAppOwner(bytes8 appId, address _newOwner) public {
        require(isOwner(appId, msg.sender), "only owner can transfer ownership");
        apps[appId].owner = _newOwner;
    }

    function isOwner(bytes8 _appId, address _owner) public view returns (bool) {
        return apps[_appId].owner == _owner;
    }

    function unregister(bytes8 _appId) public {
        require(exists(_appId), "App does not exist.");

        bytes32 hashOfName = keccak256(abi.encodePacked(apps[_appId].name));
        appNameExists[hashOfName] = false;
        delete apps[_appId];
    }

    function exists(bytes8 _appId) public view returns (bool) {
        return apps[_appId].owner != address(0x0);
    }

    function exists(string _appName) external view returns (bool) {
        bytes32 hashOfName = keccak256(abi.encodePacked(_appName));
        return appNameExists[hashOfName];
    }
}
