package types

const (
	// ModuleName defines the module name
	ModuleName = "incentive"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_capability"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	ParamsKey = "Params-value-"

	JPYXMintingRewardDenom = "ujsmn"
)

var (
	JPYXMintingClaimKeyPrefix                     = []byte{0x01} // prefix for keys that store JPYX minting claims
	JPYXMintingRewardFactorKeyPrefix              = []byte{0x02} // prefix for key that stores JPYX minting reward factors
	PreviousJPYXMintingRewardAccrualTimeKeyPrefix = []byte{0x03} // prefix for key that stores the blocktime
)
