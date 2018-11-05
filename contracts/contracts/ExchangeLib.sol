pragma solidity ^0.4.0;

import "openzeppelin-solidity/contracts/AddressUtils.sol";

library ExchangeLib {
    using AddressUtils for address;

    enum Status {NEUTRAL, PENDING, SETTLED, REJECTED, OPENED, CLOSED}

    struct Offer {
        address offeror;
        address offeree;
        address contractAddr;
        Status status;
    }

    struct Orderbook {
        mapping(bytes32 => Offer) orders;
    }

    function order(
        Orderbook storage _orderbook,
        Offer memory _offer
    ) internal returns (bytes32) {
        require(_offer.status == Status.NEUTRAL, "neutral state only");
        require(_offer.contractAddr.isContract(), "not contract address");
        bytes8 offerId = bytes8(
            keccak256(
                abi.encodePacked(
                    block.number,
                    msg.sender,
                    _offer.offeror,
                    _offer.offeree,
                    _offer.contractAddr
                )
            )
        );
        _offer.status = Status.PENDING;
        _orderbook.orders[offerId] = _offer;
        return offerId;
    }

    function settle(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(offer.status == Status.PENDING, "pending state only");
        require(msg.sender == offer.offeree, "only offeree can settle offer");
        offer.status = Status.SETTLED;
    }

    function reject(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(offer.status == Status.PENDING, "pending state only");
        require(msg.sender == offer.offeree, "only offeree can reject offer");
        offer.status = Status.REJECTED;
    }

    function open(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(offer.status == Status.SETTLED, "settled state only");
        require(msg.sender == offer.contractAddr, "only contract can open transaction");
        offer.status = Status.OPENED;
    }

    function close(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(msg.sender == offer.contractAddr, "only contract can close transaction");
        offer.status = Status.CLOSED;
    }

    function getOffer(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal view returns (Offer storage) {
        return _orderbook.orders[_offerId];
    }
}
