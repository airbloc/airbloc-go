package p2p

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestWorkers(t *testing.T) {
	Convey("Test Workers", t, func() {
		Convey("MessageAggregator", func() {
			Convey("Should launch exact number of aggregator", func() {})
			Convey("Should aggregate all messages", func() {})
			Convey("Should terminate all workers when context has done", func() {})
		})
		Convey("MessageMultiplexer", func() {
			Convey("Should termiate worker when context has done", func() {})
			Convey("Should handle message match with correct handler", func() {})
		})
	})
}
