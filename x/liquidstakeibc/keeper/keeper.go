package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	ibckeeper "github.com/cosmos/ibc-go/v6/modules/core/keeper"
	ibctmtypes "github.com/cosmos/ibc-go/v6/modules/light-clients/07-tendermint/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/persistenceOne/pstake-native/v2/x/liquidstakeibc/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey

	accountKeeper       types.AccountKeeper
	bankKeeper          types.BankKeeper
	epochsKeeper        types.EpochsKeeper
	icaControllerKeeper types.ICAControllerKeeper
	ibcKeeper           *ibckeeper.Keeper
	icqKeeper           types.ICQKeeper

	paramSpace paramtypes.Subspace

	msgRouter *baseapp.MsgServiceRouter

	authority string
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	epochsKeeper types.EpochsKeeper,
	icaControllerKeeper types.ICAControllerKeeper,
	ibcKeeper *ibckeeper.Keeper,
	icqKeeper types.ICQKeeper,

	paramSpace paramtypes.Subspace,

	msgRouter *baseapp.MsgServiceRouter,

	authority string,
) Keeper {
	// set KeyTable if it has not already been set
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:                 cdc,
		accountKeeper:       accountKeeper,
		bankKeeper:          bankKeeper,
		epochsKeeper:        epochsKeeper,
		icaControllerKeeper: icaControllerKeeper,
		ibcKeeper:           ibcKeeper,
		icqKeeper:           icqKeeper,
		storeKey:            storeKey,
		paramSpace:          paramSpace,
		msgRouter:           msgRouter,
		authority:           authority,
	}
}

// Logger returns a module-specific logger.
func (k *Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// GetParams gets the total set of liquidstakeibc parameters.
func (k *Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of liquidstakeibc parameters.
func (k *Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramSpace.SetParamSet(ctx, &params)
}

// GetDepositModuleAccount returns deposit module account interface
func (k *Keeper) GetDepositModuleAccount(ctx sdk.Context) authtypes.ModuleAccountI {
	return k.accountKeeper.GetModuleAccount(ctx, types.DepositModuleAccount)
}

// SendProtocolFee to the community pool
func (k *Keeper) SendProtocolFee(ctx sdk.Context, protocolFee sdk.Coins, moduleAccount, feeAddress string) error {
	addr, err := sdk.AccAddressFromBech32(feeAddress)
	if err != nil {
		return err
	}
	err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, moduleAccount, addr, protocolFee)
	if err != nil {
		return err
	}
	return nil
}

// GetClientState retrieves the client state given a connection id
func (k *Keeper) GetClientState(ctx sdk.Context, connectionID string) (*ibctmtypes.ClientState, error) {
	conn, found := k.ibcKeeper.ConnectionKeeper.GetConnection(ctx, connectionID)
	if !found {
		return nil, fmt.Errorf("invalid connection id, \"%s\" not found", connectionID)
	}

	clientState, found := k.ibcKeeper.ClientKeeper.GetClientState(ctx, conn.ClientId)
	if !found {
		return nil, fmt.Errorf("client id \"%s\" not found for connection \"%s\"", conn.ClientId, connectionID)
	}

	client, ok := clientState.(*ibctmtypes.ClientState)
	if !ok {
		return nil, fmt.Errorf("invalid client state for connection \"%s\"", connectionID)
	}

	return client, nil
}

// GetLatestConsensusState retrieves the last tendermint consensus state
func (k *Keeper) GetLatestConsensusState(ctx sdk.Context, connectionID string) (*ibctmtypes.ConsensusState, error) {
	conn, found := k.ibcKeeper.ConnectionKeeper.GetConnection(ctx, connectionID)
	if !found {
		return nil, fmt.Errorf("invalid connection id, \"%s\" not found", connectionID)
	}

	consensusState, found := k.ibcKeeper.ClientKeeper.GetLatestClientConsensusState(ctx, conn.ClientId)
	if !found {
		return nil, fmt.Errorf("client id \"%s\" not found for connection \"%s\"", conn.ClientId, connectionID)
	}

	state, ok := consensusState.(*ibctmtypes.ConsensusState)
	if !ok {
		return nil, fmt.Errorf("invalid consensus state for connection \"%s\"", connectionID)
	}

	return state, nil
}

// GetChainID gets the id of the host chain given a connection id
func (k *Keeper) GetChainID(ctx sdk.Context, connectionID string) (string, error) {
	clientState, err := k.GetClientState(ctx, connectionID)
	if err != nil {
		return "", fmt.Errorf("client state not found for connection \"%s\": \"%s\"", connectionID, err.Error())
	}

	return clientState.ChainId, nil
}

// RegisterICAAccount registers an ICA
func (k *Keeper) RegisterICAAccount(ctx sdk.Context, connectionID, owner string) error {
	return k.icaControllerKeeper.RegisterInterchainAccount(
		ctx,
		connectionID,
		owner,
		"",
	)
}

// RecreateHostChainICAAccounts recreates ICA if the channels are closed
func (k *Keeper) RecreateHostChainICAAccounts(ctx sdk.Context, hc *types.HostChain) (bool, error) {
	accountsRecreating := hc.DelegationAccount.ChannelState == types.ICAAccount_ICA_CHANNEL_CREATING &&
		hc.RewardsAccount.ChannelState == types.ICAAccount_ICA_CHANNEL_CREATING

	if accountsRecreating {
		return true, nil
	}

	_, isDelegateActive := k.icaControllerKeeper.GetOpenActiveChannel(
		ctx,
		hc.ConnectionId,
		icatypes.ControllerPortPrefix+k.DelegateAccountPortOwner(hc.ChainId),
	)
	if !isDelegateActive && hc.DelegationAccount.ChannelState == types.ICAAccount_ICA_CHANNEL_CREATED {
		if err := k.RegisterICAAccount(ctx, hc.ConnectionId, k.DelegateAccountPortOwner(hc.ChainId)); err != nil {
			return false, fmt.Errorf("error recreating %s delegate ica: %w", hc.ChainId, err)
		}
		accountsRecreating = true
	}

	_, isRewardsActive := k.icaControllerKeeper.GetOpenActiveChannel(
		ctx,
		hc.ConnectionId,
		icatypes.ControllerPortPrefix+k.RewardsAccountPortOwner(hc.ChainId),
	)
	if !isRewardsActive {
		if err := k.RegisterICAAccount(ctx, hc.ConnectionId, k.RewardsAccountPortOwner(hc.ChainId)); err != nil {
			return false, fmt.Errorf("error recreating %s rewards ica: %w", hc.ChainId, err)
		}
		accountsRecreating = true
	}

	return accountsRecreating, nil
}

// DelegateAccountPortOwner generates a delegate ICA port owner given the chain id
func (k *Keeper) DelegateAccountPortOwner(chainID string) string {
	return chainID + "." + types.DelegateICAType
}

// RewardsAccountPortOwner generates a rewards ICA port owner given the chain id
func (k *Keeper) RewardsAccountPortOwner(chainID string) string {
	return chainID + "." + types.RewardsICAType
}

func (k *Keeper) GetEpochNumber(ctx sdk.Context, epoch string) int64 {
	return k.epochsKeeper.GetEpochInfo(ctx, epoch).CurrentEpoch
}