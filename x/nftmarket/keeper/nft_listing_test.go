package keeper_test

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	nfttypes "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/tendermint/tendermint/crypto/ed25519"

	ununifitypes "github.com/UnUniFi/chain/types"
	"github.com/UnUniFi/chain/x/nftmarket/types"
)

// test basic functions of nft listing
func (suite *KeeperTestSuite) TestNftListingBasics() {
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	owner2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	now := time.Now().UTC()
	future := now.Add(time.Second)
	listings := []types.NftListing{
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "1",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_LISTING,
			BidToken:           "uguu",
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   time.Time{},
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "2",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_BIDDING,
			BidToken:           "uguu",
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   time.Time{},
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "3",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_END_LISTING,
			BidToken:           "uguu",
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          now,
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "1",
				NftId:   "4",
			},
			Owner:              owner.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_SELLING_DECISION,
			BidToken:           "uguu",
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: time.Time{},
			AutoRelistedCount:  0,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "2",
				NftId:   "1",
			},
			Owner:              owner2.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_SUCCESSFUL_BID,
			BidToken:           "uguu",
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: now,
			AutoRelistedCount:  0,
		},
		{
			NftId: types.NftIdentifier{
				ClassId: "2",
				NftId:   "2",
			},
			Owner:              owner2.String(),
			ListingType:        types.ListingType_DIRECT_ASSET_BORROW,
			State:              types.ListingState_LIQUIDATION,
			BidToken:           "uguu",
			MinBid:             sdk.OneInt(),
			BidActiveRank:      1,
			StartedAt:          time.Time{},
			EndAt:              now,
			FullPaymentEndAt:   now,
			SuccessfulBidEndAt: now,
			AutoRelistedCount:  0,
		},
	}

	for _, listing := range listings {
		suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)
	}

	for _, listing := range listings {
		gotListing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, listing.IdBytes())
		suite.Require().NoError(err)
		suite.Require().Equal(listing, gotListing)
	}

	// check all listings
	allListings := suite.app.NftmarketKeeper.GetAllNftListings(suite.ctx)
	suite.Require().Len(allListings, len(listings))

	// check listing by owner
	listingsByOwner := suite.app.NftmarketKeeper.GetListingsByOwner(suite.ctx, owner)
	suite.Require().Len(listingsByOwner, 4)

	// check active listings (bidding or listing status ending now)
	activeNftListings := suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, now)
	suite.Require().Len(activeNftListings, 0)
	// check active listings (bidding or listing status ending future time)
	activeNftListings = suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(activeNftListings, 2)

	// full payment listings (sell decision or end listing status ending now)
	fullPaymentNftListingsEnding := suite.app.NftmarketKeeper.GetFullPaymentNftListingsEndingAt(suite.ctx, now)
	suite.Require().Len(fullPaymentNftListingsEnding, 0)
	// full payment listings (sell decision or end listing status ending future)
	fullPaymentNftListingsEnding = suite.app.NftmarketKeeper.GetFullPaymentNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(fullPaymentNftListingsEnding, 2)

	// successful listing endings (ending now)
	successfulNftListingsEnding := suite.app.NftmarketKeeper.GetSuccessfulBidNftListingsEndingAt(suite.ctx, now)
	suite.Require().Len(successfulNftListingsEnding, 0)
	// successful listing endings (ending future)
	successfulNftListingsEnding = suite.app.NftmarketKeeper.GetSuccessfulBidNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(successfulNftListingsEnding, 1)

	// delete all the listings
	for _, listing := range listings {
		suite.app.NftmarketKeeper.DeleteNftListing(suite.ctx, listing)
	}

	// check queries after deleting all
	allListings = suite.app.NftmarketKeeper.GetAllNftListings(suite.ctx)
	suite.Require().Len(allListings, 0)

	// listings by owner
	listingsByOwner = suite.app.NftmarketKeeper.GetListingsByOwner(suite.ctx, owner)
	suite.Require().Len(listingsByOwner, 0)

	// queries for active listings
	activeNftListings = suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(activeNftListings, 0)
	// queries for full payment queue
	fullPaymentNftListingsEnding = suite.app.NftmarketKeeper.GetFullPaymentNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(fullPaymentNftListingsEnding, 0)
	// queries for successful ending queue
	successfulNftListingsEnding = suite.app.NftmarketKeeper.GetSuccessfulBidNftListingsEndingAt(suite.ctx, future)
	suite.Require().Len(successfulNftListingsEnding, 0)
}

