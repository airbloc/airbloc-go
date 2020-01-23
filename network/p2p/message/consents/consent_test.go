package consents

import (
	"testing"

	ablTypes "github.com/airbloc/airbloc-go/bind/types"

	json "github.com/json-iterator/go"
	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/perlin-network/noise/payload"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

func TestConsentMessages(t *testing.T) {
	Convey("Test Consent Messages", t, func() {
		Convey("Test ConsentRequest", func() {
			testData := map[string]interface{}{
				"message_id": uuid.NewV4(),
				"consent_data": []ablTypes.ConsentData{{
					Action:   0,
					DataType: "testtype",
					Allow:    false,
				}},
			}

			Convey("#Read", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg, err := ConsentRequest{}.Read(payload.NewReader(testBytes))
				So(err, ShouldBeNil)
				So(msg, ShouldHaveSameTypeAs, ConsentRequest{})
				So(msg.(ConsentRequest).MessageID, ShouldEqual, testData["message_id"])
				So(msg.(ConsentRequest).ConsentData, ShouldResemble, testData["consent_data"])
			})
			Convey("#Write", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg := ConsentRequest{
					MessageID:   testData["message_id"].(uuid.UUID),
					ConsentData: testData["consent_data"].([]ablTypes.ConsentData),
				}
				bytes := msg.Write()
				So(bytes, ShouldResemble, testBytes)
			})
			Convey("#ID", func() {
				id := uuid.NewV4()
				msg := ConsentRequest{MessageID: id}
				So(msg.ID(), ShouldEqual, id)
			})
			Convey("#SetID", func() {
				id := uuid.NewV4()
				msg := ConsentRequest{}
				So(msg.ID(), ShouldEqual, uuid.UUID{})
				msg.SetID(id)
				So(msg.ID(), ShouldEqual, id)
			})
		})
		Convey("Test ConsentResponse", func() {
			testData := map[string]interface{}{
				"message_id": uuid.NewV4(),
				"tx_hash":    common.HexToHash("0xdeadbeefdeadbeef"),
				"signature":  hexutil.Bytes(common.HexToHash("0xbeefdeadbeefdead").Bytes()),
			}

			Convey("#Read", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg, err := ConsentResponse{}.Read(payload.NewReader(testBytes))
				So(err, ShouldBeNil)
				So(msg, ShouldHaveSameTypeAs, ConsentResponse{})
				So(msg.(ConsentResponse).MessageID, ShouldEqual, testData["message_id"])
				So(msg.(ConsentResponse).TxHash, ShouldEqual, testData["tx_hash"])
				So(msg.(ConsentResponse).Sign, ShouldResemble, testData["signature"])
			})
			Convey("#Write", func() {
				testBytes, err := json.Marshal(testData)
				So(err, ShouldBeNil)

				msg := ConsentResponse{
					MessageID: testData["message_id"].(uuid.UUID),
					TxHash:    testData["tx_hash"].(common.Hash),
					Sign:      testData["signature"].(hexutil.Bytes),
				}
				bytes := msg.Write()
				So(bytes, ShouldResemble, testBytes)
			})
			Convey("#ID", func() {
				msg := ConsentResponse{MessageID: testData["message_id"].(uuid.UUID)}
				So(msg.ID(), ShouldEqual, testData["message_id"].(uuid.UUID))
			})
			Convey("#SetID", func() {
				msg := ConsentResponse{}
				So(msg.ID(), ShouldEqual, uuid.UUID{})
				msg.SetID(testData["message_id"].(uuid.UUID))
				So(msg.ID(), ShouldEqual, testData["message_id"].(uuid.UUID))
			})
			Convey("#Signature", func() {
				msg := ConsentResponse{Sign: testData["signature"].(hexutil.Bytes)}
				So(msg.Signature(), ShouldResemble, testData["signature"].(hexutil.Bytes))
			})
			Convey("#SetSignature", func() {
				msg := ConsentResponse{}
				So(msg.Signature(), ShouldResemble, hexutil.Bytes(nil))
				msg.SetSignature(testData["signature"].(hexutil.Bytes))
				So(msg.Signature(), ShouldResemble, testData["signature"].(hexutil.Bytes))
			})
		})
	})
}
