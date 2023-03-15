package keeper_test

import (
	"time"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/derivatives/types"
	pricefeedtypes "github.com/UnUniFi/chain/x/pricefeed/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

func (suite *KeeperTestSuite) AddPoolAssets() []types.PoolParams_Asset {
	assets := []types.PoolParams_Asset{
		{
			Denom:        "uusdc",
			TargetWeight: sdk.NewDec(1),
		},
		{
			Denom:        "uatom",
			TargetWeight: sdk.NewDec(10),
		},
	}

	for _, asset := range assets {
		suite.keeper.AddPoolAsset(suite.ctx, asset)
	}

	return assets
}

func (suite *KeeperTestSuite) TestAddPoolAsset() {
	// owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	assets := suite.AddPoolAssets()

	for _, asset := range assets {
		assetInStore := suite.keeper.GetPoolAssetByDenom(suite.ctx, asset.Denom)
		suite.Require().Equal(asset, assetInStore)
	}

	// Check if the asset was added
	allAssets := suite.keeper.GetPoolAssets(suite.ctx)

	suite.Require().Len(allAssets, len(assets))
}

func (suite *KeeperTestSuite) TestDepositPoolAsset() {
	suite.AddPoolAssets()

	depositors := []sdk.AccAddress{
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
	}

	assets := []sdk.Coin{
		{
			Denom:  "uusdc",
			Amount: sdk.NewInt(100),
		},
		{
			Denom:  "uatom",
			Amount: sdk.NewInt(10),
		},
	}

	for index, asset := range assets {
		suite.keeper.DepositPoolAsset(suite.ctx, depositors[index], asset)
	}

	for _, asset := range assets {
		amount := suite.keeper.GetAssetBalance(suite.ctx, asset.Denom)
		suite.Require().Equal(asset, amount)
	}
}

func (suite *KeeperTestSuite) TestSetPoolMarketCapSnapshot() {
	suite.AddPoolAssets()

	depositors := []sdk.AccAddress{
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
		sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes()),
	}

	height := suite.ctx.BlockHeight()

	assets := []sdk.Coin{
		{
			Denom:  "uusdc",
			Amount: sdk.NewInt(100),
		},
		{
			Denom:  "uatom",
			Amount: sdk.NewInt(10),
		},
	}

	for index, asset := range assets {
		suite.keeper.DepositPoolAsset(suite.ctx, depositors[index], asset)
	}

	marketCap := suite.keeper.GetPoolMarketCap(suite.ctx)

	// TODO: it's not working yet as we didn't add the ticker to price feed
	suite.keeper.SetPoolMarketCapSnapshot(suite.ctx, height, marketCap)

	// Check if the market cap was set
	marketCapInStore := suite.keeper.GetPoolMarketCapSnapshot(suite.ctx, height)

	suite.Require().Equal(marketCap, marketCapInStore)
}

func (suite *KeeperTestSuite) TestIsAssetValid() {
	poolAssets := suite.keeper.GetPoolAssets(suite.ctx)
	suite.Require().Len(poolAssets, 2)

	isValid := suite.keeper.IsAssetValid(suite.ctx, types.PoolParams_Asset{
		Denom:        "uatom",
		TargetWeight: sdk.NewDecWithPrec(5, 1),
	})
	suite.Require().True(isValid)

	isValid = suite.keeper.IsAssetValid(suite.ctx, types.PoolParams_Asset{
		Denom:        "xxxx",
		TargetWeight: sdk.NewDecWithPrec(5, 1),
	})
	suite.Require().False(isValid)
}

func (suite *KeeperTestSuite) TestSetGetAssetBalance() {
	coin := sdk.NewInt64Coin("uatom", 1000000)
	suite.keeper.SetAssetBalance(suite.ctx, coin)

	amount := suite.keeper.GetAssetBalance(suite.ctx, "uatom")
	suite.Require().Equal(amount, coin)
}

func (suite *KeeperTestSuite) TestGetAssetTargetAmount() {
	// get target amount at initial
	targetAmount, err := suite.keeper.GetAssetTargetAmount(suite.ctx, "uatom")
	suite.Require().NoError(err)
	suite.Require().Equal(targetAmount.String(), "0uatom")

	// set price for asset
	_, err = suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(13), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:uusdc")
	suite.Require().NoError(err)

	suite.keeper.AddPoolAsset(suite.ctx, types.PoolParams_Asset{
		Denom:        "uatom",
		TargetWeight: sdk.OneDec(),
	})
	suite.keeper.SetAssetBalance(suite.ctx, sdk.NewInt64Coin("uatom", 1000000))

	// set lp token supply
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewInt64Coin(types.LiquidityProviderTokenDenom, 1000000)})
	suite.Require().NoError(err)

	// get target amount after data set
	targetAmount, err = suite.keeper.GetAssetTargetAmount(suite.ctx, "uatom")
	suite.Require().NoError(err)
	suite.Require().Equal(targetAmount.String(), "1000000uatom")
}

func (suite *KeeperTestSuite) TestIsPriceReady() {
	suite.SetupTest()
	// get the value when nothing is set
	isReady := suite.keeper.IsPriceReady(suite.ctx)
	suite.Require().True(isReady)

	// get value after adding one pool asset
	suite.keeper.AddPoolAsset(suite.ctx, types.PoolParams_Asset{
		Denom:        "uatom",
		TargetWeight: sdk.OneDec(),
	})
	isReady = suite.keeper.IsPriceReady(suite.ctx)
	suite.Require().True(isReady)

	// get value after configuring price
	_, err := suite.app.PricefeedKeeper.SetPrice(suite.ctx, sdk.AccAddress{}, "uatom:uusdc", sdk.NewDec(13), suite.ctx.BlockTime().Add(time.Hour*3))
	suite.Require().NoError(err)
	params := suite.app.PricefeedKeeper.GetParams(suite.ctx)
	params.Markets = []pricefeedtypes.Market{
		{MarketId: "uatom:uusdc", BaseAsset: "uatom", QuoteAsset: "uusdc", Oracles: []ununifitypes.StringAccAddress{}, Active: true},
	}
	suite.app.PricefeedKeeper.SetParams(suite.ctx, params)
	err = suite.app.PricefeedKeeper.SetCurrentPrices(suite.ctx, "uatom:uusdc")
	suite.Require().NoError(err)

	isReady = suite.keeper.IsPriceReady(suite.ctx)
	suite.Require().True(isReady)
}

func (suite *KeeperTestSuite) TestUserDeposits() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	suite.keeper.SetUserDenomDepositAmount(suite.ctx, owner, "uatom", sdk.NewInt(1000))
	suite.keeper.SetUserDenomDepositAmount(suite.ctx, owner, "uusdc", sdk.NewInt(1000))
	userDeposit := suite.keeper.GetUserDenomDepositAmount(suite.ctx, owner, "uatom")
	suite.Require().Equal(userDeposit, sdk.NewInt(1000))
	userDeposit = suite.keeper.GetUserDenomDepositAmount(suite.ctx, owner, "uusdc")
	suite.Require().Equal(userDeposit, sdk.NewInt(1000))
	userDeposits := suite.keeper.GetUserDeposits(suite.ctx, owner)
	suite.Require().Equal(sdk.Coins(userDeposits).String(), "1000uatom,1000uusdc")
}
