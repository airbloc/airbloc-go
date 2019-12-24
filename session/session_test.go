package session

//const (
//	feePayerServerEndpoint = "http://localhost:3470"
//	deploymentEndpoint     = "http://localhost:8500"
//	blockchainEndpoint     = "https://api.baobab.klaytn.net:8651"
//)
//
//func newBasicAccount(t *testing.T, key *ecdsa.PrivateKey) account.Account {
//	if key == nil {
//		var err error
//		key, err = crypto.GenerateKey()
//		require.NoError(t, err)
//	}
//	return account.NewKeyedAccount(key)
//}
//
//func newFeePayedAccount(t *testing.T, key *ecdsa.PrivateKey, feePayerUrl string) account.Account {
//	if key == nil {
//		var err error
//		key, err = crypto.GenerateKey()
//		require.NoError(t, err)
//	}
//	acc, err := account.NewKeyedAccountWithFeePayer(key, feePayerUrl)
//	require.NoError(t, err)
//	return acc
//}
//
//func newDeployments(t *testing.T, endpoint string) bind.Deployments {
//	if endpoint == "" {
//		endpoint = deploymentEndpoint
//	}
//	testDeployments, err := bind.GetDeploymentsFrom(endpoint)
//	require.NoError(t, err)
//	return testDeployments
//}
//
//func TestNewSession(t *testing.T) {
//	pctx, cancelParentCtx := context.WithCancel(context.Background())
//	defer cancelParentCtx()
//
//	testClient, err := blockchain.NewClient(pctx, blockchainEndpoint)
//	require.NoError(t, err)
//	defer testClient.Close()
//
//	var (
//		testAccount     = newFeePayedAccount(t, nil, feePayerServerEndpoint)
//		testDeployments = newDeployments(t, "")
//	)
//
//	Convey("Testing Session", t, func(c C) {
//		ctx, cancel := context.WithTimeout(pctx, 1*time.Minute)
//		defer cancel()
//
//		session, err := NewSession(Config{
//			Account:     testAccount,
//			Client:      testClient,
//			Deployments: testDeployments,
//		})
//		c.So(err, ShouldBeNil)
//
//		dataTypeRegistryManager, err := managers.NewDataTypeRegistryManager(session)
//		c.So(err, ShouldBeNil)
//
//		err = dataTypeRegistryManager.Register(
//			ctx, nil,
//			testAccount.TxOpts().From.Hex(),
//			testAccount.TxOpts().From.Hash(),
//		)
//		c.So(err, ShouldBeNil)
//	})
//}
