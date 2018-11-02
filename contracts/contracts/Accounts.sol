pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/ownership/Ownable.sol";
import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "openzeppelin-solidity/contracts/ECRecovery.sol";

contract Accounts is Ownable {
    using SafeMath for uint256;

    event SignUp(
        address indexed owner,
        address indexed proxy,
        bytes8 accountId
    );

    enum Status {
        NONE,
        TEMPORARY,
        CREATED
    }

    struct Account {
        address owner;
        Status status;

        // password support using account proxy
        address proxy;
        address passwordProof;
        bytes4 passwordSalt;

        bytes32 identityHashLock;
    }

    mapping (bytes8 => Account) public accounts;
    mapping (address => bool) public isSignedUp;
    mapping (address => bytes8) private passwordToAccount;
    mapping (address => bytes8) private addressToAccount;

    uint256 public numberOfAccounts;

    constructor() public {
    }

    function create() external {
        require(!isSignedUp[msg.sender], "account already exists");

        bytes8 accountId = newAccount(msg.sender);
        emit SignUp(msg.sender, address(0x0), accountId);
    }

    function createTemporary(address proxy) {
        bytes8 accountId = newAccount(msg.sender);
        accounts[accountId].proxy = proxy;
        accounts[accountId].status = Status.TEMPORARY;
    }

    function createUsingProxy(address owner, address proxy, address passwordProof) external {
        require(!isSignedUp[msg.sender], "account already exists");

        bytes8 accountId = newAccount(msg.sender);
        setPassword(accountId, proxy, passwordProof);

        emit SignUp(owner, proxy, accountId);
    }

    function getAccountId(address sender) returns (bytes8) {
        bytes8 accountId = addressToAccount[sender];
        require(accounts[accountId].status != Status.NONE, "unknown address");
        return accountId;
    }

    function getAccountIdFromSignature(bytes message, bytes signature) returns (bytes8) {
        // TODO: use schnorr signature verification like following code
        // (msg, P, R, s) => require(R == ecadd(ecmul(s, G), ecmul(keccak256(msg, P, R), P))
        //    && Accounts[passwordToAccount[P]].status != Status.NONE);

        bytes32 hash = keccak256(message);
        address recoveredPasswordProof = ECRecovery.recover(hash, signature);

        bytes8 accountId = passwordToAccount[recoveredPasswordProof];
        if (accounts[accountId].status == Status.NONE) {
            revert("password mismatch");
        }
        return accountId;
    }

    function setPassword(bytes8 accountId, address proxy, address passwordProof) public {
        Account storage account = accounts[accountId];
        account.proxy = proxy;
        account.passwordProof = passwordProof;
        passwordToAccount[passwordProof] = accountId;
    }

    function newAccount(address owner) internal returns (bytes8 accountId) {
        // since Ethereum has a nonce in the transaction, this accountId would never colide.
        bytes memory seed = abi.encodePacked(owner, block.number);
        accountId = bytes8(keccak256(seed));

        Account storage account = accounts[accountId];
        account.owner = owner;
        account.status = Status.CREATED;

        isSignedUp[owner] = true;
    }

    function isTemporary(bytes8 accountId) external view returns (bool) {
        return accounts[accountId].status == Status.TEMPORARY;
    }
}
