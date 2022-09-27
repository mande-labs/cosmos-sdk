package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func setMinimumCommission() (sdk.Dec, sdk.Dec, sdk.Dec) {

	rate, _ := sdk.NewDecFromStr("1.000000000000000000")
	maxRate, _ := sdk.NewDecFromStr("1.000000000000000000")
	changeRate, _ := sdk.NewDecFromStr("0.000000000000000000")

	return rate, maxRate, changeRate
}
