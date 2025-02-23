package app

import (
	"encoding/json"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/CosmWasm/wasmd/x/wasm"
	dbm "github.com/cometbft/cometbft-db"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/crypto"
	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/cometbft/cometbft/libs/log"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/server"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"

	// "github.com/cosmos/cosmos-sdk/x/supply"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	// authexported "github.com/cosmos/cosmos-sdk/x/auth/exported"
	// "github.com/cosmos/cosmos-sdk/x/supply"
	// "github.com/CosmWasm/wasmd/x/wasm"
)

var emptyWasmOpts []wasm.Option = nil

// TestApp is a simple wrapper around an App. It exposes internal keepers for use in integration tests.
// This file also contains test helpers. Ideally they would be in separate package.
// Basic Usage:
//
//	Create a test app with NewTestApp, then all keepers and their methods can be accessed for test setup and execution.
//
// Advanced Usage:
//
//	Some tests call for an app to be initialized with some state. This can be achieved through keeper method calls (ie keeper.SetParams(...)).
//	However this leads to a lot of duplicated logic similar to InitGenesis methods.
//	So TestApp.InitializeFromGenesisStates() will call InitGenesis with the default genesis state.
//	and TestApp.InitializeFromGenesisStates(authState, cdpState) will do the same but overwrite the auth and cdp sections of the default genesis state
//	Creating the genesis states can be combersome, but helper methods can make it easier such as NewAuthGenStateFromAccounts below.
type TestApp struct {
	App
}

func NewTestApp() TestApp {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)
	// config.Seal()

	invCheckPeriod := uint(5)
	appOptions := make(simtestutil.AppOptionsMap, 0)
	appOptions[flags.FlagHome] = DefaultNodeHome
	appOptions[server.FlagInvCheckPeriod] = invCheckPeriod

	db := dbm.NewMemDB()
	tApp := NewApp(
		log.NewNopLogger(),
		db,
		nil,
		true,
		wasm.EnableAllProposals,
		appOptions,
		emptyWasmOpts,
	)
	return TestApp{App: *tApp}
}

// nolint
func (tApp TestApp) GetAccountKeeper() authkeeper.AccountKeeper { return tApp.AccountKeeper }
func (tApp TestApp) GetBankKeeper() bankkeeper.Keeper           { return tApp.BankKeeper }

// func (tApp TestApp) GetSupplyKeeper() supply.Keeper             { return tApp.SupplyKeeper }
func (tApp TestApp) GetStakingKeeper() *stakingkeeper.Keeper  { return tApp.StakingKeeper }
func (tApp TestApp) GetSlashingKeeper() slashingkeeper.Keeper { return tApp.SlashingKeeper }
func (tApp TestApp) GetMintKeeper() mintkeeper.Keeper         { return tApp.MintKeeper }
func (tApp TestApp) GetDistrKeeper() distrkeeper.Keeper       { return tApp.DistrKeeper }
func (tApp TestApp) GetGovKeeper() govkeeper.Keeper           { return tApp.GovKeeper }
func (tApp TestApp) GetCrisisKeeper() *crisiskeeper.Keeper    { return tApp.CrisisKeeper }
func (tApp TestApp) GetUpgradeKeeper() *upgradekeeper.Keeper  { return tApp.UpgradeKeeper }
func (tApp TestApp) GetParamsKeeper() paramskeeper.Keeper     { return tApp.ParamsKeeper }

// func (tApp TestApp) GetVVKeeper() validatorvesting.Keeper       { return tApp.vvKeeper }
// func (tApp TestApp) GetPriceFeedKeeper() pricefeedkeeper.Keeper { return tApp.PricefeedKeeper }

// func (tApp TestApp) GetHarvestKeeper() harvest.Keeper           { return tApp.harvestKeeper }
// func (tApp TestApp) GetCommitteeKeeper() committee.Keeper       { return tApp.committeeKeeper }
// func (tApp TestApp) GetIssuanceKeeper() issuance.Keeper         { return tApp.issuanceKeeper }

// InitializeFromGenesisStates calls InitChain on the app using the default genesis state, overwitten with any passed in genesis states
func (tApp TestApp) InitializeFromGenesisStates(genesisStates ...GenesisState) TestApp {
	// Create a default genesis state and overwrite with provided values
	genesisState := NewDefaultGenesisState(tApp.appCodec)
	for _, state := range genesisStates {
		for k, v := range state {
			genesisState[k] = v
		}
	}

	// Initialize the chain
	// stateBytes, err := codec.MarshalJSONIndent(tApp.cdc, genesisState)
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	if err != nil {
		panic(err)
	}
	tApp.InitChain(
		abci.RequestInitChain{
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	tApp.Commit()
	tApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: tApp.LastBlockHeight() + 1}})
	return tApp
}

