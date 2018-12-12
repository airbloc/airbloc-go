package e2e

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"testing"

	"github.com/airbloc/airbloc-go/key"
	pb "github.com/airbloc/airbloc-go/proto/rpc/v1/server"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
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
	pub, _ := base64.StdEncoding.DecodeString("BCcAUp859mBwSiXCZU3y931BcdDmR7nuCSaDIxkf0LwXMKLxsuVF6O0O4AdoiZ2enfccMaCfs7reFFg/yOiWk4w=")
	req := &pb.SignInRequest{
		Identity:     fmt.Sprintf("test-user-%d@airbloc.org", index),
		UserDelegate: pub,
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

func TestUserDelegate(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	appId := testCreateApp(t, ctx, conn)
	log.Println("Created App ID:", appId)

	schemaId := testCreateSchema(t, ctx, conn)
	log.Printf("Created Schema ID: %s\n", schemaId)

	collectionId := testCreateCollection(t, ctx, appId, schemaId, conn)
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
