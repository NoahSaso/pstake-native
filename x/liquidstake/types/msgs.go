package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	_ sdk.Msg = (*MsgLiquidStake)(nil)
	_ sdk.Msg = (*MsgLiquidUnstake)(nil)
	_ sdk.Msg = (*MsgUpdateParams)(nil)
)

// Message types for the liquidstake module
const (
	MsgTypeLiquidStake   = "liquid_stake"
	MsgTypeLiquidUnstake = "liquid_unstake"
	MsgTypeUpdateParams  = "update_params"
)

// NewMsgLiquidStake creates a new MsgLiquidStake.
func NewMsgLiquidStake(
	liquidStaker sdk.AccAddress,
	amount sdk.Coin,
) *MsgLiquidStake {
	return &MsgLiquidStake{
		DelegatorAddress: liquidStaker.String(),
		Amount:           amount,
	}
}

func (m *MsgLiquidStake) Route() string { return RouterKey }

func (m *MsgLiquidStake) Type() string { return MsgTypeLiquidStake }

func (m *MsgLiquidStake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.DelegatorAddress); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address %q: %v", m.DelegatorAddress, err)
	}
	if ok := m.Amount.IsZero(); ok {
		return errorsmod.Wrap(sdkerrors.ErrInvalidRequest, "staking amount must not be zero")
	}
	if err := m.Amount.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgLiquidStake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgLiquidStake) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m *MsgLiquidStake) GetDelegator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgLiquidUnstake creates a new MsgLiquidUnstake.
func NewMsgLiquidUnstake(
	liquidStaker sdk.AccAddress,
	amount sdk.Coin,
) *MsgLiquidUnstake {
	return &MsgLiquidUnstake{
		DelegatorAddress: liquidStaker.String(),
		Amount:           amount,
	}
}

func (m *MsgLiquidUnstake) Route() string { return RouterKey }

func (m *MsgLiquidUnstake) Type() string { return MsgTypeLiquidUnstake }

func (m *MsgLiquidUnstake) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.DelegatorAddress); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid delegator address %q: %v", m.DelegatorAddress, err)
	}
	if ok := m.Amount.IsZero(); ok {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "unstaking amount must not be zero")
	}
	if err := m.Amount.Validate(); err != nil {
		return err
	}
	return nil
}

func (m *MsgLiquidUnstake) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

func (m *MsgLiquidUnstake) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m *MsgLiquidUnstake) GetDelegator() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return addr
}

// NewMsgUpdateParams creates a new MsgUpdateParams.
func NewMsgUpdateParams(authority sdk.AccAddress, amount Params) *MsgUpdateParams {
	return &MsgUpdateParams{
		Authority: authority.String(),
		Params:    amount,
	}
}

func (m *MsgUpdateParams) Route() string {
	return RouterKey
}

// Type should return the action
func (m *MsgUpdateParams) Type() string {
	return MsgTypeUpdateParams
}

// GetSignBytes encodes the message for signing
func (m *MsgUpdateParams) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(m))
}

// GetSigners defines whose signature is required
func (m *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(m.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{addr}
}

func (m *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(m.Authority); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address %q: %v", m.Authority, err)
	}

	err := m.Params.Validate()
	if err != nil {
		return err
	}
	return nil
}
