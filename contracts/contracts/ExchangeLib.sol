pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/utils/Address.sol";

library ExchangeLib {
    using Address for address;

    enum OfferStatus {NEUTRAL, PENDING, SETTLED, REJECTED}

    struct Escrow {
        address addr;
        bytes4  sign;
        bytes   args;
    }

    function exec(
        Escrow storage _escrow,
        bytes32 _offerId
    ) internal returns (bool) {
        bytes memory data = abi.encode(_offerId);
        if (_escrow.args.length > 0) {
            data = abi.encodePacked(
                _escrow.args,
                _offerId
            );
        }
        data = abi.encodePacked(_escrow.sign, data);
        return _escrow.addr.call(data);
    }

    struct Offer {
        address     from;
        address      to;
        bytes32[]   dataIds;
        Escrow      escrow;
        OfferStatus status;
        bool        reverted;
    }

    struct Orderbook {
        mapping(bytes8 => Offer) orders;
    }

    function prepare(
        Orderbook storage _orderbook,
        Offer memory _offer
    ) internal returns (bytes8) {
        require(_offer.status == OfferStatus.NEUTRAL, "neutral state only");
        require(_offer.escrow.addr.isContract(), "not contract address");
        bytes8 offerId = bytes8(
            keccak256(
                abi.encodePacked(
                    block.number,
                    msg.sender,
                    _offer.from,
                    _offer.to,
                    _offer.escrow.addr
                )
            )
        );
        _offer.status = OfferStatus.NEUTRAL;
        _orderbook.orders[offerId] = _offer;
        return offerId;
    }

    function order(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(msg.sender == offer.from, "only offeror can order offer");
        require(offer.status == OfferStatus.NEUTRAL, "neutral state only");
        offer.status = OfferStatus.PENDING;
    }

    // settle and open
    function settle(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        Escrow storage escrow = offer.escrow;
        require(offer.status == OfferStatus.PENDING, "pending state only");
        require(msg.sender == offer.to, "only offeree can settle offer");
        offer.status = OfferStatus.SETTLED;
        return exec(escrow, _offerId);
    }

    function reject(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(offer.status == OfferStatus.PENDING, "pending state only");
        require(msg.sender == offer.to, "only offeree can reject offer");
        offer.status = OfferStatus.REJECTED;
    }

    function getOffer(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal view returns (Offer storage) {
        return _orderbook.orders[_offerId];
    }
}
