package cli

import (
	"github.com/spf13/cobra"

	"github.com/arnac-io/osmosis/osmoutils/osmocli"
	"github.com/arnac-io/osmosis/v17/x/ibc-rate-limit/client/queryproto"
	"github.com/arnac-io/osmosis/v17/x/ibc-rate-limit/types"
)

// GetQueryCmd returns the cli query commands for this module.
func GetQueryCmd() *cobra.Command {
	cmd := osmocli.QueryIndexCmd(types.ModuleName)

	cmd.AddCommand(
		osmocli.GetParams[*queryproto.ParamsRequest](
			types.ModuleName, queryproto.NewQueryClient),
	)

	return cmd
}