func (suite *KeeperTestSuite) TestListNft() {
	acc1 := suite.addrs[0]
	acc2 := suite.addrs[1]
	// suite.addrs[0] = acc1
	// suite.addrs[1] = acc2

	tests := []struct {
		testCase   string
		classId    string
		nftId      string
		nftOwner   sdk.AccAddress
		lister     sdk.AccAddress
		bidToken   string
		activeRank uint64
		mintBefore bool
		listBefore bool
		expectPass bool
	}{
		{
			testCase:   "not existing nft",
			classId:    "class1",
			nftId:      "nft1",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 1,
			mintBefore: false,
			listBefore: false,
			expectPass: false,
		},
		{
			testCase:   "already listed",
			classId:    "class2",
			nftId:      "nft2",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 1,
			mintBefore: true,
			listBefore: true,
			expectPass: false,
		},
		{
			testCase:   "not owned nft",
			classId:    "class3",
			nftId:      "nft3",
			nftOwner:   acc1,
			lister:     acc2,
			bidToken:   "uguu",
			activeRank: 1,
			mintBefore: true,
			listBefore: false,
			expectPass: false,
		},
		{
			testCase:   "unsupported bid token",
			classId:    "class4",
			nftId:      "nft4",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "xxxx",
			activeRank: 1,
			mintBefore: true,
			listBefore: false,
			expectPass: false,
		},
		{
			testCase:   "successful listing with default active rank",
			classId:    "class5",
			nftId:      "nft5",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 0,
			mintBefore: true,
			listBefore: false,
			expectPass: true,
		},
		{
			testCase:   "successful listing with non-default active rank",
			classId:    "class6",
			nftId:      "nft6",
			nftOwner:   acc1,
			lister:     acc1,
			bidToken:   "uguu",
			activeRank: 100,
			mintBefore: true,
			listBefore: false,
			expectPass: true,
		},
		{
			testCase:   "successful anther owner",
			classId:    "class7",
			nftId:      "nft7",
			nftOwner:   acc2,
			lister:     acc2,
			bidToken:   "uguu",
			activeRank: 0,
			mintBefore: true,
			listBefore: false,
			expectPass: true,
		},
	}

	for _, tc := range tests {
		if tc.mintBefore {
			suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
				Id:          tc.classId,
				Name:        tc.classId,
				Symbol:      tc.classId,
				Description: tc.classId,
				Uri:         tc.classId,
			})
			err := suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
				ClassId: tc.classId,
				Id:      tc.nftId,
				Uri:     tc.nftId,
				UriHash: tc.nftId,
			}, tc.nftOwner)
			suite.Require().NoError(err)
		}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.lister),
				NftId:         types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      tc.bidToken,
				MinBid:        sdk.ZeroInt(),
				BidActiveRank: tc.activeRank,
			})
			suite.Require().NoError(err)
		}
		err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:        ununifitypes.StringAccAddress(tc.lister),
			NftId:         types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId},
			ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:      tc.bidToken,
			MinBid:        sdk.ZeroInt(),
			BidActiveRank: tc.activeRank,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
			// get listing
			listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, (types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}).IdBytes())
			suite.Require().NoError(err)

			// check ownership is transferred
			moduleAddr := suite.app.AccountKeeper.GetModuleAddress(types.ModuleName)
			nftOwner := suite.app.NFTKeeper.GetOwner(suite.ctx, tc.classId, tc.nftId)
			suite.Require().Equal(nftOwner.String(), moduleAddr.String())

			// check bid active rank is set to default if zero
			if tc.activeRank == 0 {
				suite.Require().Equal(params.DefaultBidActiveRank, listing.BidActiveRank)
			}

			// check startedAt is set as current time
			suite.Require().Equal(suite.ctx.BlockTime(), listing.StartedAt)

			// check endAt is set from initial listing duration
			suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingPeriodInitial)), listing.EndAt)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestCancelNftListing() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase     string
		classId      string
		nftId        string
		nftOwner     sdk.AccAddress
		canceller    sdk.AccAddress
		cancelAfter  time.Duration
		numBids      int
		listBefore   bool
		endedListing bool
		expectPass   bool
	}{
		{
			testCase:     "not existing listing",
			classId:      "class1",
			nftId:        "nft1",
			nftOwner:     acc1,
			canceller:    acc1,
			cancelAfter:  time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:      0,
			listBefore:   false,
			endedListing: false,
			expectPass:   false,
		},
		{
			testCase:     "not owned nft listing",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			canceller:    acc2,
			cancelAfter:  time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:      0,
			listBefore:   true,
			endedListing: false,
			expectPass:   false,
		},
		{
			testCase:     "cancel time not pass",
			classId:      "class3",
			nftId:        "nft3",
			nftOwner:     acc1,
			canceller:    acc1,
			cancelAfter:  0,
			numBids:      0,
			listBefore:   true,
			endedListing: false,
			expectPass:   false,
		},
		{
			testCase:     "already ended listing",
			classId:      "class4",
			nftId:        "nft4",
			nftOwner:     acc1,
			canceller:    acc1,
			cancelAfter:  time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:      0,
			listBefore:   true,
			endedListing: true,
			expectPass:   false,
		},
		{
			testCase:     "successful cancel without cancel fee",
			classId:      "class5",
			nftId:        "nft5",
			nftOwner:     acc1,
			canceller:    acc1,
			cancelAfter:  time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:      0,
			listBefore:   true,
			endedListing: false,
			expectPass:   true,
		},
		{
			testCase:     "successful cancel with cancel fee",
			classId:      "class6",
			nftId:        "nft6",
			nftOwner:     acc1,
			canceller:    acc1,
			cancelAfter:  time.Second * time.Duration(params.NftListingCancelRequiredSeconds+1),
			numBids:      0,
			listBefore:   true,
			endedListing: false,
			expectPass:   true,
		},
	}

	for _, tc := range tests {
		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err := suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:         nftIdentifier,
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      "uguu",
				MinBid:        sdk.ZeroInt(),
				BidActiveRank: 2,
			})
			suite.Require().NoError(err)
		}

		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: false,
			})
			suite.Require().NoError(err)
		}

		if tc.endedListing {
			err := suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)
		}

		oldCancellerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.canceller, "uguu")
		suite.ctx = suite.ctx.WithBlockTime(suite.ctx.BlockTime().Add(tc.cancelAfter))
		err = suite.app.NftmarketKeeper.CancelNftListing(suite.ctx, &types.MsgCancelNftListing{
			Sender: ununifitypes.StringAccAddress(tc.canceller),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// check all bids are closed and returned
			nftBids := suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Len(nftBids, 0)

			// check cancel fee is reduced from listing owner
			if tc.numBids > 0 {
				cancellerNewBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.canceller, "uguu")
				suite.Require().True(cancellerNewBalance.Amount.LT(oldCancellerBalance.Amount))
			}

			// check nft ownership is returned back to owner
			owner := suite.app.NFTKeeper.GetOwner(suite.ctx, tc.classId, tc.nftId)
			suite.Require().Equal(owner, tc.nftOwner)

			// check nft listing is deleted
			_, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Error(err)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestExpandListingPeriod() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase     string
		classId      string
		nftId        string
		nftOwner     sdk.AccAddress
		executor     sdk.AccAddress
		numBids      int
		listBefore   bool
		endedListing bool
		expectPass   bool
	}{
		{
			testCase:     "not existing listing",
			classId:      "class1",
			nftId:        "nft1",
			nftOwner:     acc1,
			executor:     acc1,
			numBids:      0,
			listBefore:   false,
			endedListing: false,
			expectPass:   false,
		},
		{
			testCase:     "not owned nft listing",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			executor:     acc2,
			numBids:      0,
			listBefore:   true,
			endedListing: false,
			expectPass:   false,
		},
		{
			testCase:     "already ended nft listing",
			classId:      "class3",
			nftId:        "nft3",
			nftOwner:     acc1,
			executor:     acc2,
			numBids:      1,
			listBefore:   true,
			endedListing: true,
			expectPass:   false,
		},
		{
			testCase:     "successful nft listing period extend",
			classId:      "class4",
			nftId:        "nft4",
			nftOwner:     acc1,
			executor:     acc1,
			numBids:      1,
			listBefore:   true,
			endedListing: false,
			expectPass:   true,
		},
	}

	for _, tc := range tests {

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{coin})
		suite.NoError(err)

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:         nftIdentifier,
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      "uguu",
				MinBid:        sdk.ZeroInt(),
				BidActiveRank: 2,
			})
			suite.Require().NoError(err)
		}

		lastBidder := sdk.AccAddress{}
		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
			lastBidder = bidder

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: false,
			})
			suite.Require().NoError(err)
		}

		if tc.endedListing {
			err := suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)
		}

		oldListing, _ := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())

		oldExecutorBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.executor, "uguu")
		oldLastBidderBalance := suite.app.BankKeeper.GetBalance(suite.ctx, lastBidder, "uguu")
		err = suite.app.NftmarketKeeper.ExpandListingPeriod(suite.ctx, &types.MsgExpandListingPeriod{
			Sender: ununifitypes.StringAccAddress(tc.executor),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)

			// check fee is paid
			newExecutorBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.executor, "uguu")
			suite.Require().True(newExecutorBalance.Amount.LT(oldExecutorBalance.Amount))

			// check winner bidders get extend fee
			newLastBidderBalance := suite.app.BankKeeper.GetBalance(suite.ctx, lastBidder, "uguu")
			suite.Require().True(newLastBidderBalance.Amount.GT(oldLastBidderBalance.Amount))

			// check listing endAt is extended
			listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().NoError(err)
			suite.Require().True(oldListing.EndAt.Before(listing.EndAt))
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestSellingDecision() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase      string
		classId       string
		nftId         string
		nftOwner      sdk.AccAddress
		executor      sdk.AccAddress
		numBids       int
		enoughAutoPay bool
		autoPayment   bool
		listBefore    bool
		endedListing  bool
		expectPass    bool
	}{
		{
			testCase:      "not existing listing",
			classId:       "class1",
			nftId:         "nft1",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    false,
			endedListing:  false,
			expectPass:    false,
		},
		{
			testCase:      "not owned nft listing",
			classId:       "class2",
			nftId:         "nft2",
			nftOwner:      acc1,
			executor:      acc2,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  false,
			expectPass:    false,
		},
		{
			testCase:      "already ended nft listing",
			classId:       "class3",
			nftId:         "nft3",
			nftOwner:      acc1,
			executor:      acc2,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  true,
			expectPass:    false,
		},
		{
			testCase:      "successful nft selling decision with automatic payment",
			classId:       "class4",
			nftId:         "nft4",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   true,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
		{
			testCase:      "successful nft selling decision with automatic payment enabled with not enough balance",
			classId:       "class5",
			nftId:         "nft5",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: false,
			autoPayment:   true,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
		{
			testCase:      "successful nft selling decision without automatic payment",
			classId:       "class6",
			nftId:         "nft6",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{coin})
		suite.NoError(err)

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:         nftIdentifier,
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      "uguu",
				MinBid:        sdk.ZeroInt(),
				BidActiveRank: 2,
			})
			suite.Require().NoError(err)
		}

		lastBidder := sdk.AccAddress{}
		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
			lastBidder = bidder

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			mintCoin := coin
			if !tc.enoughAutoPay {
				mintCoin = sdk.NewInt64Coin("uguu", int64(1000000*(i+1)/2))
			}
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{mintCoin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{mintCoin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: tc.autoPayment,
			})
			suite.Require().NoError(err)
		}

		if tc.endedListing {
			err := suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)
		}

		err = suite.app.NftmarketKeeper.SellingDecision(suite.ctx, &types.MsgSellingDecision{
			Sender: ununifitypes.StringAccAddress(tc.executor),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)
			if tc.autoPayment {
				bid, err := suite.app.NftmarketKeeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), lastBidder)
				suite.Require().NoError(err)
				if tc.enoughAutoPay {
					// check automatic payment execution when user has enough balance
					suite.Require().Equal(bid.PaidAmount, bid.Amount.Amount)
				} else {
					// check automatic payment when the user does not have enough balance
					suite.Require().NotEqual(bid.PaidAmount, bid.Amount.Amount)
				}
			}

			// check full payment end time update
			listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().NoError(err)
			suite.Require().Equal(listing.State, types.ListingState_SELLING_DECISION)
			suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingFullPaymentPeriod)), listing.FullPaymentEndAt)
		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestEndNftListing() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	acc2 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase      string
		classId       string
		nftId         string
		nftOwner      sdk.AccAddress
		executor      sdk.AccAddress
		numBids       int
		enoughAutoPay bool
		autoPayment   bool
		listBefore    bool
		endedListing  bool
		expectPass    bool
	}{
		{
			testCase:      "not existing listing",
			classId:       "class1",
			nftId:         "nft1",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    false,
			endedListing:  false,
			expectPass:    false,
		},
		{
			testCase:      "not owned nft listing",
			classId:       "class2",
			nftId:         "nft2",
			nftOwner:      acc1,
			executor:      acc2,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  false,
			expectPass:    false,
		},
		{
			testCase:      "already ended nft listing",
			classId:       "class3",
			nftId:         "nft3",
			nftOwner:      acc1,
			executor:      acc2,
			numBids:       1,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  true,
			expectPass:    false,
		},
		{
			testCase:      "successful nft listing ending when no bid",
			classId:       "class4",
			nftId:         "nft4",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       0,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
		{
			testCase:      "successful nft listing ending with automatic payment enabled with not enough balance",
			classId:       "class5",
			nftId:         "nft5",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       1,
			enoughAutoPay: false,
			autoPayment:   true,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
		{
			testCase:      "successful nft listing ending without automatic payment",
			classId:       "class6",
			nftId:         "nft6",
			nftOwner:      acc1,
			executor:      acc1,
			numBids:       4,
			enoughAutoPay: true,
			autoPayment:   false,
			listBefore:    true,
			endedListing:  false,
			expectPass:    true,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.executor, sdk.Coins{coin})
		suite.NoError(err)

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		if tc.listBefore {
			err := suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
				Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:         nftIdentifier,
				ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
				BidToken:      "uguu",
				MinBid:        sdk.ZeroInt(),
				BidActiveRank: 2,
			})
			suite.Require().NoError(err)
		}

		lastBidder := sdk.AccAddress{}
		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
			lastBidder = bidder

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			mintCoin := coin
			if !tc.enoughAutoPay {
				mintCoin = sdk.NewInt64Coin("uguu", int64(1000000*(i+1)/2))
			}
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{mintCoin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{mintCoin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: tc.autoPayment,
			})
			suite.Require().NoError(err)
		}

		if tc.endedListing {
			err := suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
				Sender: ununifitypes.StringAccAddress(tc.nftOwner),
				NftId:  nftIdentifier,
			})
			suite.Require().NoError(err)
		}

		err = suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
			Sender: ununifitypes.StringAccAddress(tc.executor),
			NftId:  nftIdentifier,
		})

		if tc.expectPass {
			suite.Require().NoError(err)
			if tc.autoPayment {
				bid, err := suite.app.NftmarketKeeper.GetBid(suite.ctx, nftIdentifier.IdBytes(), lastBidder)
				suite.Require().NoError(err)
				if tc.enoughAutoPay {
					// check automatic payment execution when user has enough balance
					suite.Require().Equal(bid.PaidAmount, bid.Amount.Amount)
				} else {
					// check automatic payment when the user does not have enough balance
					suite.Require().NotEqual(bid.PaidAmount, bid.Amount.Amount)
				}
			}

			if tc.numBids == 0 {
				// successful end when there's no bid - delete listing directly and transfer nft back to owner
				_, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().Error(err)
			} else {
				// check full payment end time update
				listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_END_LISTING)
				suite.Require().Equal(suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingFullPaymentPeriod)), listing.FullPaymentEndAt)

				// check non-active bids are cancelled automatically
				bids := suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().True(len(bids) <= int(listing.BidActiveRank))
			}

		} else {
			suite.Require().Error(err)
		}
	}
}

