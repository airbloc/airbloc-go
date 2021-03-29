package kas

import "github.com/klaytn/klaytn/params"

type Config struct {
	Endpoint        string `default:"https://node-api.klaytnapi.com/v1/klaytn"`
	Network         string `default:"cypress"`
	AccessKeyID     string
	SecretAccessKey string
}

var networkNameToChainIDs = map[string]uint64{
	"cypress": params.CypressNetworkId,
	"baobab":  params.BaobabNetworkId,
}
