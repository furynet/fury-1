package keeper_test

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/fury-zone/fury/modules/guardian/keeper"
	"github.com/fury-zone/fury/modules/guardian/types"
	"github.com/fury-zone/fury/simapp"
)

var (
	pks = []cryptotypes.PubKey{
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB50"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB51"),
		newPubKey("0B485CFC0EECC619440448436F8FC9DF40566F2369E72400281454CB552AFB52"),
	}
	addrs = []sdk.AccAddress{
		sdk.AccAddress(pks[0].Address()),
		sdk.AccAddress(pks[1].Address()),
		sdk.AccAddress(pks[2].Address()),
	}
)

type KeeperTestSuite struct {
	suite.Suite

	cdc    *codec.LegacyAmino
	ctx    sdk.Context
	keeper keeper.Keeper
	app    *simapp.SimApp
}

func (suite *KeeperTestSuite) SetupTest() {
	app := simapp.Setup(suite.T(), false)

	suite.app = app
	suite.cdc = app.LegacyAmino()
	suite.ctx = app.BaseApp.NewContext(false, tmproto.Header{})
	suite.keeper = app.GuardianKeeper
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestAddSuper() {
	super := types.NewSuper("test", types.Genesis, addrs[0], addrs[1])

	suite.keeper.AddSuper(suite.ctx, super)
	addedSuper, found := suite.keeper.GetSuper(suite.ctx, addrs[0])
	suite.True(found)
	suite.True(super.Equal(addedSuper))

	var supers []types.Super
	suite.keeper.IterateSupers(
		suite.ctx,
		func(super types.Super) bool {
			supers = append(supers, super)
			return false
		},
	)

	suite.Equal(1, len(supers))
	suite.Contains(supers, super)
}

func (suite *KeeperTestSuite) TestDeleteSuper() {
	super := types.NewSuper("test", types.Genesis, addrs[0], addrs[1])

	suite.keeper.AddSuper(suite.ctx, super)
	addedSuper, found := suite.keeper.GetSuper(suite.ctx, addrs[0])
	suite.True(found)
	suite.True(super.Equal(addedSuper))

	address, _ := sdk.AccAddressFromBech32(super.Address)
	suite.keeper.DeleteSuper(suite.ctx, address)

	_, found = suite.keeper.GetSuper(suite.ctx, addrs[0])
	suite.False(found)
}

func (suite *KeeperTestSuite) TestQuerySupers() {
	super := types.NewSuper("test", types.Genesis, addrs[0], addrs[1])
	suite.keeper.AddSuper(suite.ctx, super)

	var supers []types.Super
	querier := keeper.NewQuerier(suite.keeper, suite.cdc)
	res, sdkErr := querier(suite.ctx, []string{types.QuerySupers}, abci.RequestQuery{})
	suite.NoError(sdkErr)

	err := suite.cdc.UnmarshalJSON(res, &supers)
	suite.NoError(err)
	suite.Len(supers, 1)
	suite.Contains(supers, super)
}

func newPubKey(pk string) (res cryptotypes.PubKey) {
	pkBytes, err := hex.DecodeString(pk)
	if err != nil {
		panic(err)
	}

	pubkey := &ed25519.PubKey{Key: pkBytes}

	return pubkey
}
