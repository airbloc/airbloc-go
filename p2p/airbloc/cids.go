package airbloc

const (
	CIDReEncryption = 0x20
	CIDIdentity     = 0x21
	CIDDataSync     = 0x22
	CIDDAC          = 0x23
	CIDTest         = 0xFF
)

var cids = []uint64{
	CIDReEncryption,
	CIDIdentity,
	CIDDataSync,
	CIDDAC,
	CIDTest,
}