func (suite *KeeperTestSuite) TestProcessEndingNftListings() {
	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase            string
		classId             string
		nftId               string
		nftOwner            sdk.AccAddress
		numBids             int
		relistedCount       uint64
		expectedToEnd       bool
		expectedToBeRemoved bool
	}{
		{
			testCase:            "no bid nft listing extend when relisted count not reached the limit",
			classId:             "class1",
			nftId:               "nft1",
			nftOwner:            acc1,
			numBids:             0,
			relistedCount:       0,
			expectedToEnd:       false,
			expectedToBeRemoved: false,
		},
		{
			testCase:            "no bid nft listing end when relisted count reached",
			classId:             "class2",
			nftId:               "nft2",
			nftOwner:            acc1,
			numBids:             0,
			relistedCount:       params.AutoRelistingCountIfNoBid,
			expectedToEnd:       true,
			expectedToBeRemoved: true,
		},
		{
			testCase:            "bids existing nft listing end when relisted count not reached",
			classId:             "class3",
			nftId:               "nft3",
			nftOwner:            acc1,
			numBids:             1,
			relistedCount:       0,
			expectedToEnd:       true,
			expectedToBeRemoved: false,
		},
		{
			testCase:            "bids existing nft listing end when relisted count reached",
			classId:             "class4",
			nftId:               "nft4",
			nftOwner:            acc1,
			numBids:             1,
			relistedCount:       params.AutoRelistingCountIfNoBid,
			expectedToEnd:       true,
			expectedToBeRemoved: false,
		},
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.nftOwner, sdk.Coins{coin})
		suite.NoError(err)

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
			NftId:         nftIdentifier,
			ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:      "uguu",
			MinBid:        sdk.ZeroInt(),
			BidActiveRank: 2,
		})
		suite.Require().NoError(err)
		listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
		suite.Require().NoError(err)
		listing.AutoRelistedCount = tc.relistedCount
		suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)

		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: true,
			})
			suite.Require().NoError(err)
		}

		suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second * time.Duration(params.NftListingPeriodInitial+1)))
		suite.app.NftmarketKeeper.ProcessEndingNftListings(suite.ctx)

		if tc.expectedToBeRemoved {
			_, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().Error(err)
		} else {
			listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
			suite.Require().NoError(err)

			if tc.expectedToEnd {
				suite.Require().Equal(listing.State, types.ListingState_END_LISTING)
			} else {
				suite.Require().NotEqual(listing.State, types.ListingState_END_LISTING)
			}
		}
	}
}

