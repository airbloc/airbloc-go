pragma solidity ^0.4.24;

import "./Exchange.sol";
import "./ExchangeLib.sol";

contract SimpleContract {

    struct Agreement {
        bool offeror;
        bool offeree;
    }

    Exchange private exchange;
    mapping(bytes32 => Agreement) agreements;
    mapping(bytes32 => uint256) balances;

    constructor(Exchange _exchange) public {
        exchange = _exchange;
    }

    function open(bytes32 _offerId) public payable {
        (
            address offeror, 
            , // address offeree,
            address contractAddr
        ) = exchange.getOffer(_offerId);
        require(msg.sender == offeror, "should have authority");
        require(contractAddr == address(this), "not this contract");
        exchange.open(_offerId);
        balances[_offerId] = msg.value;
        agreements[_offerId] = Agreement(false, false);
    }

    function close(bytes32 _offerId) public {
        (
            address offeror,
            address offeree,
            address contractAddr
        ) = exchange.getOffer(_offerId);

        require(contractAddr == address(this), "not this contract");

        bool isOfferor = msg.sender == offeror;
        bool isOfferee = msg.sender == offeree;

        require(
            isOfferor && isOfferee,
            "should have authority");

        if (
            agreements[_offerId].offeror &&
            agreements[_offerId].offeree
        ) {
            bool reverted = exchange.close(_offerId);
            uint256 amount = balances[_offerId];
            delete agreements[_offerId];

            if (!reverted) {
                offeror.transfer(amount);
            } else {
                offeree.transfer(amount);
            }

            return;
        }

        if (isOfferor) {
            agreements[_offerId].offeror = true;
        } else {
            agreements[_offerId].offeree = true;
        }
    }
}