package e2e

import (
	"log"

	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/klaytn/klaytn/crypto"
	"github.com/stretchr/testify/require"
)

func (t *T) prepareEscrow() *pb.Contract {
	return nil
	//self := bind.NewKeyedTransactor(t.config.TransactorPrivateKey)
	//
	//klayent, err := client.DialContext(t.ctx, t.config.EthereumEndpoint)
	//require.NoError(t, err)
	//
	//token, err := testadapter.NewSimpleToken(t.config.DeployedContracts["ERC20Mintable"], common.Hash{}, big.NewInt(0), klayent)
	//require.NoError(t, err)
	//
	//// mint 10000 Tokens
	//tx, err := token.Mint(self, self.From, new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.Ether)))
	//require.NoError(t, err)
	//_, err = klayBind.WaitMined(t.ctx, klayent, tx)
	//require.NoError(t, err)
	//log.Println("10000 new tokens are minted.")
	//
	//// approve SimpleContract (Trade Escrow Contract) to take 10000 tokens from me
	//tx, err = token.Approve(self, t.config.DeployedContracts["SimpleContract"], new(big.Int).Mul(big.NewInt(10000), big.NewInt(params.Ether)))
	//require.NoError(t, err)
	//_, err = klayBind.WaitMined(t.ctx, klayent, tx)
	//require.NoError(t, err)
	//log.Println("Allowed taking 10000 tokens")
	//
	//mintBalance, _ := token.BalanceOf(nil, self.From)
	//log.Println("Currently", new(big.Float).Quo(
	//	new(big.Float).SetInt(mintBalance),
	//	new(big.Float).SetInt(big.NewInt(params.Ether))),
	//	"tokens are minted now.")
	//
	//simpleContract, err := abi.JSON(strings.NewReader(testadapter.ERC20EscrowABI))
	//require.NoError(t, err)
	//
	//// prepare escrow condition details: SimpleContract.transact(ERC20Mintable.address, 100 Tokens)
	//escrowFunc := simpleContract.Methods["transact"]
	//escrowFuncSign := []byte(escrowFunc.Sig())
	//escrowFuncSelector := crypto.Keccak256Hash(escrowFuncSign).Bytes()[:4]
	//// address, uint256, bytes8
	//escrowFuncArgs, err := escrowFunc.Inputs[:len(escrowFunc.Inputs)-1].Pack(
	//	t.config.DeployedContracts["ERC20Mintable"],
	//	new(big.Int).Mul(big.NewInt(100), big.NewInt(params.KLAY)),
	//)
	//require.NoError(t, err)
	//
	//return &pb.Contract{
	//	Type: pb.Contract_SMART,
	//	SmartEscrow: &pb.SmartContract{
	//		Address:    t.config.DeployedContracts["SimpleContract"].Hex(),
	//		EscrowSign: escrowFuncSelector,
	//		EscrowArgs: escrowFuncArgs,
	//	},
	//}
}

// testExchange tests trading bundle data uploaded before.
func (t *T) testExchange(bundleId string) *pb.BundleInfoResponse {
	/**
	Exchange TODOs:
	- Add common api
		- get current account
		- get current node status
		- etc...
	- sign/args generation (or just make input of func to abi)
	*/
	data := pb.NewDataClient(t.conn)
	bundleInfo, err := data.GetBundleInfo(t.ctx, &pb.BundleInfoRequest{BundleId: bundleId})
	require.NoError(t, err)

	exchange := pb.NewExchangeClient(t.conn)
	req := &pb.OrderRequest{
		To:       crypto.PubkeyToAddress(t.config.TransactorPrivateKey.PublicKey).Hex(),
		Contract: t.prepareEscrow(),
		DataIds:  bundleInfo.DataInfoes,
	}
	log.Println(bundleInfo.DataInfoes)

	offerId, err := exchange.Prepare(t.ctx, req)
	require.NoError(t, err)

	log.Println("OfferId :", offerId.GetOfferId())

	_, err = exchange.Order(t.ctx, offerId)
	require.NoError(t, err)

	// accept offer
	receipt, err := exchange.Settle(t.ctx, offerId)
	require.NoError(t, err)

	// print offer result
	d, _ := json.MarshalIndent(receipt, "", "    ")
	log.Println(string(d))

	return bundleInfo
}
