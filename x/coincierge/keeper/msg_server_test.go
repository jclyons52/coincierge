package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/jclyons52/coincierge/testutil/keeper"
	"github.com/jclyons52/coincierge/x/coincierge/keeper"
	"github.com/jclyons52/coincierge/x/coincierge/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.CoinciergeKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
