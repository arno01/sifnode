package ante

import (
	"strings"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	disptypes "github.com/Sifchain/sifnode/x/dispensation/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
)

func NewAnteHandler(options ante.HandlerOptions) (sdk.AnteHandler, error) {
	if options.AccountKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "account keeper is required for ante builder")
	}
	if options.BankKeeper == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "bank keeper is required for ante builder")
	}
	if options.SignModeHandler == nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrLogic, "sign mode handler is required for ante builder")
	}
	var sigGasConsumer = options.SigGasConsumer
	if sigGasConsumer == nil {
		sigGasConsumer = ante.DefaultSigVerificationGasConsumer
	}
	return sdk.ChainAnteDecorators(
		ante.NewSetUpContextDecorator(), // outermost AnteDecorator. SetUpContext must be called first
		NewAdjustGasPriceDecorator(),    // Custom decorator to adjust gas price for specific msg types
		ante.NewRejectExtensionOptionsDecorator(),
		ante.NewMempoolFeeDecorator(),
		ante.NewValidateBasicDecorator(),
		ante.NewTxTimeoutHeightDecorator(),
		ante.NewValidateMemoDecorator(options.AccountKeeper),
		ante.NewConsumeGasForTxSizeDecorator(options.AccountKeeper),
		ante.NewDeductFeeDecorator(options.AccountKeeper, options.BankKeeper, options.FeegrantKeeper),
		ante.NewSetPubKeyDecorator(options.AccountKeeper), // SetPubKeyDecorator must be called before all signature verification decorators
		ante.NewValidateSigCountDecorator(options.AccountKeeper),
		ante.NewSigGasConsumeDecorator(options.AccountKeeper, sigGasConsumer),
		ante.NewSigVerificationDecorator(options.AccountKeeper, options.SignModeHandler),
		ante.NewIncrementSequenceDecorator(options.AccountKeeper),
	), nil

}

// AdjustGasPriceDecorator is a custom decorator to reduce fee prices .
type AdjustGasPriceDecorator struct {
}

// NewAdjustGasPriceDecorator create a new instance of AdjustGasPriceDecorator
func NewAdjustGasPriceDecorator() AdjustGasPriceDecorator {
	return AdjustGasPriceDecorator{}
}

// AnteHandle adjusts the gas price based on the tx type.
func (r AdjustGasPriceDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (newCtx sdk.Context, err error) {
	msgs := tx.GetMsgs()
	if len(msgs) == 1 && (strings.Contains(strings.ToLower(sdk.MsgTypeURL(msgs[0])), strings.ToLower(disptypes.MsgTypeCreateDistribution)) ||
		strings.Contains(strings.ToLower(sdk.MsgTypeURL(msgs[0])), strings.ToLower(disptypes.MsgTypeRunDistribution))) {
		minGasPrice := sdk.DecCoin{
			Denom:  "rowan",
			Amount: sdk.MustNewDecFromStr("0.00000005"),
		}
		if !minGasPrice.IsValid() {
			return ctx, sdkerrors.Wrap(sdkerrors.ErrLogic, "invalid gas price")
		}
		ctx = ctx.WithMinGasPrices(sdk.NewDecCoins(minGasPrice))
		return next(ctx, tx, simulate)
	}
	minFee := sdk.ZeroInt()
	for i := range msgs {
		msgTypeURLLower := strings.ToLower(sdk.MsgTypeURL(msgs[i]))
		if strings.Contains(msgTypeURLLower, strings.ToLower(banktypes.TypeMsgSend)) ||
			strings.Contains(msgTypeURLLower, strings.ToLower(banktypes.TypeMsgMultiSend)) ||
			strings.Contains(msgTypeURLLower, "createuserclaim") ||
			strings.Contains(msgTypeURLLower, "swap") ||
			strings.Contains(msgTypeURLLower, "removeliquidity") ||
			strings.Contains(msgTypeURLLower, "addliquidity") {
			minFee = sdk.NewInt(100000000000000000) // 0.1
		} else if strings.Contains(msgTypeURLLower, "transfer") && minFee.LTE(sdk.NewInt(10000000000000000)) {
			minFee = sdk.NewInt(10000000000000000) // 0.01
		}
	}
	if minFee.Equal(sdk.ZeroInt()) {
		return next(ctx, tx, simulate)
	}
	feeTx, ok := tx.(sdk.FeeTx)
	if !ok {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrTxDecode, "tx must be a FeeTx")
	}
	fees := feeTx.GetFee()
	rowanFee := sdk.ZeroInt()
	for j := range fees {
		if strings.EqualFold("rowan", fees[j].Denom) {
			rowanFee = fees[j].Amount
		}
	}
	if rowanFee.LTE(sdk.ZeroInt()) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrLogic, "unsupported fee asset")
	}
	if rowanFee.LT(minFee) {
		return ctx, sdkerrors.Wrap(sdkerrors.ErrLogic, "tx fee is too low")
	}
	return next(ctx, tx, simulate)
}
