pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./AppRegistry.sol";
import "./SchemaRegistry.sol";
import "./SparseMerkleTree.sol";
import "./Utils.sol";
import "./Accounts.sol";

contract CollectionRegistry {
    using SafeMath for uint256;

    event Registration(address indexed registrar, bytes8 indexed appId, bytes8 collectionId);
    event Unregistration(bytes8 indexed collectionId, bytes8 indexed appId);
    event Allowed(bytes8 indexed collectionId, bytes8 indexed userId);
    event Denied(bytes8 indexed collectionId, bytes8 indexed userId);

    struct Collection {
        bytes8 appId;
        bytes8 schemaId;
        IncentivePolicy policy;
        mapping (bytes8 => Auth) dataCollectionOf;
    }

    struct IncentivePolicy {
        uint256 self;
        uint256 owner;
    }

    struct Auth {
        bool isAllowed;
        uint256 authorizedAt;
    }

    mapping (bytes8 => Collection) collections;

    Accounts accounts;
    AppRegistry apps;
    SchemaRegistry schemas;

    constructor(Accounts _accounts, AppRegistry _appReg, SchemaRegistry _schemaReg) public {
        apps = _appReg;
        schemas = _schemaReg;
        accounts = _accounts;
    }

    function register(bytes8 _appId, bytes8 _schemaId, uint256 _ratio) public {
        require(apps.checkOwner(_appId, msg.sender), "only owner can register collection.");
        require(schemas.exists(_schemaId), "given schema does not exist");

        bytes32 unique = keccak256(abi.encodePacked(_appId, _schemaId, _ratio));
        bytes8 collectionId = Utils.generateId(unique, msg.sender);

        Collection storage collection = collections[collectionId];
        collection.appId = _appId;
        collection.schemaId = _schemaId;

        // calculate with ETH. ex) 35ETH == 0.35%
        collection.policy = IncentivePolicy({
            self: _ratio,
            owner: uint256(100 ether).sub(_ratio)
        });

        emit Registration(msg.sender, _appId, collectionId);
    }

    function unregister(bytes8 _id) public {
        require(exists(_id), "collection does not exist");

        bytes8 appId = collections[_id].appId;
        require(apps.checkOwner(appId, msg.sender), "only owner can register collection.");

        delete collections[_id];
        emit Unregistration(_id, appId);
    }

    function get(bytes8 _id) public view returns (bytes8 appId, bytes8 schemaId, uint256 incentiveRatioSelf) {
        require(exists(_id), "collection does not exist");

        appId = collections[_id].appId;
        schemaId = collections[_id].schemaId;
        incentiveRatioSelf = collections[_id].policy.self;
    }

    function allow(bytes8 _id) public {
        bytes8 userId = accounts.getAccountId(msg.sender);
        modifyAuth(_id, userId, true);

        emit Allowed(_id, userId);
    }

    function allowByPassword(bytes8 _id, bytes passwordSignature) public {
        bytes32 inputHash = keccak256(abi.encodePacked(_id));
        bytes8 userId = accounts.getAccountIdFromSignature(inputHash, passwordSignature);

        modifyAuth(_id, userId, true);
        emit Allowed(_id, userId);
    }

    function deny(bytes8 _id) public {
        bytes8 userId = accounts.getAccountId(msg.sender);

        modifyAuth(_id, userId, false);
        emit Denied(_id, userId);
    }

    function denyByPassword(bytes8 _id, bytes passwordSignature) public {
        bytes32 inputHash = keccak256(abi.encodePacked(_id));
        bytes8 userId = accounts.getAccountIdFromSignature(inputHash, passwordSignature);

        modifyAuth(_id, userId, false);
        emit Denied(_id, userId);
    }

    function modifyAuth(bytes8 _id, bytes8 _userId, bool _allow) internal {
        require(exists(_id), "Collection does not exist.");
        Auth storage auth = collections[_id].dataCollectionOf[_userId];

        if (auth.authorizedAt != 0 && accounts.isTemporary(_userId)) {
            // temporary account can't change DAuth settings that already set.
            revert("The account is currently locked.");
        }

        auth.isAllowed = _allow;
        auth.authorizedAt = block.number;
    }

    function exists(bytes8 _id) public view returns (bool) {
        return (collections[_id].appId != bytes8(0x0));
    }

    function isCollectionAllowed(bytes8 collectionId, bytes8 user) public view returns (bool) {
        return isCollectionAllowedAt(collectionId, user, block.number);
    }

    function isCollectionAllowedAt(bytes8 collectionId, bytes8 user, uint256 blockNumber) public view returns (bool) {
        return collections[collectionId].dataCollectionOf[user].isAllowed
            && collections[collectionId].dataCollectionOf[user].authorizedAt < blockNumber;
    }
}
