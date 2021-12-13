package coincierge_test

import (
	"testing"

	keepertest "github.com/jclyons52/coincierge/testutil/keeper"
	"github.com/jclyons52/coincierge/x/coincierge"
	"github.com/jclyons52/coincierge/x/coincierge/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.CoinciergeKeeper(t)
	coincierge.InitGenesis(ctx, *k, genesisState)
	got := coincierge.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	// this line is used by starport scaffolding # genesis/test/assert
}
