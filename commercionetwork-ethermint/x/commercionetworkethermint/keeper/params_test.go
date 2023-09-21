package keeper_test

import (
	"testing"

	testkeeper "commercionetwork-ethermint/testutil/keeper"
	"commercionetwork-ethermint/x/commercionetworkethermint/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.CommercionetworkethermintKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