func (suite *KeeperTestSuite) TestActiveNftListingsEndingAtQueueRemovalOnNftListingEnd() {
	suite.SetupTest()

	classId := "class1"
	nftId := "nf1"
	nftOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()

	suite.ctx = suite.ctx.WithBlockTime(now)
	coin := sdk.NewInt64Coin("uguu", int64(1000000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{coin})
	suite.NoError(err)

	suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
		Id:          classId,
		Name:        classId,
		Symbol:      classId,
		Description: classId,
		Uri:         classId,
	})
	err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     nftId,
		UriHash: nftId,
	}, nftOwner)
	suite.Require().NoError(err)

	nftIdentifier := types.NftIdentifier{ClassId: classId, NftId: nftId}
	err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
		Sender:        ununifitypes.StringAccAddress(nftOwner),
		NftId:         nftIdentifier,
		ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:      "uguu",
		MinBid:        sdk.ZeroInt(),
		BidActiveRank: 2,
	})
	suite.Require().NoError(err)

	listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().NoError(err)
	suite.Require().True(listing.IsActive())

	// check number before end listing
	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
	activeNftListings := suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, now.Add(time.Second*time.Duration(params.NftListingPeriodInitial+1)))
	suite.Require().Len(activeNftListings, 1)

	err = suite.app.NftmarketKeeper.EndNftListing(suite.ctx, &types.MsgEndNftListing{
		Sender: ununifitypes.StringAccAddress(nftOwner),
		NftId:  nftIdentifier,
	})
	suite.Require().NoError(err)

	// check number after end listing
	activeNftListings = suite.app.NftmarketKeeper.GetActiveNftListingsEndingAt(suite.ctx, now.Add(time.Second*time.Duration(params.NftListingPeriodInitial+1)))
	suite.Require().Len(activeNftListings, 0)
}

