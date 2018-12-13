package e2e

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
)

func testExchangeProcess(
	t *testing.T,
	ctx context.Context,
	req *pb.OrderRequest,
	reject, revert bool,
	aliceConn, bobConn *grpc.ClientConn,
) {
	alice := pb.NewExchangeClient(aliceConn)
	bob := pb.NewExchangeClient(bobConn)

	var receipt *pb.Receipt = nil

	offerId := testExchangePrepare(t, ctx, req, alice)
	testExchangeOrder(t, ctx, offerId, alice)
	switch {
	case reject && revert:
		testExchangeReject(t, ctx, offerId, bob)
		testExchangeCloseRevert(t, ctx, offerId, bob)
		require.True(t, t.Failed())
	case !reject && revert:
		testExchangeSettle(t, ctx, offerId, bob)
		receipt = testExchangeCloseRevert(t, ctx, offerId, bob)
		require.False(t, t.Failed())
	case reject && !revert:
		testExchangeReject(t, ctx, offerId, bob)
		testExchangeCloseNonRevert(t, ctx, offerId, bob)
		require.True(t, t.Failed())
	case !(reject && revert):
		testExchangeSettle(t, ctx, offerId, bob)
		receipt = testExchangeCloseNonRevert(t, ctx, offerId, bob)
		require.False(t, t.Failed())
	}

	require.NotNil(t, receipt)
}

func testExchangePrepare(
	t *testing.T,
	ctx context.Context,
	req *pb.OrderRequest,
	exchange pb.ExchangeClient,
) *pb.OfferId {
	if len(req.DataIds) >= 20 {
		require.FailNow(t, "number data ids should least 20")
	}

	offerID, err := exchange.Prepare(ctx, req)
	require.NoError(t, err)

	offer, err := exchange.GetOffer(ctx, offerID)
	require.NoError(t, err)
	require.Equal(t, pb.Status_NEUTRAL, offer.GetStatus())
	return offerID
}

func testExchangeOrder(
	t *testing.T,
	ctx context.Context,
	offerID *pb.OfferId,
	exchange pb.ExchangeClient,
) {
	_, err := exchange.Order(ctx, offerID)
	require.NoError(t, err)

	offer, err := exchange.GetOffer(ctx, offerID)
	require.NoError(t, err)
	require.Equal(t, pb.Status_PENDING, offer.GetStatus())
}

func testExchangeSettle(
	t *testing.T,
	ctx context.Context,
	offerID *pb.OfferId,
	exchange pb.ExchangeClient,
) {
	_, err := exchange.Settle(ctx, offerID)
	require.NoError(t, err)

	offer, err := exchange.GetOffer(ctx, offerID)
	require.NoError(t, err)
	require.Equal(t, pb.Status_OPENED, offer.GetStatus())
}

func testExchangeReject(
	t *testing.T,
	ctx context.Context,
	offerID *pb.OfferId,
	exchange pb.ExchangeClient,
) {
	_, err := exchange.Reject(ctx, offerID)
	require.NoError(t, err)

	offer, err := exchange.GetOffer(ctx, offerID)
	require.NoError(t, err)
	require.Equal(t, pb.Status_REJECTED, offer.GetStatus())
}

func testExchangeCloseNonRevert(
	t *testing.T,
	ctx context.Context,
	offerID *pb.OfferId,
	exchange pb.ExchangeClient,
) *pb.Receipt {
	receipt, err := exchange.CloseOrder(ctx, offerID)
	assert.NoError(t, err)

	offer, err := exchange.GetOffer(ctx, offerID)
	assert.NoError(t, err)
	assert.Equal(t, pb.Status_CLOSED, offer.GetStatus())
	assert.Equal(t, false, offer.GetReverted())
	return receipt
}

func testExchangeCloseRevert(
	t *testing.T,
	ctx context.Context,
	offerID *pb.OfferId,
	exchange pb.ExchangeClient,
) *pb.Receipt {
	receipt, err := exchange.CloseOrder(ctx, offerID)
	assert.NoError(t, err)

	offer, err := exchange.GetOffer(ctx, offerID)
	assert.NoError(t, err)
	assert.Equal(t, pb.Status_CLOSED, offer.GetStatus())
	assert.Equal(t, true, offer.GetReverted())
	return receipt
}
