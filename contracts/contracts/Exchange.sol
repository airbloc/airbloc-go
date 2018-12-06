pragma solidity ^0.4.24;

import "./ExchangeLib.sol";
import "openzeppelin-solidity/contracts/introspection/IERC165.sol";
import "openzeppelin-solidity/contracts/utils/ReentrancyGuard.sol";

contract Exchange is ReentrancyGuard {
    using ExchangeLib for ExchangeLib.Offer;
    using ExchangeLib for ExchangeLib.Orderbook;

    event OfferPrepared(bytes8 indexed _offerId);
    event OfferPresented(bytes8 indexed _offerId);
    event OfferSettled(bytes8 indexed _offerId);
    event OfferRejected(bytes8 indexed _offerId);
    event Receipt(bytes8 indexed _offerId, address indexed _offeror, address indexed _to);

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
        address _to,
        address _escrow,
        bytes4 _escrowOpenSign,
        bytes memory _escrowOpenArgs,
        bytes4 _escrowCloseSign,
        bytes memory _escrowCloseArgs,
        bytes16[] memory _dataIds
    ) public {
        require(_to != address(0), "invalid offere address");
        require(_escrow != address(0), "invalid contract address");
        require(IERC165(_escrow).supportsInterface(_escrowOpenSign), "open interface not supported");
        require(IERC165(_escrow).supportsInterface(_escrowCloseSign), "close interface not supported");

        bytes8 offerId = orderbook.prepare(
            ExchangeLib.Offer({
                from: msg.sender,
                to: _to,
                dataIds: _dataIds,
                escrow: ExchangeLib.Escrow({
                    addr: _escrow,
                    openSign: _escrowOpenSign,
                    openArgs: _escrowOpenArgs,
                    closeSign: _escrowCloseSign,
                    closeArgs: _escrowCloseArgs
                }),
                status: ExchangeLib.OfferStatus.NEUTRAL,
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
        require(offer.status == ExchangeLib.OfferStatus.NEUTRAL, "neutral state only");
        require(msg.sender == offer.from, "only from can modify offer");
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

    function settle(bytes8 _offerId) public nonReentrant {
        // add settle options
        orderbook.settle(_offerId);
        require(orderbook.open(_offerId), "failed to open escrow transaction");
        emit OfferSettled(_offerId);
    }

    function reject(bytes8 _offerId) public {
        orderbook.reject(_offerId);
        emit OfferRejected(_offerId);
    }

    function close(bytes8 _offerId) public nonReentrant returns (bool) {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);

        require(orderbook.close(_offerId), "failed to close escrow transaction");
        // add some options (timeout, brokers, etc..)
        toIndex[offer.to].push(_offerId);
        fromIndex[offer.from].push(_offerId);
        escrowIndex[offer.escrow.addr].push(_offerId);
        emit Receipt(_offerId, offer.to, offer.from);
        return offer.reverted;
    }

    function getReceiptsByOfferor(address _from) public view returns (bytes8[] memory) {return toIndex[_from];}
    function getReceiptsByOfferee(address _to) public view returns (bytes8[] memory) {return fromIndex[_to];}
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
            address, // from
            address, // to
            address, // escrow.addr
            bool     // reverted
        )
    {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);
        return (
            offer.from,
            offer.to,
            offer.escrow.addr,
            offer.reverted
        );
    }

    function getOffer(bytes8 _offerId)
        public
        view
        returns (
            address,         //from
            address,         //to
            bytes16[] memory, //dataIds
            // Escrow
            address,      // addr
            bytes4,       // open sign
            bytes memory, // open args
            bytes4,       // close sign
            bytes memory, // close args
            // Status
            ExchangeLib.OfferStatus, // status
            bool                     // reverted
        )
    {
        ExchangeLib.Offer storage offer = _getOffer(_offerId);
        ExchangeLib.Escrow storage escrow = offer.escrow;
        return (
            offer.from,
            offer.to,
            offer.dataIds,
            escrow.addr,
            escrow.openSign,
            escrow.openArgs,
            escrow.closeSign,
            escrow.closeArgs,
            offer.status,
            offer.reverted
        );
    }
}
