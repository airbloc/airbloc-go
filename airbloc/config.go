package airbloc

import (
	ablbind "github.com/airbloc/airbloc-go/bind"
)

type Config struct {
	Account     Account
	Client      *Client
	Deployments ablbind.Deployments
}
