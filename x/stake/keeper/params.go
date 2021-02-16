package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/public-awesome/stargaze/x/stake/types"
)

// ParamKeyTable for staking module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.KeyTable{}
}

// GetParams returns the total set of curating parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	return types.Params{}
}

// SetParams sets the curating parameters to the param space.
func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {}
