module github.com/airbloc/airbloc-go

go 1.13

replace (
	git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
	github.com/klaytn/klaytn => github.com/airbloc/klaytn v1.1.1-0.20191108080408-b619a41d5235
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.12.0
)

require (
	github.com/airbloc/airframe v0.0.0-20190822094319-71d50205227c
	github.com/airbloc/logger v1.1.3
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
	github.com/syndtr/goleveldb v1.0.1-0.20190318030020-c3a204f8e965
	github.com/tidwall/pretty v0.0.0-20180105212114-65a9db5fad51 // indirect
	go.mongodb.org/mongo-driver v1.0.0
	golang.org/x/crypto v0.0.0-20190911031432-227b76d455e7
)
