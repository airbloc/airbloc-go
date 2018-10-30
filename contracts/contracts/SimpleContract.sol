pragma solidity ^0.4.24;

import "./Exchange.sol";

contract SimpleContract {

    struct Agreement {
        bool offeror;
        bool offeree;
    }

    Exchange private exchange;
    mapping(bytes32 => Agreement) agreements;

    constructor(Exchange _exchange) public {
        exchange = _exchange;
    }

    function open(bytes32 _offerId) public payable {
        (address offeror,,,) = exchange.getOrder(_offerId);
        require(msg.sender == offeror, "should have authority");
        exchange.open(_offerId, new address[](0), msg.value, 0);
    }

    function close(bytes32 _offerId) public {
        (
            address offeror,
            address offeree,
        ) = exchange.getParticipants(_offerId);

        bool isOfferor = msg.sender == offeror;
        bool isOfferee = msg.sender == offeree;

        require(
            isOfferor && isOfferee,
            "should have authority");

        if (
            agreements[_offerId].offeror &&
            agreements[_offerId].offeree
        ) {
            (bool reverted, uint256 amount) = exchange.close(_offerId);

            if (!reverted) {
                offeror.transfer(amount);
            } else {
                offeree.transfer(amount);
            }

            delete agreements[_offerId];
            return;
        }

        if (isOfferor) {
            agreements[_offerId].offeror = true;
        } else {
            agreements[_offerId].offeree = true;
        }
    }
}