pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./Utils.sol";
import "./SparseMerkleTree.sol";
import "./Accounts.sol";

contract DataRegistry is Ownable {

    bytes32 constant LEAF_INCLUDED = 0x0000000000000000000000000000000000000000000000000000000000000001;

    event BundleRegistered(bytes8 indexed collectionId, uint64 index);
    event Punished(address provider);

    struct Bundle {
        bytes32 usersRoot;
        bytes32 bundleDataHash;
        string uri;
        uint256 createdAt;

        // will be further used in Data Availability Challenge
        uint64 proofOfPosessionCount;
    }

    struct Auth {
        bool allowed;
        uint256 authedAt;
    }

    struct Collection {
        address owner; // App
        mapping (address => Auth) dauth;
        Bundle[] bundles;
    }

    mapping (bytes8 => Collection) collections;

    Accounts accounts;
    SparseMerkleTree sparseMerkleTree;

    constructor(Accounts _accounts, SparseMerkleTree _smt) public {
        accounts = _accounts;
        sparseMerkleTree = _smt;
    }

    function registerBundle(bytes8 collectionId, bytes32 usersRoot, bytes32 dataHash, string uri) external /*onlyProvider*/ {
        Collection storage collection = collections[collectionId];
        require(collection.owner != address(0x0), "Collection does not exist.");

        Bundle storage bundle;
        bundle.usersRoot = usersRoot;
        bundle.bundleDataHash = dataHash;
        bundle.uri = uri;

        uint64 bundleIndex = uint64(collection.bundles.length);
        collection.bundles.push(bundle);

        emit BundleRegistered(collectionId, bundleIndex);
    }

    function challenge(bytes8 collectionId, uint64 bundleIndex, bytes proof) external /*onlyConsumer*/ {
        Collection storage collection = collections[collectionId];
        require(
            collection.owner != address(0x0),
            "Collection does not exist."
        );

        Bundle storage bundle = collection.bundles[bundleIndex];
        require(
            isCollectionAllowedAt(collectionId, msg.sender, bundle.createdAt),
            "You have been allowed to collect the data at that time. Why is it a problem?"
        );

        uint64 userId = uint64(accounts.getAccountId(msg.sender));
        bool challengeIsTrue = sparseMerkleTree.checkMembership(LEAF_INCLUDED, bundle.usersRoot, userId, proof);
        if (challengeIsTrue) {
            // punish(collection.owner);
            emit Punished(collection.owner);
        } else {
            revert("Proof failed");
        }
    }

    function isMyDataIncluded(bytes8 collectionId, uint64 bundleIndex, bytes proof) public view returns (bool) {
        Collection storage collection = collections[collectionId];
        require(
            collection.owner != address(0x0),
            "Collection does not exist."
        );
        uint64 userId = uint64(accounts.getAccountId(msg.sender));
        bytes32 root = collection.bundles[bundleIndex].usersRoot;

        return sparseMerkleTree.checkMembership(LEAF_INCLUDED, root, userId, proof);
    }

    function isCollectionAllowed(bytes8 collectionId, address user) public view returns (bool) {
        return isCollectionAllowedAt(collectionId, user, block.number);
    }

    function isCollectionAllowedAt(bytes8 collectionId, address user, uint256 blockNumber) internal view returns (bool) {
        return collections[collectionId].dauth[user].allowed
            && collections[collectionId].dauth[user].authedAt < blockNumber;
    }
}