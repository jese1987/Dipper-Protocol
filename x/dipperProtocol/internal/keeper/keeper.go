package keeper

import (
	"github.com/Dipper-Protocol/x/dipperProtocol/internal/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

// Keeper maintains the link to storage and exposes getter/setter methods for the various parts of the state machine
type Keeper struct {
	CoinKeeper bank.Keeper

	storeKey sdk.StoreKey // Unexposed key to access store from sdk.Context

	cdc *codec.Codec // The wire codec for binary encoding/decoding.
}

// NewKeeper creates new instances of the dipperProtocol Keeper
func NewKeeper(coinKeeper bank.Keeper, storeKey sdk.StoreKey, cdc *codec.Codec) Keeper {
	return Keeper{
		CoinKeeper: coinKeeper,
		storeKey:   storeKey,
		cdc:        cdc,
	}
}

// Gets the entire Whois metadata struct for a name
func (k Keeper) GetWhois(ctx sdk.Context, name string) types.Whois {
	store := ctx.KVStore(k.storeKey)
	if !k.IsNamePresent(ctx, name) {
		return types.NewWhois()
	}
	bz := store.Get([]byte(name))
	var whois types.Whois
	k.cdc.MustUnmarshalBinaryBare(bz, &whois)
	return whois
}

// Sets the entire Whois metadata struct for a name
func (k Keeper) SetWhois(ctx sdk.Context, name string, whois types.Whois) {
	if whois.Owner.Empty() {
		return
	}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(whois))
}

// Deletes the entire Whois metadata struct for a name
func (k Keeper) DeleteWhois(ctx sdk.Context, name string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete([]byte(name))
}

// ResolveName - returns the string that the name resolves to
func (k Keeper) ResolveName(ctx sdk.Context, name string) string {
	return k.GetWhois(ctx, name).Value
}

// SetName - sets the value string that a name resolves to
func (k Keeper) SetName(ctx sdk.Context, name string, value string) {
	whois := k.GetWhois(ctx, name)
	whois.Value = value
	k.SetWhois(ctx, name, whois)
}

// HasOwner - returns whether or not the name already has an owner
func (k Keeper) HasOwner(ctx sdk.Context, name string) bool {
	return !k.GetWhois(ctx, name).Owner.Empty()
}

// GetOwner - get the current owner of a name
func (k Keeper) GetOwner(ctx sdk.Context, name string) sdk.AccAddress {
	return k.GetWhois(ctx, name).Owner
}

// SetOwner - sets the current owner of a name
func (k Keeper) SetOwner(ctx sdk.Context, name string, owner sdk.AccAddress) {
	whois := k.GetWhois(ctx, name)
	whois.Owner = owner
	k.SetWhois(ctx, name, whois)
}

// GetPrice - gets the current price of a name
func (k Keeper) GetPrice(ctx sdk.Context, name string) sdk.Coins {
	return k.GetWhois(ctx, name).Price
}

// SetPrice - sets the current price of a name
func (k Keeper) SetPrice(ctx sdk.Context, name string, price sdk.Coins) {
	whois := k.GetWhois(ctx, name)
	whois.Price = price
	k.SetWhois(ctx, name, whois)
}

// Get an iterator over all names in which the keys are the names and the values are the whois
func (k Keeper) GetNamesIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, nil)
}

// Check if the name is present in the store or not
func (k Keeper) IsNamePresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

//Dipper Bank

func (k Keeper) GetBillBank(ctx sdk.Context) types.BillBank {
	store := ctx.KVStore(k.storeKey)
	if !k.IsObjectPresent(ctx, types.DipperBank){
		return *types.NewBillBank()
	}
	bz := store.Get([]byte(types.DipperBank))
	var billBank = types.NewBillBank()
	k.cdc.MustUnmarshalBinaryBare(bz, &billBank)
	return *billBank
}

