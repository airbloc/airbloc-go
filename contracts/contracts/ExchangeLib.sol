pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/utils/Address.sol";

library ExchangeLib {
    using Address for address;

    enum OfferStatus {NEUTRAL, PENDING, SETTLED, REJECTED, OPENED, CLOSED}

    struct Escrow {
        address addr;
        bytes4  openSign;
        bytes   openArgs;
        bytes4  closeSign;
        bytes   closeArgs;
    }

    function exec(
        Escrow storage _escrow,
        bytes4 _sign,
        bytes memory _args,
        bytes20 _offerId
    ) internal returns (bool) {
        bytes memory args = abi.encode(_offerId);
        if (_args.length > 0) {
            args = abi.encodePacked(
                _args,
                _offerId
            );
        }
        return _escrow.addr.delegatecall(
            abi.encodePacked(
                _sign,
                args
            )
        );
    }

    struct Offer {
        address     from;
        address     to;
        bytes20[]   dataIds;
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
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(offer.status == OfferStatus.PENDING, "pending state only");
        require(msg.sender == offer.to, "only offeree can settle offer");
        offer.status = OfferStatus.SETTLED;
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

    function open(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        Escrow storage escrow = offer.escrow;

        require(offer.status == OfferStatus.SETTLED, "settled state only");
        require(msg.sender == offer.to, "only escrow can open transaction");

        offer.status = OfferStatus.OPENED;
        return exec(escrow, escrow.openSign, escrow.openArgs, _offerId);
    }

    function close(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        Escrow storage escrow = offer.escrow;
      
        require(offer.status == OfferStatus.OPENED, "opened state only");
        require(msg.sender == offer.escrow.addr, "only contract can close transaction");

        offer.status = OfferStatus.CLOSED;
        return exec(escrow, escrow.closeSign, escrow.closeArgs, _offerId);
    }

    function getOffer(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal view returns (Offer storage) {
        return _orderbook.orders[_offerId];
    }
}
