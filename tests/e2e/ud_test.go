package e2e

import (
	"context"
	"fmt"
	"github.com/airbloc/airbloc-go/key"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
	"testing"
)

func TestCreateAccount(t *testing.T) {

}

func TestCreateProxyAccount(t *testing.T) {

}

func TestWatchDAuthEvent(t *testing.T) {

}

func TestChallengeDAuthEvent(t *testing.T) {

}

func TestChallengeDataTransaction(t *testing.T) {
	k, _ := key.Generate()
	log.Println(hexutil.Encode(crypto.FromECDSA(k.PrivateKey)))
}

func testUserSignup(conn *grpc.ClientConn, index int) string {
	dauth := pb.NewDAuthClient(conn)
	req := &pb.SignInRequest{
		Identity:     fmt.Sprintf("test-user-%d@airbloc.org", index),
		UserDelegate: common.FromHex("0x3949EfD67b33D4FEe196BDe8E945acE3F9ee3EE6"),
	}
	resp, err := dauth.SignIn(context.Background(), req)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to create account").Error())
	}
	return resp.GetAccountId()
}

func testDAuth(conn *grpc.ClientConn, collectionId string, accountId string, allow bool) {
	dauth := pb.NewDAuthClient(conn)
	req := &pb.DAuthRequest{
		CollectionId: collectionId,
		AccountId:    accountId,
	}
	dauthModifyMethod := dauth.Allow
	if !allow {
		dauthModifyMethod = dauth.Deny
	}
	_, err := dauthModifyMethod(context.Background(), req)
	if err != nil {
		log.Fatalln(errors.Wrap(err, "failed to modify dauth").Error())
	}
}

func TestUserDelegateMain(t *testing.T) {
	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	appId := testCreateApp(conn)
	log.Println("Created App ID:", appId)

	schemaId := testCreateSchema(conn)
	log.Printf("Created Schema ID: %s\n", schemaId)

	collectionId := testCreateCollection(appId, schemaId, conn)
	log.Printf("Created Collection: %s\n", collectionId)

	// create 10 accounts through DAuth signup
	var accountIds []string
	for i := 1; i <= 10; i++ {
		accountId := testUserSignup(conn, i)
		accountIds = append(accountIds, accountId)
	}

	// DAuth it
	for i, accountId := range accountIds {
		allow := i%2 == 0
		testDAuth(conn, collectionId, accountId, allow)
		log.Printf("%s dauth: allow=%v", accountId, allow)
	}
}
