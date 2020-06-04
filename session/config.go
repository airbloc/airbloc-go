package session

import (
	"github.com/airbloc/airbloc-go/account"
	ablbind "github.com/airbloc/airbloc-go/bind"
)

type Config struct {
	Account     account.Account
	Client      ablbind.ContractBackend
	Deployments ablbind.Deployments
}
