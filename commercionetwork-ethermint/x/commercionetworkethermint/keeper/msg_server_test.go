package keeper_test

import (
	"context"
	"testing"

	keepertest "commercionetwork-ethermint/testutil/keeper"
	"commercionetwork-ethermint/x/commercionetworkethermint/keeper"
	"commercionetwork-ethermint/x/commercionetworkethermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CommercionetworkethermintKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
