package message

import "github.com/perlin-network/noise"

var (
	OpcodeUsersSignUpRequest      = noise.RegisterMessage(noise.NextAvailableOpcode(), (SignUpRequest)(nil))
	OpcodeUsersSignUpResponse     = noise.RegisterMessage(noise.NextAvailableOpcode(), (SignUpResponse)(nil))
	OpcodeUsersUnlockRequest      = noise.RegisterMessage(noise.NextAvailableOpcode(), (UnlockRequest)(nil))
	OpcodeUsersUnlockResponse     = noise.RegisterMessage(noise.NextAvailableOpcode(), (UnlockResponse)(nil))
	OpcodeConsentsConsentRequest  = noise.RegisterMessage(noise.NextAvailableOpcode(), (ConsentRequest)(nil))
	OpcodeConsentsConsentResponse = noise.RegisterMessage(noise.NextAvailableOpcode(), (ConsentResponse)(nil))
)
