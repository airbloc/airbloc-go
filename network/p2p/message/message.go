package message

import (
	"github.com/airbloc/airbloc-go/network/p2p/message/consents"
	"github.com/airbloc/airbloc-go/network/p2p/message/users"
	"github.com/perlin-network/noise"
	uuid "github.com/satori/go.uuid"
)

type Message interface {
	noise.Message
	ID() uuid.UUID
	SetID(id uuid.UUID)
}

var (
	// assertion
	_ = []Message{
		(*users.SignUpRequest)(nil),
		(*users.SignUpResponse)(nil),
		(*users.UnlockRequest)(nil),
		(*users.UnlockResponse)(nil),
		(*consents.ConsentRequest)(nil),
		(*consents.ConsentResponse)(nil),
	}

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
