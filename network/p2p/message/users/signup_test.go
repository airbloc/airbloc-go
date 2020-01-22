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

func TestSignUpMessages(t *testing.T) {
	Convey("Test SignUp Messages", t, func() {
		Convey("Test SignUpRequest", func() {
			testData := map[string]interface{}{
				"message_id":    uuid.NewV4(),
				"identity_hash": common.HexToHash("0xdeadbeefdeadbeef"),
			}

			Convey("#Read", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg, err := SignUpRequest{}.Read(payload.NewReader(testBytes))
				So(err, ShouldBeNil)
				So(msg, ShouldHaveSameTypeAs, SignUpRequest{})
				So(msg.(SignUpRequest).MessageID, ShouldEqual, testData["message_id"])
				So(msg.(SignUpRequest).IdentityHash, ShouldEqual, testData["identity_hash"])
			})
			Convey("#Write", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg := SignUpRequest{
					MessageID:    testData["message_id"].(uuid.UUID),
					IdentityHash: testData["identity_hash"].(common.Hash),
				}
				bytes := msg.Write()
				So(bytes, ShouldResemble, testBytes)
			})
			Convey("#ID", func() {
				id := uuid.NewV4()
				msg := SignUpRequest{MessageID: id}
				So(msg.ID(), ShouldEqual, id)
			})
			Convey("#SetID", func() {
				id := uuid.NewV4()
				msg := SignUpRequest{}
				So(msg.ID(), ShouldEqual, uuid.UUID{})
				msg.SetID(id)
				So(msg.ID(), ShouldEqual, id)
			})
		})
		Convey("Test SignUpResponse", func() {
			testData := map[string]interface{}{
				"message_id": uuid.NewV4(),
				"tx_hash":    common.HexToHash("0xdeadbeefdeadbeef"),
				"signature":  hexutil.Bytes(common.HexToHash("0xbeefdeadbeefdead").Bytes()),
			}

			Convey("#Read", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg, err := SignUpResponse{}.Read(payload.NewReader(testBytes))
				So(err, ShouldBeNil)
				So(msg, ShouldHaveSameTypeAs, SignUpResponse{})
				So(msg.(SignUpResponse).MessageID, ShouldEqual, testData["message_id"])
				So(msg.(SignUpResponse).TxHash, ShouldEqual, testData["tx_hash"])
				So(msg.(SignUpResponse).Signature, ShouldResemble, testData["signature"])
			})
			Convey("#Write", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg := SignUpResponse{
					MessageID: testData["message_id"].(uuid.UUID),
					TxHash:    testData["tx_hash"].(common.Hash),
					Signature: testData["signature"].(hexutil.Bytes),
				}
				bytes := msg.Write()
				So(bytes, ShouldResemble, testBytes)
			})
			Convey("#ID", func() {
				id := uuid.NewV4()
				msg := SignUpResponse{MessageID: id}
				So(msg.ID(), ShouldEqual, id)
			})
			Convey("#SetID", func() {
				id := uuid.NewV4()
				msg := SignUpResponse{}
				So(msg.ID(), ShouldEqual, uuid.UUID{})
				msg.SetID(id)
				So(msg.ID(), ShouldEqual, id)
			})
		})
	})
}
