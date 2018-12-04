pragma solidity ^0.4.24;

import "openzeppelin-solidity/contracts/utils/Address.sol";

library ExchangeLib {
    using Address for address;

    enum Status {NEUTRAL, PENDING, SETTLED, REJECTED, OPENED, CLOSED}

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
        bytes memory _args
    ) internal returns (bool) {
        return _escrow.addr.delegatecall(
            abi.encodePacked(
                _sign,
                _args      
            )
        );
    }

    struct Offer {
        address  from;
        address  to;
        bytes16[] dataIds;
        Escrow   escrow;
        Status   status;
        bool     reverted;
    }

    struct Orderbook {
        mapping(bytes8 => Offer) orders;
    }

    function prepare(
        Orderbook storage _orderbook,
        Offer memory _offer
    ) internal returns (bytes8) {
        require(_offer.status == Status.NEUTRAL, "neutral state only");
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
        _offer.status = Status.NEUTRAL;
        _orderbook.orders[offerId] = _offer;
        return offerId;
    }

    function order(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(msg.sender == offer.from, "only offeror can order offer");
        require(offer.status == Status.NEUTRAL, "neutral state only");
        offer.status = Status.PENDING;
    }

    // settle and open
    function settle(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(offer.status == Status.PENDING, "pending state only");
        require(msg.sender == offer.to, "only offeree can settle offer");
        offer.status = Status.SETTLED;
    }

    function reject(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal {
        Offer storage offer = _orderbook.orders[_offerId];
        require(offer.status == Status.PENDING, "pending state only");
        require(msg.sender == offer.to, "only offeree can reject offer");
        offer.status = Status.REJECTED;
    }

    function open(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        Escrow storage escrow = offer.escrow;

        require(offer.status == Status.SETTLED, "settled state only");
        require(msg.sender == offer.to, "only escrow can open transaction");

        offer.status = Status.OPENED;
        return exec(escrow, escrow.openSign,escrow.openArgs);
    }

    function close(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal returns (bool) {
        Offer storage offer = _orderbook.orders[_offerId];
        Escrow storage escrow = offer.escrow;
      
        require(offer.status == Status.OPENED, "opened state only");
        require(msg.sender == offer.escrow.addr, "only contract can close transaction");

        offer.status = Status.CLOSED;
        return exec(escrow, escrow.closeSign, escrow.closeArgs);
    }

    function getOffer(
        Orderbook storage _orderbook,
        bytes8 _offerId
    ) internal view returns (Offer storage) {
        return _orderbook.orders[_offerId];
    }
}
