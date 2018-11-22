pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "./Utils.sol";
import "./SparseMerkleTree.sol";
import "./Accounts.sol";
import "./CollectionRegistry.sol";

contract DataRegistry is Ownable {

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
    mapping (bytes8 => Bundle[]) bundles;

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
        bundle.uri = uri;

        uint64 bundleIndex = uint64(bundles[collectionId].length);
        bundles[collectionId].push(bundle);

        emit BundleRegistered(collectionId, bundleIndex);
    }

    function challenge(bytes8 collectionId, uint64 bundleIndex, bytes proof) external view /*onlyConsumer*/ {
        require(collections.exists(collectionId), "Collection does not exist.");
        bytes8 userId = accounts.getAccountId(msg.sender);

        Bundle storage bundle = bundles[collectionId][bundleIndex];
        require(
            collections.isCollectionAllowedAt(collectionId, userId, bundle.createdAt),
            "You have been allowed to collect the data at that time. Why is it a problem?"
        );

        bool challengeIsTrue = sparseMerkleTree.checkMembership(bundle.usersRoot, uint64(userId), proof);
        if (challengeIsTrue) {
            // punish(collection.owner);
//            emit Punished(collection.owner);
        } else {
            revert("Proof failed");
        }
    }

    function isMyDataIncluded(bytes8 collectionId, uint64 bundleIndex, bytes proof) public view returns (bool) {
        require(collections.exists(collectionId), "Collection does not exist.");

        uint64 userId = uint64(accounts.getAccountId(msg.sender));
        bytes32 root = bundles[collectionId][bundleIndex].usersRoot;

        return sparseMerkleTree.checkMembership(root, userId, proof);
    }

}
