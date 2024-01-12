package types

const (
	EventTypeLiquidStake                           = "liquid_stake"
	EventTypeLiquidStakeLSM                        = "liquid_stake_lsm"
	EventTypeLiquidUnstake                         = "liquid_unstake"
	EventTypeRedeem                                = "redeem"
	EventTypePacket                                = "ics27_packet"
	EventTypeTimeout                               = "timeout"
	EventTypeSlashing                              = "validator_slash"
	EventTypeUpdateParams                          = "update_params"
	EventTypeChainDisabled                         = "chain_disabled"
	EventTypeValidatorStatusUpdate                 = "validator_status_update"
	EventTypeValidatorExchangeRateUpdate           = "validator_exchange_rate_update"
	EventTypeValidatorDelegableStateUpdate         = "validator_delegable_state_update"
	EventTypeDoDelegation                          = "send_delegation"
	EventTypeDoDelegationDeposit                   = "send_individual_delegation"
	EventTypeClaimedUnbondings                     = "claimed_unbondings"
	EventTypeRedeemTokensForShares                 = "redeem_lsm_tokens_shares"
	EventTypeCValueUpdate                          = "c_value_update"
	EventTypeDelegationWorkflow                    = "delegation_workflow"
	EventTypeUndelegationWorkflow                  = "undelegation_workflow"
	EventTypeValidatorUndelegationWorkflow         = "validator_undelegation_workflow"
	EventTypeRewardsWorkflow                       = "rewards_workflow"
	EventTypeLSMWorkflow                           = "lsm_workflow"
	EventTypeRewardsTransfer                       = "rewards_transfer"
	EventTypeUnbondingMaturedReceived              = "unbonding_matured"
	EventTypeValidatorUnbondingMaturedReceived     = "validator_unbonding_matured"
	EventAutocompoundRewardsReceived               = "autocompound_rewards_received"
	EventStakingDepositTransferReceived            = "staking_deposit_received"
	EventStakingDepositTransferTimeout             = "staking_deposit_timeout"
	EventLSMDepositTransferReceived                = "lsm_deposit_received"
	EventLSMDepositTransferTimeout                 = "lsm_deposit_timeout"
	EventICAChannelCreated                         = "ica_channel_created"
	EventSuccessfulDelegation                      = "successful_delegation"
	EventSuccessfulUndelegation                    = "successful_undelegation"
	EventBurn                                      = "stk-burn"
	EventSuccessfulUndelegationTransfer            = "successful_undelegation_transfer"
	EventSuccessfulValidatorUndelegationTransfer   = "successful_validator_undelegation_transfer"
	EventSuccessfulLSMRedeem                       = "successful_lsm_redeem"
	EventSuccessfulRedelegation                    = "successful_redelegation"
	EventUnsuccessfulDelegation                    = "unsuccessful_delegation"
	EventUnsuccessfulUndelegation                  = "unsuccessful_undelegation"
	EventUnsuccessfulUndelegationInitiation        = "unsuccessful_undelegation_initiation"
	EventUnsuccessfulUndelegationTransfer          = "unsuccessful_undelegation_transfer"
	EventUnsuccessfulValidatorUndelegationTransfer = "unsuccessful_validator_undelegation_transfer"
	EventUnsuccessfulLSMRedeem                     = "unsuccessful_lsm_redeem"
	EventUnsuccessfulRedelegate                    = "unsuccessful_redelegate"
	EventFailedClaimUnbondings                     = "failed_claim_unbondings"

	AttributeInputAmount                     = "input_amount"
	AttributeOutputAmount                    = "output_amount"
	AttributeDelegatorAddress                = "address"
	AttributePstakeDepositFee                = "pstake_deposit_fee"
	AttributePstakeUnstakeFee                = "pstake_unstake_fee"
	AttributePstakeRedeemFee                 = "pstake_redeem_fee"
	AttributePstakeAutocompoundFee           = "autocompound_fee"
	AttributeChainID                         = "chain_id"
	AttributeNewCValue                       = "new_c_value"
	AttributeOldCValue                       = "old_c_value"
	AttributeEpoch                           = "epoch_number"
	AttributeValidatorAddress                = "validator_address"
	AttributeExistingDelegation              = "existing_delegation"
	AttributeUpdatedDelegation               = "updated_delegation"
	AttributeSlashedAmount                   = "slashed_amount"
	AttributeKeyAuthority                    = "authority"
	AttributeKeyUpdatedParams                = "updated_params"
	AttributeKeyAck                          = "acknowledgement"
	AttributeKeyAckSuccess                   = "success"
	AttributeKeyAckError                     = "error"
	AttributeKeyValidatorNewStatus           = "validator_new_status"
	AttributeKeyValidatorOldStatus           = "validator_old_status"
	AttributeKeyValidatorNewExchangeRate     = "validator_new_exchange_rate"
	AttributeKeyValidatorOldExchangeRate     = "validator_old_exchange_rate"
	AttributeKeyValidatorDelegable           = "validator_delegable"
	AttributeTotalDelegatedAmount            = "total_delegated_amount"
	AttributeIBCSequenceID                   = "ibc_sequence_id"
	AttributeICAMessages                     = "ica_messages"
	AttributeClaimAmount                     = "claimed_amount"
	AttributeClaimAddress                    = "claim_address"
	AttributeClaimStatus                     = "claim_status"
	AttributeModuleMintedAmount              = "minted_amount"
	AttributeModuleLSMTokenizedAmount        = "lsm_tokenized_amount"
	AttributeModuleStakedAmount              = "staked_amount"
	AttributeModuleAmountOnPersistence       = "amount_on_persistence"
	AttributeModuleAmountOnHostChain         = "amount_on_host_chain"
	AttributeModuleUnbondingAmount           = "unbonding_amount"
	AttributeTotalEpochDepositAmount         = "deposit_amount"
	AttributeTotalEpochUnbondingAmount       = "unbonding_amount"
	AttributeTotalEpochBurnAmount            = "burn_amount"
	AttributeValidatorUnbondingAmount        = "validator_unbonding_amount"
	AttributeLSMDepositsSharesAmount         = "lsm_deposits_shares_amount"
	AttributeRewardsTransferAmount           = "rewards_transfer_amount"
	AttributeRewardsBalanceAmount            = "rewards_balance_amount"
	AttributeUnbondingMaturedAmount          = "unbonding_matured_amount"
	AttributeValidatorUnbondingMaturedAmount = "validator_unbonding_matured_amount"
	AttributeAutocompoundTransfer            = "autocompound_transfer_amount"
	AttributeICAPortOwner                    = "ica_port_owner"
	AttributeICAChannelID                    = "ica_channel_id"
	AttributeICAAddress                      = "ica_address"
	AttributeDelegatedAmount                 = "delegated_amount"
	AttributeUndelegatedAmount               = "undelegated_amount"
	AttributeRedeemedAmount                  = "redeemed_amount"
	AttributeRedelegatedAmount               = "redelegated_amount"
	AttributeValidatorSrcAddress             = "redelegation_validator_src-address"
	AttributeValidatorDstAddress             = "redelegation_validator_dst-address"

	AttributeValueCategory = ModuleName
)
