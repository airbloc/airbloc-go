pragma solidity ^0.4.24;

import "./ExchangeLib.sol";
import "openzeppelin-solidity/contracts/introspection/IERC165.sol";

contract Exchange {
    using ExchangeLib for ExchangeLib.Offer;
    using ExchangeLib for ExchangeLib.Orderbook;

    event OfferPrepared(bytes8 indexed _offerId);
    event OfferPresented(bytes8 indexed _offerId);
    event OfferSettled(bytes8 indexed _offerId);
    event OfferRejected(bytes8 indexed _offerId);
    event Receipt(bytes8 indexed _offerId, address indexed _offeror, address indexed _offeree);

    ExchangeLib.Orderbook orderbook;
    mapping(address => bytes8[]) public toIndex;
    mapping(address => bytes8[]) public fromIndex;
    mapping(address => bytes8[]) public escrowIndex;

    uint256 constant DEFAULT_TIMEOUT = 240; // block = 3600 sec = 60 min = 1 hour
    uint256 constant MAX_OPT_LENGTH = 10;

    constructor() public {
        orderbook = ExchangeLib.Orderbook();
    }

    function prepare(
        address _offeror,
        address _offeree,
        address _escrow,
        bytes4 _sign,
        bytes memory _args,
        bytes16[] memory _dataIds
    ) public {
        require(_offeror != address(0), "invalid offeror address");
        require(_offeree != address(0), "invalid offere address");
        require(_escrow != address(0), "invalid contract address");
        require(IERC165(_escrow).supportsInterface(_sign), "interface not supported");

        bytes8 offerId = orderbook.prepare(
            ExchangeLib.Offer({
                offeror: _offeror,
                offeree: _offeree,
                dataIds: _dataIds,
                escrow: ExchangeLib.Escrow({
                    addr: _escrow,
                    sign: _sign,
                    args: _args
                }),
                status: ExchangeLib.Status.NEUTRAL,
                reverted: false
            })
        );
        emit OfferPrepared(offerId);
    }

    function addDataIds(
        bytes8 _offerId,
        bytes16[] memory _dataIds
    ) public {
        ExchangeLib.Offer storage offer = orderbook.getOffer(_offerId);
        require(offer.status == ExchangeLib.Status.NEUTRAL, "neutral state only");
        require(msg.sender == offer.offeror, "only offeror can modify offer");
        require(_dataIds.length <= 255, "dataIds length exceeded (max 255)");

        for (uint8 i = 0; i < _dataIds.length; i++) {
            offer.dataIds.push(_dataIds[i]);            
        }
    }

    function order(bytes8 _offerId) public {
        // add order options
        orderbook.order(_offerId);
        emit OfferPresented(_offerId);
    }

    function settle(bytes8 _offerId) public {
        // add settle options
        orderbook.settle(_offerId);
        require(orderbook.open(_offerId), "failed to open escrow transaction");
        emit OfferSettled(_offerId);
    }

    function reject(bytes8 _offerId) public {
        orderbook.reject(_offerId);
        emit OfferRejected(_offerId);
    }

    function close(bytes8 _offerId) public returns (bool) {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);
        orderbook.close(_offerId);
        // add some options (timeout, brokers, etc..)
        toIndex[offer.offeree].push(_offerId);
        fromIndex[offer.offeror].push(_offerId);
        escrowIndex[offer.escrow.addr].push(_offerId);
        emit Receipt(_offerId, offer.offeree, offer.offeror);
        return offer.reverted;
    }

    function getReceiptsByOfferor(address _offeror) public view returns (bytes8[] memory) {return toIndex[_offeror];}
    function getReceiptsByOfferee(address _offeree) public view returns (bytes8[] memory) {return fromIndex[_offeree];}
    function getReceiptsByEscrow(address _escrow) public view returns (bytes8[] memory) {return escrowIndex[_escrow];}

    function _getOffer(bytes8 _offerId)
        internal
        view
        returns (ExchangeLib.Offer storage)
    {
        return orderbook.getOffer(_offerId);
    }

    function getOfferCompact(bytes8 _offerId)
        public
        view
        returns (
            address, // offeror
            address, // offeree
            address, // escrow.addr
            bool     // reverted
        )
    {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);
        return (
            offer.offeror,
            offer.offeree,
            offer.escrow.addr,
            offer.reverted
        );
    }

    function getOffer(bytes8 _offerId)
        public
        view
        returns (
            address,         //offeror
            address,         //offeree
            bytes16[] memory, //dataIds
            // Escrow
            address,      // addr
            bytes4,       // sign
            bytes memory, // args
            // Status
            ExchangeLib.Status, // status
            bool                // reverted
        )
    {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);
        ExchangeLib.Escrow storage escrow = offer.escrow;
        return (
            offer.offeror, 
            offer.offeree, 
            offer.dataIds,
            escrow.addr,
            escrow.sign,
            escrow.args,
            offer.status,
            offer.reverted
        );
    }
}
