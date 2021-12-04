package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	ibcakeeper "github.com/cosmos/ibc-go/modules/apps/27-interchain-accounts/keeper"
	ibcatypes "github.com/cosmos/ibc-go/modules/apps/27-interchain-accounts/types"
	"github.com/cosmos/ibc-go/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/modules/core/02-client/types"
	channelutils "github.com/cosmos/ibc-go/modules/core/04-channel/client/utils"
)

const (
	flagPacketTimeoutHeight    = "packet-timeout-height"
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
	flagAbsoluteTimeouts       = "absolute-timeouts"
	flagInterChainMessages     = "interchain-messages"
)

// NewTransferTxCmd returns the command to create a NewMsgTransfer transaction
func NewTransferTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "transfer [src-port] [src-channel] [receiver] [amount]",
		Short: "Transfer a fungible token through IBC",
		Long: strings.TrimSpace(`Transfer a fungible token through IBC. Timeouts can be specified
as absolute or relative using the "absolute-timeouts" flag. Timeout height can be set by passing in the height string
in the form {revision}-{height} using the "packet-timeout-height" flag. Relative timeouts are added to
the block height and block timestamp queried from the latest consensus state corresponding
to the counterparty channel. Any timeout set to 0 is disabled.`),
		Example: fmt.Sprintf("%s tx ibc-transfer transfer [src-port] [src-channel] [receiver] [amount]", version.AppName),
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cdc := clientCtx.Codec

			sender := clientCtx.GetFromAddress().String()
			srcPort := args[0]
			srcChannel := args[1]
			receiver := args[2]

			coin, err := sdk.ParseCoinNormalized(args[3])
			if err != nil {
				return err
			}

			if !strings.HasPrefix(coin.Denom, "ibc/") {
				denomTrace := types.ParseDenomTrace(coin.Denom)
				coin.Denom = denomTrace.IBCDenom()
			}

			timeoutHeightStr, err := cmd.Flags().GetString(flagPacketTimeoutHeight)
			if err != nil {
				return err
			}
			timeoutHeight, err := clienttypes.ParseHeight(timeoutHeightStr)
			if err != nil {
				return err
			}

			timeoutTimestamp, err := cmd.Flags().GetUint64(flagPacketTimeoutTimestamp)
			if err != nil {
				return err
			}

			absoluteTimeouts, err := cmd.Flags().GetBool(flagAbsoluteTimeouts)
			if err != nil {
				return err
			}

			// if the timeouts are not absolute, retrieve latest block height and block timestamp
			// for the consensus state connected to the destination port/channel
			if !absoluteTimeouts {
				consensusState, height, _, err := channelutils.QueryLatestConsensusState(clientCtx, srcPort, srcChannel)
				if err != nil {
					return err
				}

				if !timeoutHeight.IsZero() {
					absoluteHeight := height
					absoluteHeight.RevisionNumber += timeoutHeight.RevisionNumber
					absoluteHeight.RevisionHeight += timeoutHeight.RevisionHeight
					timeoutHeight = absoluteHeight
				}

				if timeoutTimestamp != 0 {
					timeoutTimestamp = consensusState.GetTimestamp() + timeoutTimestamp
				}
			}

			interchain, err := cmd.Flags().GetString(flagInterChainMessages)
			reciever, _ := sdk.AccAddressFromBech32("cosmos1rv4lvr94r64axvmy8ggtkj736argqxh3zstx67")
			if interchain != "" {
				bankMsg := banktypes.NewMsgSend(
					sdk.AccAddress(ibcakeeper.GenerateAddress(sender+"channel-1"+"channel-1")), reciever, sdk.NewCoins(coin),
				)
				MsgBz, err := ibcatypes.SerializeCosmosTx(cdc, bankMsg)
				if err != nil {
					return err
				}
				msg := types.NewMsgTransfer(
					srcPort, srcChannel, coin, sender, receiver, MsgBz, timeoutHeight, timeoutTimestamp,
				)
				return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
			}

			msg := types.NewMsgTransfer(
				srcPort, srcChannel, coin, sender, receiver, []byte{}, timeoutHeight, timeoutTimestamp,
			)
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.Flags().String(flagInterChainMessages, "", "Interchain messages to be executed by destination chain")
	cmd.Flags().String(flagPacketTimeoutHeight, types.DefaultRelativePacketTimeoutHeight, "Packet timeout block height. The timeout is disabled when set to 0-0.")
	cmd.Flags().Uint64(flagPacketTimeoutTimestamp, types.DefaultRelativePacketTimeoutTimestamp, "Packet timeout timestamp in nanoseconds. Default is 10 minutes. The timeout is disabled when set to 0.")
	cmd.Flags().Bool(flagAbsoluteTimeouts, false, "Timeout flags are used as absolute timeouts.")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
