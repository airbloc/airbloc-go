pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/ECRecovery.sol";
import "./Utils.sol";

contract Accounts is Ownable {
    using SafeMath for uint256;
    using ECRecovery for bytes32;

    event SignUp(address indexed owner, bytes8 accountId);
    event TemporaryCreated(address indexed proxy, bytes32 indexed identityHash, bytes8 accountId);
    event Unlocked(bytes32 indexed identityHash, bytes8 indexed accountId, address newOwner);

    enum Status {
        NONE,
        TEMPORARY,
        CREATED
    }

    struct Account {
        address owner;
        Status status;

        address delegate;

        // password support using account proxy
        address proxy;
        address passwordProof;
    }

    mapping (bytes8 => Account) public accounts;
    mapping (address => bytes8) private passwordToAccount;
    mapping (address => bytes8) private addressToAccount;

    mapping (bytes32 => bytes8) public identityHashToAccount;

    uint256 public numberOfAccounts;

    function create() external {
        require(
            addressToAccount[msg.sender] == bytes8(0),
            "you can make only one account per one Ethereum Account");

        bytes8 accountId = Utils.generateId(bytes32(0), msg.sender);
        accounts[accountId].owner = msg.sender;
        accounts[accountId].status = Status.CREATED;

        addressToAccount[msg.sender] = accountId;
        emit SignUp(msg.sender, accountId);
    }

    function createTemporary(bytes32 identityHash) external {
        require(identityHashToAccount[identityHash] == bytes8(0), "account already exists");

        bytes8 accountId = Utils.generateId(identityHash, msg.sender);
        accounts[accountId].proxy = msg.sender;
        accounts[accountId].delegate = msg.sender;
        accounts[accountId].status = Status.TEMPORARY;

        identityHashToAccount[identityHash] = accountId;
        emit TemporaryCreated(msg.sender, identityHash, accountId);
    }

    function unlockTemporary(bytes32 identityPreimage, address newOwner, bytes passwordSignature) external {
        // check that keccak256(identityPreimage) == account.identityHash
        bytes32 identityHash = keccak256(abi.encodePacked(identityPreimage));
        bytes8 accountId = identityHashToAccount[identityHash];

        require(isTemporary(accountId));
        Account storage account = accounts[accountId];

        require(
            msg.sender == account.proxy,
            "account must be unlocked through the account proxy"
        );
        require(
            addressToAccount[msg.sender] == bytes8(0),
            "you can make only one account per one Ethereum Account"
        );
        account.owner = newOwner;
        addressToAccount[newOwner] = accountId;

        bytes memory message = abi.encodePacked(identityPreimage, newOwner);
        setPassword(accountId, message, passwordSignature);

        account.status = Status.CREATED;
        emit Unlocked(identityHash, accountId, newOwner);
    }

    function createUsingProxy(address owner, bytes passwordSignature) external {
        require(
            addressToAccount[owner] == bytes8(0),
            "you can make only one account per one Ethereum Account");

        bytes8 accountId = Utils.generateId(bytes32(owner), msg.sender);
        accounts[accountId].owner = owner;
        accounts[accountId].proxy = msg.sender;
        accounts[accountId].delegate = msg.sender;
        accounts[accountId].status = Status.CREATED;

        bytes memory message = abi.encodePacked(owner);
        setPassword(accountId, message, passwordSignature);

        addressToAccount[owner] = accountId;
        emit SignUp(owner, accountId);
    }

    function setDelegate(address delegate) external {
        // the delegate and the proxy cannot modify delegate.
        // a delegate can be set only through the account owner's direct transaction.
        require(addressToAccount[msg.sender] != bytes8(0),
            "Account does not exist.");

        Account storage account = accounts[addressToAccount[msg.sender]];
        account.delegate = delegate;
    }

    function setPassword(bytes8 accountId, bytes memory message, bytes memory passwordSignature) internal {
        // user uses his/her own password to derive a sign key.
        // since ECRECOVER returns address (not public key itself),
        // we need to use address as a password proof.
        address passwordProof = keccak256(message).recover(passwordSignature);

        // password proof should be unique, since unique account ID is also used for key derivation
        require(passwordToAccount[passwordProof] == bytes8(0x0), "password proof is not unique");

        accounts[accountId].passwordProof = passwordProof;
        passwordToAccount[passwordProof] = accountId;
    }

    function getAccountId(address sender) public view returns (bytes8) {
        bytes8 accountId = addressToAccount[sender];
        require(accounts[accountId].status != Status.NONE, "unknown address");
        return accountId;
    }

    function getAccountIdFromSignature(bytes32 messageHash, bytes signature) public view returns (bytes8) {
        address passwordProof = messageHash.recover(signature);
        bytes8 accountId = passwordToAccount[passwordProof];

        if (accounts[accountId].status == Status.NONE) {
            revert("password mismatch");
        }
        return accountId;
    }

    function isTemporary(bytes8 accountId) public view returns (bool) {
        return accounts[accountId].status == Status.TEMPORARY;
    }

    function isDelegateOf(address sender, bytes8 accountId) public view returns (bool) {
        return accounts[accountId].delegate == sender;
    }
}
