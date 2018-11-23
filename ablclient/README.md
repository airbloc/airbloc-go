# airbloc-go/ablclient
airbloc-go/ablclient is the Go client library and SDK (Software Development Kit) for Airbloc Protocol.

For full compatibility, it is recommended to vendor builds using etcd's vendored packages, using tools like golang/dep, as in vendor directories.

## Install

```
go get github.com/airbloc/airbloc-go/ablclient
```

## Usage

Currently, only the account functionality has been developed. API client support and DAuth support will be added soon.

```go
package main

import (
    "fmt"
    "log"

	"github.com/airbloc/airbloc-go/ablclient"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9124", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	abl := ablclient.NewClient(conn)

    // generate Ethereum wallet address
	priv, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalln("failed to generate a private key", err.Error())
	}

	walletAddress := crypto.PubkeyToAddress(priv.PublicKey)
	password := "myPassword1234!"

	session, err := accounts.Create(walletAddress, password)
	if err != nil {
		log.Fatalln("failed to create Airbloc account", err.Error())
	}

	fmt.Printf("Account ID: %s\n", session.AccountId.String())
}
```
