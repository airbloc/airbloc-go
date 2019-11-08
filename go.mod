module github.com/airbloc/airbloc-go

go 1.13

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/klaytn/klaytn => github.com/airbloc/klaytn v1.1.1-0.20191001045213-4c5c6da28f1e
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.12.0
)

require (
	github.com/airbloc/airframe v0.0.0-20190822094319-71d50205227c
	github.com/airbloc/logger v1.1.3
	github.com/aristanetworks/goarista v0.0.0-20190924011532-60b7b74727fd // indirect
	github.com/dgraph-io/badger v1.5.5
	github.com/gin-gonic/gin v1.3.0
	github.com/gopherjs/gopherjs v0.0.0-20190915194858-d3ddacdb130f // indirect
	github.com/json-iterator/go v1.1.7
	github.com/klaytn/klaytn v1.1.1
	github.com/libp2p/go-libp2p-crypto v0.0.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/pkg/errors v0.8.1
	github.com/smartystreets/assertions v1.0.1 // indirect
	github.com/smartystreets/goconvey v0.0.0-20190731233626-505e41936337
	github.com/stretchr/testify v1.4.0
	github.com/syndtr/goleveldb v1.0.0
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	go.uber.org/multierr v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20190909091759-094676da4a83
	golang.org/x/net v0.0.0-20190930134127-c5a3c61f89f3 // indirect
	golang.org/x/tools v0.0.0-20190925020647-22afafe3322a // indirect
)
