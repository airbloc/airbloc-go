pragma solidity ^0.4.0;

library ExchangeLib {
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
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        if (offer.status != Status.PENDING) {return false;}
        if (msg.sender != offer.offeree) {return false;}
        offer.status = Status.SETTLED;
        return true;
    }

    function reject(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        if (offer.status != Status.PENDING) {return false;}       
        if (msg.sender != offer.offeree) {return false;}
        offer.status = Status.REJECTED;
        return true;
    }

    function open(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        if (offer.status != Status.SETTLED) {return false;}
        if (msg.sender != offer.contractAddr) {return false;}
        offer.status = Status.OPENED;
        return true;
    }

    function close(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        if (offer.status != Status.OPENED) {return false;}
        if (msg.sender != offer.contractAddr) {return false;}
        offer.status = Status.CLOSED;
        return true;
    }

    function getOffer(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal view returns (Offer storage) {
        return _orderbook.orders[_offerId];
    }
}