func (suite *KeeperTestSuite) TestHandleFullPaymentPeriodEndings() {

	acc1 := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)

	tests := []struct {
		testCase     string
		classId      string
		nftId        string
		nftOwner     sdk.AccAddress
		numBids      int
		listingState types.ListingState
		fullPay      bool
	}{
		{
			testCase:     "selling decision listing when highest bid is paid",
			classId:      "class1",
			nftId:        "nft1",
			nftOwner:     acc1,
			numBids:      2,
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      true,
		}, // add successful listing state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status
		{
			testCase:     "selling decision listing when highest bid is not paid and no more bids",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			numBids:      1,
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      false,
		}, // status => ListingState_LISTING
		{
			testCase:     "selling decision listing when highest bid is not paid, and more bids",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			numBids:      2,
			listingState: types.ListingState_SELLING_DECISION,
			fullPay:      false,
		}, // status => ListingState_BIDDING
		{
			testCase:     "ended listing, when fully paid bid exists",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			numBids:      2,
			listingState: types.ListingState_END_LISTING,
			fullPay:      true,
		}, // add successful bid state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status, close all the other bids
		{
			testCase:     "ended listing, when fully paid bid does not exist",
			classId:      "class2",
			nftId:        "nft2",
			nftOwner:     acc1,
			numBids:      2,
			listingState: types.ListingState_END_LISTING,
			fullPay:      true,
		}, // all the bids closed, pay depositCollected, nft listing delete, transfer nft to fully paid bidder
	}

	for _, tc := range tests {
		suite.SetupTest()

		now := time.Now().UTC()
		suite.ctx = suite.ctx.WithBlockTime(now)

		coin := sdk.NewInt64Coin("uguu", int64(1000000000))
		err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
		suite.NoError(err)
		err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, tc.nftOwner, sdk.Coins{coin})
		suite.NoError(err)

		suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
			Id:          tc.classId,
			Name:        tc.classId,
			Symbol:      tc.classId,
			Description: tc.classId,
			Uri:         tc.classId,
		})
		err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
			ClassId: tc.classId,
			Id:      tc.nftId,
			Uri:     tc.nftId,
			UriHash: tc.nftId,
		}, tc.nftOwner)
		suite.Require().NoError(err)

		nftIdentifier := types.NftIdentifier{ClassId: tc.classId, NftId: tc.nftId}
		err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
			Sender:        ununifitypes.StringAccAddress(tc.nftOwner),
			NftId:         nftIdentifier,
			ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
			BidToken:      "uguu",
			MinBid:        sdk.ZeroInt(),
			BidActiveRank: 2,
		})
		suite.Require().NoError(err)
		listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
		suite.Require().NoError(err)

		for i := 0; i < tc.numBids; i++ {
			bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

			// init tokens to addr
			coin := sdk.NewInt64Coin("uguu", int64(1000000*(i+1)))
			err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
			suite.NoError(err)
			err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
			suite.NoError(err)

			err := suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
				Sender:           ununifitypes.StringAccAddress(bidder),
				NftId:            nftIdentifier,
				Amount:           coin,
				AutomaticPayment: true,
			})
			suite.Require().NoError(err)

			if tc.fullPay {
				err := suite.app.NftmarketKeeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
					Sender: ununifitypes.StringAccAddress(bidder),
					NftId:  nftIdentifier,
				})
				suite.Require().NoError(err)
			}
		}

		listing.State = tc.listingState
		suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)

		oldNftOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.nftOwner, "uguu")
		suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second * time.Duration(params.NftListingPeriodInitial+1)))
		suite.app.NftmarketKeeper.HandleFullPaymentsPeriodEndings(suite.ctx)

		switch tc.listingState {
		case types.ListingState_SELLING_DECISION:
			if tc.fullPay {
				// add successful listing state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_SUCCESSFUL_BID)
				suite.Require().Equal(listing.SuccessfulBidEndAt, suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingNftDeliveryPeriod)))
			} else if tc.numBids > 1 {
				// status => ListingState_BIDDING
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_BIDDING)
			} else {
				// status => ListingState_LISTING
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_LISTING)
			}
		case types.ListingState_END_LISTING:
			if tc.fullPay {
				// add successful bid state with SuccessfulBidEndAt field + types.ListingState_SUCCESSFUL_BID status, close all the other bids
				listing, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(listing.State, types.ListingState_SUCCESSFUL_BID)
				suite.Require().Equal(listing.SuccessfulBidEndAt, suite.ctx.BlockTime().Add(time.Second*time.Duration(params.NftListingNftDeliveryPeriod)))
			} else {
				// all the bids closed, pay depositCollected, nft listing delete, transfer nft to fully paid bidder
				_, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().Error(err)

				bids := suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().Len(bids, 0)

				newOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, tc.nftOwner, "uguu")
				suite.Require().True(newOwnerBalance.Amount.GT(oldNftOwnerBalance.Amount))

				nft, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
				suite.Require().NoError(err)
				suite.Require().Equal(nft.Owner, tc.nftOwner)
			}
		}
	}
}

