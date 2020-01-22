package message

import (
	"encoding/json"

	"github.com/airbloc/airbloc-go/network/p2p/message/consents"
	"github.com/airbloc/airbloc-go/network/p2p/message/users"

	"github.com/perlin-network/noise"
	"github.com/perlin-network/noise/payload"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

type Message interface {
	noise.Message
	ID() uuid.UUID
	SetID(id uuid.UUID)
}

type NoResponse struct{ id uuid.UUID }

func (NoResponse) Read(reader payload.Reader) (noise.Message, error) {
	var resp NoResponse
	err := json.NewDecoder(reader).Decode(&resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode no response message")
	}
	return resp, nil
}
func (n NoResponse) Write() []byte {
	respBytes, _ := json.Marshal(n)
	return respBytes
}
func (n NoResponse) ID() uuid.UUID      { return n.id }
func (n NoResponse) SetID(id uuid.UUID) { n.id = id }

var (
	// assertion
	_ = []Message{
		(*NoResponse)(nil),
		(*users.SignUpRequest)(nil),
		(*users.SignUpResponse)(nil),
		(*users.UnlockRequest)(nil),
		(*users.UnlockResponse)(nil),
		(*consents.ConsentRequest)(nil),
		(*consents.ConsentResponse)(nil),
	}

	OpcodeNoResponse              = noise.RegisterMessage(noise.NextAvailableOpcode(), (*NoResponse)(nil))
	OpcodeUsersSignUpRequest      = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.SignUpRequest)(nil))
	OpcodeUsersSignUpResponse     = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.SignUpResponse)(nil))
	OpcodeUsersUnlockRequest      = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.UnlockRequest)(nil))
	OpcodeUsersUnlockResponse     = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.UnlockResponse)(nil))
	OpcodeConsentsConsentRequest  = noise.RegisterMessage(noise.NextAvailableOpcode(), (*consents.ConsentRequest)(nil))
	OpcodeConsentsConsentResponse = noise.RegisterMessage(noise.NextAvailableOpcode(), (*consents.ConsentResponse)(nil))

	RequestOpcodes = []noise.Opcode{
		OpcodeUsersSignUpRequest,
		OpcodeUsersUnlockRequest,
		OpcodeConsentsConsentRequest,
	}
	ResponseOpcodes = []noise.Opcode{
		OpcodeUsersSignUpResponse,
		OpcodeUsersUnlockResponse,
		OpcodeConsentsConsentResponse,
	}
	Opcodes = append(RequestOpcodes, ResponseOpcodes...)
)
