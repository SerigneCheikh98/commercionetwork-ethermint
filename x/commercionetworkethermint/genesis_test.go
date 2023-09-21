package commercionetworkethermint_test

import (
	"testing"

	keepertest "commercionetwork-ethermint/testutil/keeper"
	"commercionetwork-ethermint/testutil/nullify"
	"commercionetwork-ethermint/x/commercionetworkethermint"
	"commercionetwork-ethermint/x/commercionetworkethermint/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CommercionetworkethermintKeeper(t)
	commercionetworkethermint.InitGenesis(ctx, *k, genesisState)
	got := commercionetworkethermint.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