func (suite *KeeperTestSuite) TestDelieverSuccessfulBids() {
	suite.SetupTest()

	classId := "class1"
	nftId := "nf1"
	nftOwner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())
	now := time.Now().UTC()

	suite.ctx = suite.ctx.WithBlockTime(now)
	coin := sdk.NewInt64Coin("uguu", int64(1000000000))
	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, nftOwner, sdk.Coins{coin})
	suite.NoError(err)

	suite.app.NFTKeeper.SaveClass(suite.ctx, nfttypes.Class{
		Id:          classId,
		Name:        classId,
		Symbol:      classId,
		Description: classId,
		Uri:         classId,
	})
	err = suite.app.NFTKeeper.Mint(suite.ctx, nfttypes.NFT{
		ClassId: classId,
		Id:      nftId,
		Uri:     nftId,
		UriHash: nftId,
	}, nftOwner)
	suite.Require().NoError(err)

	nftIdentifier := types.NftIdentifier{ClassId: classId, NftId: nftId}
	err = suite.app.NftmarketKeeper.ListNft(suite.ctx, &types.MsgListNft{
		Sender:        ununifitypes.StringAccAddress(nftOwner),
		NftId:         nftIdentifier,
		ListingType:   types.ListingType_DIRECT_ASSET_BORROW,
		BidToken:      "uguu",
		MinBid:        sdk.ZeroInt(),
		BidActiveRank: 2,
	})
	suite.Require().NoError(err)

	bidder := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	// init tokens to addr
	err = suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{coin})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToAccount(suite.ctx, minttypes.ModuleName, bidder, sdk.Coins{coin})
	suite.NoError(err)

	err = suite.app.NftmarketKeeper.PlaceBid(suite.ctx, &types.MsgPlaceBid{
		Sender:           ununifitypes.StringAccAddress(bidder),
		NftId:            nftIdentifier,
		Amount:           coin,
		AutomaticPayment: true,
	})
	suite.Require().NoError(err)
	err = suite.app.NftmarketKeeper.PayFullBid(suite.ctx, &types.MsgPayFullBid{
		Sender: ununifitypes.StringAccAddress(bidder),
		NftId:  nftIdentifier,
	})
	suite.Require().NoError(err)

	listing, err := suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().NoError(err)
	listing.SuccessfulBidEndAt = now
	listing.State = types.ListingState_SUCCESSFUL_BID
	suite.app.NftmarketKeeper.SetNftListing(suite.ctx, listing)

	suite.ctx = suite.ctx.WithBlockTime(now.Add(time.Second))
	oldNftOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")

	suite.app.NftmarketKeeper.DelieverSuccessfulBids(suite.ctx)

	// check nft transfer
	newNftOwner := suite.app.NFTKeeper.GetOwner(suite.ctx, classId, nftId)
	suite.Require().NoError(err)
	suite.Require().Equal(newNftOwner.String(), bidder.String())

	// check fee paid
	newOwnerBalance := suite.app.BankKeeper.GetBalance(suite.ctx, nftOwner, "uguu")
	suite.Require().True(newOwnerBalance.Amount.GT(oldNftOwnerBalance.Amount))

	// check bid deleted
	bids := suite.app.NftmarketKeeper.GetBidsByNft(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Len(bids, 0)

	// check nft listing deleted
	_, err = suite.app.NftmarketKeeper.GetNftListingByIdBytes(suite.ctx, nftIdentifier.IdBytes())
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestProcessPaymentWithCommissionFee() {
	amount := sdk.NewInt(1000000)
	owner := sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address().Bytes())

	err := suite.app.BankKeeper.MintCoins(suite.ctx, minttypes.ModuleName, sdk.Coins{sdk.NewCoin("uguu", amount)})
	suite.NoError(err)
	err = suite.app.BankKeeper.SendCoinsFromModuleToModule(suite.ctx, minttypes.ModuleName, types.ModuleName, sdk.Coins{sdk.NewCoin("uguu", amount)})
	suite.NoError(err)

	suite.app.NftmarketKeeper.ProcessPaymentWithCommissionFee(suite.ctx, owner, "uguu", amount)

	params := suite.app.NftmarketKeeper.GetParamSet(suite.ctx)
	fee := amount.Mul(sdk.NewInt(int64(params.NftListingCommissionFee))).Quo(sdk.NewInt(100))
	listingPayment := amount.Sub(fee)

	// check fee paid to NftTradingFee
	tradingFeeModuleAcc := suite.app.AccountKeeper.GetModuleAddress(types.NftTradingFee)
	tradingFeeBal := suite.app.BankKeeper.GetBalance(suite.ctx, tradingFeeModuleAcc, "uguu")
	suite.Require().Equal(tradingFeeBal, sdk.NewCoin("uguu", fee))

	// check fee to lister
	ownerBal := suite.app.BankKeeper.GetBalance(suite.ctx, owner, "uguu")
	suite.Require().Equal(ownerBal, sdk.NewCoin("uguu", listingPayment))
}
