pragma solidity ^0.4.24;

import "./Exchange.sol";
import "./ExchangeLib.sol";
import "openzeppelin-solidity/contracts/token/ERC20/IERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/SafeERC20.sol";
import "openzeppelin-solidity/contracts/token/ERC20/ERC20Mintable.sol";
import "openzeppelin-solidity/contracts/utils/ReentrancyGuard.sol";

contract SimpleContract is ReentrancyGuard {
    using SafeERC20 for IERC20;

    Exchange private exchange;

    constructor(Exchange _exchange) public {
        exchange = _exchange;
    }

    function transact(address _token, uint256 _amount, bytes8 _offerId) public nonReentrant {
        (
        address from,
        address to,
        address contractAddr
        ) = exchange.getOfferCompact(_offerId);

        require(msg.sender == address(exchange), "should have authority");
        require(contractAddr == address(this), "not this contract");

        IERC20 token = IERC20(_token);
        require(token.allowance(from, address(this)) <= token.balanceOf(from) , "low balance");
        token.safeTransferFrom(from, to, _amount);
    }
}
