package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"fairyring/x/keyshare/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdRegisterValidator())
	cmd.AddCommand(CmdSendKeyshare())
	cmd.AddCommand(CmdCreateLatestPubKey())
	cmd.AddCommand(CmdCreateAuthorizedAddress())
	cmd.AddCommand(CmdUpdateAuthorizedAddress())
	cmd.AddCommand(CmdDeleteAuthorizedAddress())

	cmd.AddCommand(CmdCreateGeneralKeyShare())
	// this line is used by starport scaffolding # 1

	return cmd
}
