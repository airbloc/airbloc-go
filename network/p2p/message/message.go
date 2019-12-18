package message

import (
	"github.com/airbloc/airbloc-go/network/p2p/message/consents"
	"github.com/airbloc/airbloc-go/network/p2p/message/users"
	"github.com/perlin-network/noise"
)

var (
	OpcodeUsersSignUpRequest      = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.SignUpRequest)(nil))
	OpcodeUsersSignUpResponse     = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.SignUpResponse)(nil))
	OpcodeUsersUnlockRequest      = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.UnlockRequest)(nil))
	OpcodeUsersUnlockResponse     = noise.RegisterMessage(noise.NextAvailableOpcode(), (*users.UnlockResponse)(nil))
	OpcodeConsentsConsentRequest  = noise.RegisterMessage(noise.NextAvailableOpcode(), (*consents.ConsentRequest)(nil))
	OpcodeConsentsConsentResponse = noise.RegisterMessage(noise.NextAvailableOpcode(), (*consents.ConsentResponse)(nil))

	// OpcodeMap is opcode to message string list for logging & debugging
	OpcodeMap = map[noise.Opcode]string{
		OpcodeUsersSignUpRequest:      "Message_Users_SignUp_Request",
		OpcodeUsersSignUpResponse:     "Message_Users_SignUp_Response",
		OpcodeUsersUnlockRequest:      "Message_Users_Unlock_Request",
		OpcodeUsersUnlockResponse:     "Message_Users_Unlock_Response",
		OpcodeConsentsConsentRequest:  "Message_Consents_Consent_Request",
		OpcodeConsentsConsentResponse: "Message_Consents_Consent_Response",
	}

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
