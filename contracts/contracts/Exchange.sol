pragma solidity ^0.4.24;

import "./ExchangeLib.sol";
import "openzeppelin-solidity/contracts/AddressUtils.sol";
import "openzeppelin-solidity/contracts/access/Whitelist.sol";

contract Exchange is Whitelist {
    using AddressUtils for address;
    using ExchangeLib for ExchangeLib.Offer;
    using ExchangeLib for ExchangeLib.Orderbook;

    event OfferPresented(bytes32 indexed _offerId, address _contract);
    event OfferSettled(bytes32 indexed _offerId, address _offeree);
    event OfferRejected(bytes32 indexed _offerId, address _offeree);
    event OfferOpened(bytes32 indexed _offerId, uint256 _timeout);
    event OfferClosed(bytes32 indexed _offerId, address _offeror, address _offeree, uint256 _amount, bool _reverted);

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
        require(_contract.isContract(), "invalid contract address");

        bytes32 offerId = orderbook.order(
            ExchangeLib.Offer({
                offeror: _offeror,
                offeree: _offeree,
                contractAddr: _contract,
                state: ExchangeLib.Status.NEUTRAL
            })
        );
        emit OfferPresented(offerId, _contract);
    }

    function settle(bytes32 _offerId) public {
        require(orderbook.settle(_offerId), "settle error");
        emit OfferSettled(_offerId, msg.sender);
    }

    function reject(bytes32 _offerId) public {
        require(orderbook.reject(_offerId), "reject error");
        emit OfferRejected(_offerId, msg.sender);
    }

    function open(
        bytes32 _offerId,
        address[] _opt,
        uint256 _amount,
        uint256 _timeout
    ) public {
        ExchangeLib.Order storage targetOrder = orderbook.getOrder(_offerId);

        require(targetOrder.contractAddr != address(0), "order not found");
        require(!targetOrder.isPending, "order not settled");
        require(msg.sender == targetOrder.contractAddr, "invalid sender");
        require(_opt.length <= MAX_OPT_LENGTH, "opt length too long");

        uint256 timeout = _timeout;
        if (_timeout == 0) {timeout = DEFAULT_TIMEOUT;}
        timeout += block.number;

        bool result = orderbook.open(
            _offerId,
            ExchangeLib.Offer({
                participants: ExchangeLib.Participants({
                    offeror: targetOrder.offeror,
                    offeree: targetOrder.offeree,
                    opt: _opt
                }),
                amount: _amount,
                timeout: timeout,
                contractAddr: targetOrder.contractAddr
            })
        );
        require(result, "open error");
        emit OfferOpened(_offerId, timeout);
    }

    function close(bytes32 _offerId) public returns (bool, uint256) {
        ExchangeLib.Offer storage offer = orderbook.getOffer(_offerId);
        (
            address offeror,
            address offeree,
            uint256 amount
        ) = offer.getSummary();

        require(offer.amount != 0, "order not found");
        require(msg.sender == offer.contractAddr, "invalid sender");

        bool reverted = block.number > offer.timeout;
        require(orderbook.close(_offerId), "close error");
        emit OfferClosed(_offerId, offeror, offeree, amount, reverted);

        return (reverted, amount);
    }

    function getOffer(bytes32 _offerId)
        public
        view
        returns (address, address, address[], uint256, uint256, address)
    {
        ExchangeLib.Offer storage offer = orderbook.getOffer(_offerId);
        ExchangeLib.Participants storage participants = offer.participants;
        return (
            participants.offeror,
            participants.offeree,
            participants.opt,
            offer.amount,
            offer.timeout,
            offer.contractAddr
        );
    }

    function getParticipants(bytes32 _offerId)
        public
        view
        returns (address, address, address[])
    {
        ExchangeLib.Participants storage participants = orderbook.getOffer(_offerId).participants;
        return (
            participants.offeror,
            participants.offeree,
            participants.opt
        );
    }

    function getOrder(bytes32 _offerId)
        public
        view
        returns (address, address, address, bool)
    {
        ExchangeLib.Order storage targetOrder = orderbook.getOrder(_offerId);
        return (
            targetOrder.offeror,
            targetOrder.offeree,
            targetOrder.contractAddr,
            targetOrder.isPending
        );
    }
}
