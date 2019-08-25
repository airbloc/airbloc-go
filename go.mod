module github.com/airbloc/airbloc-go

go 1.12

replace (
	github.com/airbloc/airframe => github.com/frostornge/airframe v0.0.0-20190822082946-296a12f1f9f2
	github.com/ethereum/go-ethereum/crypto => github.com/klaytn/klaytn v1.1.0
	github.com/klaytn/klaytn => github.com/frostornge/klaytn v0.0.0-20190825165342-de85be0ec1b7
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.12.0
)

require (
	cloud.google.com/go v0.37.1
	github.com/airbloc/airframe v0.0.0-20190404074616-f11a21f025cf
	github.com/airbloc/logger v1.1.3
	github.com/aws/aws-sdk-go v1.19.7
	github.com/dgraph-io/badger v1.5.5
	github.com/gin-gonic/gin v1.3.0
	github.com/golang/mock v1.2.0
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.0.0
	github.com/ipfs/go-cid v0.0.1
	github.com/ipfs/go-datastore v0.0.1
	github.com/jackpal/gateway v1.0.5 // indirect
	github.com/jinzhu/configor v1.0.0
	github.com/json-iterator/go v1.1.7
	github.com/klaytn/klaytn v1.1.1
	github.com/libp2p/go-libp2p v0.0.2
	github.com/libp2p/go-libp2p-crypto v0.0.1
	github.com/libp2p/go-libp2p-host v0.0.2
	github.com/libp2p/go-libp2p-kad-dht v0.0.5
	github.com/libp2p/go-libp2p-net v0.0.1
	github.com/libp2p/go-libp2p-peer v0.1.0
	github.com/libp2p/go-libp2p-peerstore v0.0.1
	github.com/libp2p/go-libp2p-protocol v0.0.1
	github.com/libp2p/go-libp2p-pubsub v0.0.1
	github.com/mcuadros/go-defaults v1.1.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/multiformats/go-multiaddr v0.0.2
	github.com/onsi/ginkgo v1.9.0
	github.com/onsi/gomega v1.6.0
	github.com/pkg/errors v0.8.1
	github.com/soheilhy/cmux v0.1.4
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51 // indirect
	github.com/urfave/cli v1.21.0
	github.com/valyala/fasthttp v1.4.0
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4
	google.golang.org/api v0.2.0
	google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	google.golang.org/grpc v1.23.0
)
