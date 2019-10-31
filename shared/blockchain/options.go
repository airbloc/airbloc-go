package blockchain

import "crypto/ecdsa"

const (
	KlaytnBaobab  = "https://api.baobab.klaytn.net:8651"
	KlaytnCypress = "https://api.cypress.klaytn.net:8651"
)

type Options struct {
	Endpoint         string
	Key              *ecdsa.PrivateKey
	FeePayerKey      *ecdsa.PrivateKey
	FeePayerEndpoint string
}