func (k Keeper) SetBillBank(ctx sdk.Context, bb types.BillBank) {
	//if len(oracle.TokensPrice) == 0 {
	//	return
	//}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(types.DipperBank), k.cdc.MustMarshalBinaryBare(bb))
}

//NetValueOf
func (k Keeper) GetNetValueOf(ctx sdk.Context, user sdk.AccAddress) float64 {
	return k.GetBillBank(ctx).NetValueOf(user)
}

//Borrow methods
func (k Keeper)GetBorrowBalanceOf(ctx sdk.Context, symbol string, user sdk.AccAddress) float64 {
	return k.GetBillBank(ctx).BorrowBalanceOf(symbol, user)
}

func (k Keeper)GetBorrowValueOf(ctx sdk.Context, symbol string, user sdk.AccAddress) float64 {
	return k.GetBillBank(ctx).BorrowValueOf(symbol, user)
}

func (k Keeper)GetBorrowValueEstimate(ctx sdk.Context, amount float64, symbol string) float64{
	return k.GetBillBank(ctx).BorrowValueEstimate(amount, symbol)
}

func (k Keeper)BankBorrow(ctx sdk.Context, amount float64, symbol string, user sdk.AccAddress) {
	bank := k.GetBillBank(ctx)
	bank.Borrow(amount, symbol, user)
	//if err != nil {
	//	return
	//}
	k.SetBillBank(ctx, bank)
}

func (k Keeper)BankRepay(ctx sdk.Context, amount float64, symbol string, user sdk.AccAddress) {
	bank := k.GetBillBank(ctx)
	bank.Repay(amount, symbol, user)
	k.SetBillBank(ctx, bank)
}


//Supply methods
func (k Keeper)GetSupplyBalanceOf(ctx sdk.Context, symbol string, user sdk.AccAddress) float64 {
	return k.GetBillBank(ctx).SupplyBalanceOf(symbol, user)
}

func (k Keeper)GetSupplyValueOf(ctx sdk.Context, symbol string, user sdk.AccAddress) float64 {
	return k.GetBillBank(ctx).SupplyValueOf(symbol, user)
}

func (k Keeper)BankDeposit(ctx sdk.Context, amount float64, symbol string, user sdk.AccAddress) {
	bank := k.GetBillBank(ctx)
	bank.Deposit(amount, symbol, user)
	k.SetBillBank(ctx, bank)
}

func (k Keeper)BankWithdraw(ctx sdk.Context, amount float64, symbol string, user sdk.AccAddress) {
	bank := k.GetBillBank(ctx)
	bank.Withdraw(amount, symbol, user)
	k.SetBillBank(ctx, bank)
}

//Orcale methods
// Gets the entire Whois metadata struct for a name
func (k Keeper) GetBankOracle(ctx sdk.Context, name string) types.Oracle {
	store := ctx.KVStore(k.storeKey)
	if !k.IsObjectPresent(ctx, name) {
		return *types.NewOracle()
	}
	bz := store.Get([]byte(name))
	var oracle types.Oracle
	k.cdc.MustUnmarshalBinaryBare(bz, &oracle)
	return oracle
}

func (k Keeper) SetBankOracle(ctx sdk.Context, name string, oracle types.Oracle) {
	//if len(oracle.TokensPrice) == 0 {
	//	return
	//}
	store := ctx.KVStore(k.storeKey)
	store.Set([]byte(name), k.cdc.MustMarshalBinaryBare(oracle))
}

func (k Keeper) IsObjectPresent(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has([]byte(name))
}

func (k Keeper)GetOraclePrice(ctx sdk.Context, name string, symbol string) float64 {
	oracle := k.GetBankOracle(ctx, name)
	return oracle.GetPrice(symbol)
}

func (k Keeper)SetOraclePrice(ctx sdk.Context, name string, symbol string, price float64) {
	oracle := k.GetBankOracle(ctx, name)
	oracle.SetPrice(symbol, price)
	k.SetBankOracle(ctx, name, oracle)
}