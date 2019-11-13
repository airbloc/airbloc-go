package session

import (
	"github.com/airbloc/airbloc-go/account"
	ablbind "github.com/airbloc/airbloc-go/bind"
	"github.com/airbloc/airbloc-go/blockchain"
)

type Config struct {
	Account     account.Account
	Client      *blockchain.Client
	Deployments ablbind.Deployments
}
