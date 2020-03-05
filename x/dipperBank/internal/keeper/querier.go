package keeper

import (
	"fmt"
	"github.com/Dipper-Protocol/codec"
	"github.com/Dipper-Protocol/x/dipperBank/internal/types"
	"strconv"

	sdk "github.com/Dipper-Protocol/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// query endpoints supported by the dipperBank Querier
const (
	QueryResolve = "resolve"
	QueryWhois   = "whois"
	QueryNames   = "names"

	QueryOraclePrice = "oracleprice"
	QueryNetValueOf = "netvalue"
	QueryBorrowBalanceOf = "borrowbalance"
	QueryBorrowValueOf = "borrowvalue"
	QueryBorrowValueEstimate = "borrowvalueestimate"
	QuerySupplyBalanceOf = "supplybalance"
	QuerySupplyValueOf = "supplyvalue"
)

// NewQuerier is the module level router for state queries
func NewQuerier(keeper Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) (res []byte, err sdk.Error) {
		switch path[0] {
		case QueryOraclePrice:
			return queryOraclePrice(ctx, path[1:], req, keeper)
		case QueryNetValueOf:
			return queryNetValueOf(ctx, path[1:], req, keeper)
		case QueryBorrowBalanceOf:
			return queryBorrowBalanceOf(ctx, path[1:], req, keeper)
		case QueryBorrowValueOf:
			return queryBorrowValueOf(ctx, path[1:], req, keeper)
		case QueryBorrowValueEstimate:
			return queryBorrowValueEstimate(ctx, path[1:], req, keeper)
		case QuerySupplyBalanceOf:
			return querySupplyBalanceOf(ctx, path[1:], req, keeper)
		case QuerySupplyValueOf:
			return querySupplyValueOf(ctx, path[1:], req, keeper)
		default:
			return nil, sdk.ErrUnknownRequest("unknown dipperBank query endpoint")
		}
	}
}


func queryOraclePrice(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	price := keeper.GetOraclePrice(ctx, path[0])
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: fmt.Sprintf("%d", price)})
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}

func queryNetValueOf(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	addr, err := sdk.AccAddressFromHex(path[0])
	if err != nil {
		panic("could not get right addr")
	}
	value := keeper.GetNetValueOf(ctx, addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: fmt.Sprintf("%d", value)})
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}

func queryBorrowBalanceOf(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	addr, err := sdk.AccAddressFromHex(path[1])
	if err != nil {
		panic("could not get right addr")
	}
	value := keeper.GetBorrowBalanceOf(ctx, path[0], addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: fmt.Sprintf("%d", value)})
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}

func queryBorrowValueOf(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	addr, err := sdk.AccAddressFromHex(path[1])
	if err != nil {
		panic("could not get right addr")
	}
	value := keeper.GetBorrowValueOf(ctx, path[0], addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: fmt.Sprintf("%d", value)})
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}

func queryBorrowValueEstimate(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	amount, err := strconv.ParseInt(path[0], 10, 64)
	if err != nil {
		panic("could not get estimate num")
	}
	value := keeper.GetBorrowValueEstimate(ctx, amount, path[1])
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: fmt.Sprintf("%d", value)})
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}

func querySupplyBalanceOf(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	addr, err := sdk.AccAddressFromHex(path[1])
	if err != nil {
		panic("could not get right addr")
	}
	value := keeper.GetSupplyBalanceOf(ctx, path[0], addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: fmt.Sprintf("%d", value)})
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}

func querySupplyValueOf(ctx sdk.Context, path []string, req abci.RequestQuery, keeper Keeper) ([]byte, sdk.Error) {
	addr, err := sdk.AccAddressFromHex(path[1])
	if err != nil {
		panic("could not get right addr")
	}
	value := keeper.GetSupplyValueOf(ctx, path[0], addr)
	res, err := codec.MarshalJSONIndent(keeper.cdc, types.QueryResResolve{Value: fmt.Sprintf("%d", value)})
	if err != nil {
		panic("could not marshal result to JSON")
	}
	return res, nil
}





