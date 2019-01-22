pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./Utils.sol";
import "./SparseMerkleTree.sol";
import "./Accounts.sol";
import "./CollectionRegistry.sol";

contract DataRegistry is Ownable {

    event BundleUnregistered(bytes8 indexed collectionId, bytes8 indexed bundleId);
    event BundleRegistered(bytes8 indexed collectionId, bytes8 indexed bundleId);
    event Punished(address provider);

    struct Bundle {
        bytes32 usersRoot;
        bytes32 bundleDataHash;
        string uri;
        uint256 createdAt;
        bytes8 collectionId;

        // will be further used in Data Availability Challenge
        uint64 proofOfPosessionCount;
    }
    mapping (bytes8 => mapping (bytes8 => Bundle)) public bundles;

    Accounts accounts;
    CollectionRegistry collections;
    SparseMerkleTree sparseMerkleTree;

    constructor(Accounts _accounts, CollectionRegistry _collections, SparseMerkleTree _smt) public {
        accounts = _accounts;
        collections = _collections;
        sparseMerkleTree = _smt;
    }

    function registerBundle(bytes8 collectionId, bytes32 usersRoot, bytes32 dataHash, string uri) external /*onlyProvider*/ {
        require(collections.exists(collectionId), "Collection does not exist.");

        Bundle memory bundle;
        bundle.usersRoot = usersRoot;
        bundle.bundleDataHash = dataHash;
        bundle.collectionId = collectionId;
        bundle.uri = uri;

        bytes8 empty = "";
        bytes8 bundleId = bytes8(keccak256(abi.encodePacked(block.number, msg.sender, collectionId, dataHash)));
        bundles[empty][bundleId] = bundle;

        emit BundleRegistered(collectionId, bundleId);
    }

    function challenge(bytes8 collectionId, bytes8 bundleId, bytes proof) external view /*onlyConsumer*/ {
        require(collections.exists(collectionId), "Collection does not exist.");
        bytes8 userId = accounts.getAccountId(msg.sender);

        bytes8 empty = "";
        Bundle storage bundle = bundles[empty][bundleId];
        require(
            collections.isCollectionAllowedAt(collectionId, userId, bundle.createdAt),
            "You have been allowed to collect the data at that time. Why is it a problem?"
        );

        bool challengeResult = sparseMerkleTree.checkMembership(bundle.usersRoot, uint64(userId), proof);
        require(challengeResult, "Proof failed");
//        punish(collection.owner);
//        emit Punished(collection.owner);
    }

    function isMyDataIncluded(bytes8 collectionId, bytes8 bundleId, bytes proof) public view returns (bool) {
        require(collections.exists(collectionId), "Collection does not exist.");

        bytes8 empty = "";
        uint64 userId = uint64(accounts.getAccountId(msg.sender));
        bytes32 root = bundles[empty][bundleId].usersRoot;

        return sparseMerkleTree.checkMembership(root, userId, proof);
    }

}
