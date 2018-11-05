pragma solidity ^0.4.24;

import "./ExchangeLib.sol";
import "openzeppelin-solidity/contracts/access/Whitelist.sol";

contract Exchange is Whitelist {
    using ExchangeLib for ExchangeLib.Offer;
    using ExchangeLib for ExchangeLib.Orderbook;

    event OfferPresented(bytes32 indexed _offerId, address _contract);
    event OfferSettled(bytes32 indexed _offerId, address _offeree);
    event OfferRejected(bytes32 indexed _offerId, address _offeree);
    event OfferOpened(bytes32 indexed _offerId);
    event OfferClosed(bytes32 indexed _offerId, address _offeror, address _offeree, bool _reverted);

    ExchangeLib.Orderbook orderbook;
    uint256 constant DEFAULT_TIMEOUT = 240; // block = 3600 sec = 60 min = 1 hour
    uint256 constant MAX_OPT_LENGTH = 10;

    constructor() public {
        orderbook = ExchangeLib.Orderbook();
    }

    function order(
        address _offeror,
        address _offeree,
        address _contract
    ) public {
        require(_offeror != address(0), "invalid offeror address");
        require(_offeree != address(0), "invalid offere address");
        require(_contract != address(0), "invalid contract address");

        bytes32 offerId = orderbook.order(
            ExchangeLib.Offer({
                offeror: _offeror,
                offeree: _offeree,
                contractAddr: _contract,
                status: ExchangeLib.Status.NEUTRAL
            })
        );
        emit OfferPresented(offerId, _contract);
    }

    function settle(bytes32 _offerId) public {
        orderbook.settle(_offerId);
        emit OfferSettled(_offerId, msg.sender);
    }

    function reject(bytes32 _offerId) public {
        orderbook.reject(_offerId);
        emit OfferRejected(_offerId, msg.sender);
    }

    function open(bytes32 _offerId) public {
        orderbook.open(_offerId);
        emit OfferOpened(_offerId);
    }

    function close(bytes32 _offerId) public returns (bool) {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);
        orderbook.close(_offerId);
        bool reverted = false;

        // TODO: add some options (timeout, brokers, etc..)
        emit OfferClosed(_offerId, offer.offeror, offer.offeree, reverted);
        return reverted;
    }

    function _getOffer(bytes32 _offerId)
        internal
        view
        returns (ExchangeLib.Offer storage)
    {
        return orderbook.getOffer(_offerId);
    }

    function getOffer(bytes32 _offerId)
        public
        view
        returns (address, address, address)
    {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);
        return (
            offer.offeror,
            offer.offeree,
            offer.contractAddr
        );
    }
}
