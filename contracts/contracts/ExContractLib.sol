pragma solidity ^0.4.0;

library ExContractLib {
    struct Offer {
        Participants participants;
        uint256 amount;
        uint256 timeout;
        address contractAddr;
    }

    struct Participants {
        address offeror;
        address offeree;
        address[] opt;
    }

    function getSummary(Offer storage _offer)
        internal
        view
        returns (address, address, uint256)
    {
        return (
            _offer.participants.offeror,
            _offer.participants.offeree,
            _offer.amount
        );
    }

    struct Order {
        address offeror;
        address offeree;
        address contractAddr;
        bool isPending;
    }

    struct Orderbook {
        mapping(bytes32 => Offer) orders;
        mapping(bytes32 => Order) pending;
    }

    function order(
        Orderbook storage _orderbook,
        Order memory _order
    ) internal returns (bytes32) {
        bytes32 offerId = keccak256(
            abi.encodePacked(
                _order.offeror,
                _order.offeree,
                _order.contractAddr,
                block.number
            )
        );
        _order.isPending = true;
        _orderbook.pending[offerId] = _order;
        return offerId;
    }

    function settle(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal returns (bool) {
        Order storage targetOrder = _orderbook.pending[_offerId];
        if (msg.sender != targetOrder.offeree) {return false;}
        targetOrder.isPending = false;
        return true;
    }

    function reject(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal returns (bool) {
        Order storage targetOrder = _orderbook.pending[_offerId];
        if (msg.sender != targetOrder.offeree) {return false;}
        delete _orderbook.pending[_offerId];
        return true;
    }

    function open(
        Orderbook storage _orderbook,
        bytes32 _offerId,
        Offer memory _offer
    ) internal returns (bool) {
        _orderbook.orders[_offerId] = _offer;
        return true;
    }

    function close(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal returns (bool) {
        delete _orderbook.orders[_offerId];
        return true;
    }

    function getOrder(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal view returns (Order storage) {
        return _orderbook.pending[_offerId];
    }

    function getOffer(
        Orderbook storage _orderbook,
        bytes32 _offerId
    ) internal view returns (Offer storage) {
        return _orderbook.orders[_offerId];
    }
}
