package types

import (
	"github.com/Dipper-Protocol/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc = codec.New()

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgBankBorrow{}, "dipperProtocol/BankBorrow", nil)
	cdc.RegisterConcrete(MsgBankRepay{}, "dipperProtocol/BankRepay", nil)
	cdc.RegisterConcrete(MsgBankDeposit{}, "dipperProtocol/BankDeposit", nil)
	cdc.RegisterConcrete(MsgBankWithdraw{}, "dipperProtocol/BankWithdraw", nil)
	cdc.RegisterConcrete(MsgSetOraclePrice{}, "dipperProtocol/SetOraclePrice", nil)
}