// InitializeFromGenesisStatesWithTime calls InitChain on the app using the default genesis state, overwitten with any passed in genesis states and genesis Time
func (tApp TestApp) InitializeFromGenesisStatesWithTime(genTime time.Time, genesisStates ...GenesisState) TestApp {
	// Create a default genesis state and overwrite with provided values
	genesisState := NewDefaultGenesisState(tApp.appCodec)
	for _, state := range genesisStates {
		for k, v := range state {
			genesisState[k] = v
		}
	}

	// Initialize the chain
	// stateBytes, err := codec.MarshalJSONIndent(tApp.cdc, genesisState)
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	if err != nil {
		panic(err)
	}
	tApp.InitChain(
		abci.RequestInitChain{
			Time:          genTime,
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
		},
	)
	tApp.Commit()
	tApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: tApp.LastBlockHeight() + 1, Time: genTime}})
	return tApp
}

// InitializeFromGenesisStatesWithTimeAndChainID calls InitChain on the app using the default genesis state, overwitten with any passed in genesis states and genesis Time
func (tApp TestApp) InitializeFromGenesisStatesWithTimeAndChainID(genTime time.Time, chainID string, genesisStates ...GenesisState) TestApp {
	// Create a default genesis state and overwrite with provided values
	genesisState := NewDefaultGenesisState(tApp.appCodec)
	for _, state := range genesisStates {
		for k, v := range state {
			genesisState[k] = v
		}
	}

	// Initialize the chain
	// stateBytes, err := codec.MarshalJSONIndent(tApp.cdc, genesisState)
	stateBytes, err := json.MarshalIndent(genesisState, "", "  ")
	if err != nil {
		panic(err)
	}
	tApp.InitChain(
		abci.RequestInitChain{
			Time:          genTime,
			Validators:    []abci.ValidatorUpdate{},
			AppStateBytes: stateBytes,
			ChainId:       chainID,
		},
	)
	tApp.Commit()
	tApp.BeginBlock(abci.RequestBeginBlock{Header: tmproto.Header{Height: tApp.LastBlockHeight() + 1, Time: genTime}})
	return tApp
}

func (tApp TestApp) CheckBalance(t *testing.T, ctx sdk.Context, owner sdk.AccAddress, expectedCoins sdk.Coins) {
	// acc := tApp.GetAccountKeeper().GetAccount(ctx, owner)
	// require.NotNilf(t, acc, "account with address '%s' doesn't exist", owner)
	require.Equal(t, expectedCoins, tApp.GetBankKeeper().GetAllBalances(ctx, owner))
}

// Create a new auth genesis state from some addresses and coins. The state is returned marshalled into a map.
func NewAuthGenState(tApp TestApp, addresses []sdk.AccAddress, coins []sdk.Coins) GenesisState {
	// Create GenAccounts
	accounts := authtypes.GenesisAccounts{}
	for i := range addresses {
		accounts = append(accounts, authtypes.NewBaseAccountWithAddress(addresses[i]))
	}
	// Create the auth genesis state
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), accounts)
	// Create the bank genesis state
	bankGenesis := banktypes.DefaultGenesisState()
	for i := range addresses {
		bankGenesis.Balances = append(bankGenesis.Balances, banktypes.Balance{Address: addresses[i].String(), Coins: coins[i]})
	}
	// return GenesisState{authtypes.ModuleName: authtypes.ModuleCdc.MustMarshalJSON(authGenesis), banktypes.ModuleName: banktypes.ModuleCdc.MustMarshalJSON(bankGenesis)}
	return GenesisState{authtypes.ModuleName: tApp.appCodec.MustMarshalJSON(authGenesis), banktypes.ModuleName: tApp.appCodec.MustMarshalJSON(bankGenesis)}
}

// Create a new auth genesis state from some addresses and coins. The state is returned marshalled into a map.
func NewAuthGenStateModAcc(tApp TestApp, moduleAccounts []*authtypes.ModuleAccount) GenesisState {
	// Create GenAccounts
	accounts := authtypes.GenesisAccounts{}
	for i := range accounts {
		accounts = append(accounts, moduleAccounts[i])
	}
	// Create the auth genesis state
	authGenesis := authtypes.NewGenesisState(authtypes.DefaultParams(), accounts)

	return GenesisState{authtypes.ModuleName: tApp.appCodec.MustMarshalJSON(authGenesis)}
}

// GeneratePrivKeyAddressPairsFromRand generates (deterministically) a total of n secp256k1 private keys and addresses.
func GeneratePrivKeyAddressPairs(n int) (keys []crypto.PrivKey, addrs []sdk.AccAddress) {
	config := sdk.GetConfig()
	config.SetBech32PrefixForAccount(AccountAddressPrefix, AccountPubKeyPrefix)
	config.SetBech32PrefixForValidator(ValidatorAddressPrefix, ValidatorPubKeyPrefix)
	config.SetBech32PrefixForConsensusNode(ConsNodeAddressPrefix, ConsNodePubKeyPrefix)

	r := rand.New(rand.NewSource(12345)) // make the generation deterministic
	keys = make([]crypto.PrivKey, n)
	addrs = make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		secret := make([]byte, 32)
		_, err := r.Read(secret)
		if err != nil {
			panic("Could not read randomness")
		}
		keys[i] = secp256k1.GenPrivKeySecp256k1(secret)
		addrs[i] = sdk.AccAddress(keys[i].PubKey().Address())
	}
	return
}
