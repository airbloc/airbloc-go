package users

import (
	"testing"

	json "github.com/json-iterator/go"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/perlin-network/noise/payload"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnlockMessages(t *testing.T) {
	Convey("Test Unlock Messages", t, func() {
		Convey("Test UnlockRequest", func() {
			testData := map[string]interface{}{
				"message_id":        uuid.NewV4(),
				"identity_preimage": common.HexToHash("0xdeadbeefdeadbeef"),
				"new_owner":         common.HexToAddress("0xbeefdeadbeefdead"),
			}

			Convey("#Read", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg, err := UnlockRequest{}.Read(payload.NewReader(testBytes))
				So(err, ShouldBeNil)
				So(msg, ShouldHaveSameTypeAs, UnlockRequest{})
				So(msg.(UnlockRequest).MessageId, ShouldEqual, testData["message_id"])
				So(msg.(UnlockRequest).IdentityPreimage, ShouldEqual, testData["identity_preimage"])
				So(msg.(UnlockRequest).NewOwner, ShouldEqual, testData["new_owner"])
			})
			Convey("#Write", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg := UnlockRequest{
					MessageId:        testData["message_id"].(uuid.UUID),
					IdentityPreimage: testData["identity_preimage"].(common.Hash),
					NewOwner:         testData["new_owner"].(common.Address),
				}
				bytes := msg.Write()
				So(bytes, ShouldResemble, testBytes)
			})
			Convey("#ID", func() {
				id := uuid.NewV4()
				msg := UnlockRequest{MessageId: id}
				So(msg.ID(), ShouldEqual, id)
			})
			Convey("#SetID", func() {
				id := uuid.NewV4()
				msg := UnlockRequest{}
				So(msg.ID(), ShouldEqual, uuid.UUID{})
				msg.SetID(id)
				So(msg.ID(), ShouldEqual, id)
			})
		})
		Convey("Test UnlockResponse", func() {
			testData := map[string]interface{}{
				"message_id": uuid.NewV4(),
				"tx_hash":    common.HexToHash("0xdeadbeefdeadbeef"),
				"signature":  hexutil.Bytes(common.HexToHash("0xbeefdeadbeefdead").Bytes()),
			}

			Convey("#Read", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg, err := UnlockResponse{}.Read(payload.NewReader(testBytes))
				So(err, ShouldBeNil)
				So(msg, ShouldHaveSameTypeAs, UnlockResponse{})
				So(msg.(UnlockResponse).MessageId, ShouldEqual, testData["message_id"])
				So(msg.(UnlockResponse).TxHash, ShouldEqual, testData["tx_hash"])
				So(msg.(UnlockResponse).Sign, ShouldResemble, testData["signature"])
			})
			Convey("#Write", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg := UnlockResponse{
					MessageId: testData["message_id"].(uuid.UUID),
					TxHash:    testData["tx_hash"].(common.Hash),
					Sign:      testData["signature"].(hexutil.Bytes),
				}
				bytes := msg.Write()
				So(bytes, ShouldResemble, testBytes)
			})
			Convey("#ID", func() {
				id := uuid.NewV4()
				msg := UnlockResponse{MessageId: id}
				So(msg.ID(), ShouldEqual, id)
			})
			Convey("#SetID", func() {
				id := uuid.NewV4()
				msg := UnlockResponse{}
				So(msg.ID(), ShouldEqual, uuid.UUID{})
				msg.SetID(id)
				So(msg.ID(), ShouldEqual, id)
			})
			Convey("#Signature", func() {
				msg := UnlockResponse{Sign: testData["signature"].(hexutil.Bytes)}
				So(msg.Signature(), ShouldResemble, testData["signature"].(hexutil.Bytes))
			})
			Convey("#SetSignature", func() {
				msg := UnlockResponse{}
				So(msg.Signature(), ShouldResemble, hexutil.Bytes(nil))
				msg.SetSignature(testData["signature"].(hexutil.Bytes))
				So(msg.Signature(), ShouldResemble, testData["signature"].(hexutil.Bytes))
			})
		})
	})
}
