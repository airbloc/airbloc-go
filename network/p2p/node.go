package p2p

import (
	"github.com/airbloc/airbloc-go/account"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/transport"
	"github.com/pkg/errors"
)

func NewAirblocNode(host string, port uint16, tp transport.Layer, acc account.Account) (*noise.Node, error) {
	param := noise.DefaultParams()
	if host != "" {
		param.Host = host
	}
	if port != 0 {
		param.Port = port
	}
	if tp != nil {
		param.Transport = tp
	}

	node, err := noise.NewNode(param)
	if err != nil {
		return nil, errors.Wrap(err, "new node")
	}

	node.Set("node.address", acc.TxOpts().From)
	return node, nil
}
