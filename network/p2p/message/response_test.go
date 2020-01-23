package message

import (
	"testing"

	"github.com/airbloc/airbloc-go/account"

	"github.com/klaytn/klaytn/common"
	"github.com/klaytn/klaytn/common/hexutil"
	"github.com/klaytn/klaytn/crypto"
	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"
)

var _ Message = (*TestResponse)(nil)

type TestResponse struct {
	MessageID uuid.UUID
	Sign      hexutil.Bytes
}

func (resp TestResponse) Read(payload.Reader) (noise.Message, error) { return nil, nil }
func (resp TestResponse) Write() []byte                              { return nil }
func (resp TestResponse) ID() uuid.UUID                              { return resp.MessageID }
func (resp *TestResponse) SetID(id uuid.UUID)                        { resp.MessageID = id }
func (resp TestResponse) Signature() hexutil.Bytes                   { return resp.Sign }
func (resp *TestResponse) SetSignature(sign hexutil.Bytes)           { resp.Sign = sign }

func TestResponseMessageUtils(t *testing.T) {
	Convey("Test Response Message Utils", t, func() {
		key, err := crypto.GenerateKey()
		So(err, ShouldBeNil)
		acc := account.NewKeyedAccount(key)

		Convey("#SignResponseMessage", func() {
			Convey("Should done correctly", func() {
				msg := &TestResponse{MessageID: uuid.NewV4()}
				err = SignResponseMessage(msg, acc)
				So(err, ShouldBeNil)
				So(msg.Sign, ShouldNotEqual, (hexutil.Bytes)(nil))
			})
			Convey("Should fail if response is nil", func() {
				err = SignResponseMessage(nil, acc)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "nil response")
			})
			Convey("Should fail if response.ID is empty", func() {
				err = SignResponseMessage(&TestResponse{}, acc)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "empty uuid")
			})
		})
		Convey("#VerifyResponseMessage", func() {
			Convey("Should done correctly", func() {
				msg := &TestResponse{MessageID: uuid.NewV4()}

				var id [32]byte
				copy(id[:], msg.ID().Bytes())

				sig, err := crypto.Sign(id[:], key)
				So(err, ShouldBeNil)

				msg.SetSignature(sig)

				ok, err := VerifyResponseMessage(msg, acc.Address())
				So(ok, ShouldBeTrue)
				So(err, ShouldBeNil)
			})
			Convey("Should fail if response is nil", func() {
				ok, err := VerifyResponseMessage(nil, common.Address{})
				So(ok, ShouldBeFalse)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "nil response")
			})
			Convey("Should fail if response.ID is empty", func() {
				ok, err := VerifyResponseMessage(&TestResponse{}, common.Address{})
				So(ok, ShouldBeFalse)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "empty uuid")
			})
			Convey("Should fail when ecrecover fails", func() {
				ok, err := VerifyResponseMessage(&TestResponse{
					MessageID: uuid.NewV4(),
					Sign:      []byte{0xde, 0xed, 0xbe, 0xef, 0xde, 0xed, 0xbe, 0xef},
				}, common.Address{})
				So(ok, ShouldBeFalse)
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "failed to recover signer's public key from signature")
			})
		})
	})
}
