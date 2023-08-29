package cli

import (
	"fmt"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/version"
	"github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/arnac-io/osmosis/osmoutils/osmocli"
	"github.com/arnac-io/osmosis/v17/x/protorev/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewCmdTx returns the cli transaction commands for this module
func NewCmdTx() *cobra.Command {
	txCmd := osmocli.TxIndexCmd(types.ModuleName)
	osmocli.AddTxCmd(txCmd, CmdSetDeveloperAccount)
	osmocli.AddTxCmd(txCmd, CmdSetMaxPoolPointsPerTx)
	osmocli.AddTxCmd(txCmd, CmdSetMaxPoolPointsPerBlock)
	txCmd.AddCommand(
		CmdSetDeveloperHotRoutes().BuildCommandCustomFn(),
		CmdSetInfoByPoolType().BuildCommandCustomFn(),
		CmdSetBaseDenoms().BuildCommandCustomFn(),
		CmdSetProtoRevAdminAccountProposal(),
		CmdSetProtoRevEnabledProposal(),
	)
	return txCmd
}

// CmdSetDeveloperHotRoutes implements the command to set the protorev hot routes
func CmdSetDeveloperHotRoutes() *osmocli.TxCliDesc {
	desc := osmocli.TxCliDesc{
		Use:   "set-hot-routes [path/to/routes.json]",
		Short: "set the protorev hot routes",
		Long: `Must provide a json file with all of the hot routes that will be set. 
		Sample json file:
		[
			{
				"token_in": "uosmo",
				"token_out": "ibc/123...",
				"arb_routes" : [
					{
						"trades": [
							{
								"pool": 1,
								"token_in": "uosmo",
								"token_out": "uatom"
							},
							{
								"pool": 2,
								"token_in": "uatom",
								"token_out": "ibc/123..."
							},
							{
								"pool": 0,
								"token_in": "ibc/123...",
								"token_out": "uosmo"
							}
						],
						"step_size": 1000000
					}
				]
			}
		]
		`,
		Example:          fmt.Sprintf(`$ %s tx protorev set-hot-routes routes.json --from mykey`, version.AppName),
		NumArgs:          1,
		ParseAndBuildMsg: BuildSetHotRoutesMsg,
	}

	return &desc
}

// CmdSetDeveloperAccount implements the command to set the protorev developer account
func CmdSetDeveloperAccount() (*osmocli.TxCliDesc, *types.MsgSetDeveloperAccount) {
	return &osmocli.TxCliDesc{
		Use:     "set-developer-account [sdk.AccAddress]",
		Short:   "set the protorev developer account",
		NumArgs: 1,
		ParseAndBuildMsg: func(clientCtx client.Context, args []string, flags *pflag.FlagSet) (sdk.Msg, error) {
			developer, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return nil, err
			}

			return &types.MsgSetDeveloperAccount{
				DeveloperAccount: developer.String(),
				Admin:            clientCtx.GetFromAddress().String(),
			}, nil
		},
	}, &types.MsgSetDeveloperAccount{}
}

// CmdSetMaxPoolPointsPerTx implements the command to set the max pool points per tx
func CmdSetMaxPoolPointsPerTx() (*osmocli.TxCliDesc, *types.MsgSetMaxPoolPointsPerTx) {
	return &osmocli.TxCliDesc{
		Use:     "set-max-pool-points-per-tx [uint64]",
		Short:   "set the max pool points that can be consumed per tx",
		NumArgs: 1,
		ParseAndBuildMsg: func(clientCtx client.Context, args []string, flags *pflag.FlagSet) (sdk.Msg, error) {
			maxPoolPointsPerTx, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return nil, err
			}

			return &types.MsgSetMaxPoolPointsPerTx{
				MaxPoolPointsPerTx: maxPoolPointsPerTx,
				Admin:              clientCtx.GetFromAddress().String(),
			}, nil
		},
	}, &types.MsgSetMaxPoolPointsPerTx{}
}

// CmdSetMaxPoolPointsPerBlock implements the command to set the max pool points per block
func CmdSetMaxPoolPointsPerBlock() (*osmocli.TxCliDesc, *types.MsgSetMaxPoolPointsPerBlock) {
	return &osmocli.TxCliDesc{
		Use:     "set-max-pool-points-per-block [uint64]",
		Short:   "set the max pool points that can be consumed per block",
		NumArgs: 1,
		ParseAndBuildMsg: func(clientCtx client.Context, args []string, flags *pflag.FlagSet) (sdk.Msg, error) {
			maxPoolPointsPerBlock, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return nil, err
			}

			return &types.MsgSetMaxPoolPointsPerBlock{
				MaxPoolPointsPerBlock: maxPoolPointsPerBlock,
				Admin:                 clientCtx.GetFromAddress().String(),
			}, nil
		},
	}, &types.MsgSetMaxPoolPointsPerBlock{}
}

// CmdSetInfoByPoolType implements the command to set the pool information used throughout the module
func CmdSetInfoByPoolType() *osmocli.TxCliDesc {
	desc := osmocli.TxCliDesc{
		Use:   "set-info-by-pool-type [path/to/pool_info.json]",
		Short: "set the protorev pool type info",
		Long: `Must provide a json file with all the pool info that will be set. This does NOT set info for a single pool type.
		All information must be provided across all pool types in the json file.
		Sample json file:
		{
			"stable" : {
				"weight" : 1,
			},
			"concentrated" : {
				"weight" : 1,
				"max_ticks_crossed": 10,
			},
			"balancer" : {
				"weight" : 1,
			},
			"cosmwasm" : {
				"weight_maps" : [
					{"contract_address" : "cosmos123...", "weight" : 1}
				],
			},
		}
		`,
		Example:          fmt.Sprintf(`$ %s tx protorev set-info-by-pool-type pool_info.json --from mykey`, version.AppName),
		NumArgs:          1,
		ParseAndBuildMsg: BuildSetInfoByPoolTypeMsg,
	}

	return &desc
}

// CmdSetBaseDenoms implements the command to set the base denoms used in the highest liquidity method
func CmdSetBaseDenoms() *osmocli.TxCliDesc {
	desc := osmocli.TxCliDesc{
		Use:   "set-base-denoms [path/to/denoms.json]",
		Short: "set the protorev base denoms",
		Long: `Must provide a json file with all the base denoms that will be set. 
		Sample json file:
		[
			{
				"step_size" : 10000,
				"denom" : "uosmo"
			},
			{
				"step_size" : 10000,
				"denom" : "atom"
			}
		]
		`,
		Example:          fmt.Sprintf(`$ %s tx protorev set-base-denoms denoms.json --from mykey`, version.AppName),
		NumArgs:          1,
		ParseAndBuildMsg: BuildSetBaseDenomsMsg,
	}

	return &desc
}

// CmdSetProtoRevAdminAccountProposal implements the command to submit a SetProtoRevAdminAccountProposal
func CmdSetProtoRevAdminAccountProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-admin-account-proposal [sdk.AccAddress]",
		Args:    cobra.ExactArgs(1),
		Short:   "submit a set protorev admin account proposal to set the admin account for x/protorev",
		Example: fmt.Sprintf(`$ %s tx protorev set-protorev-admin-account osmo123... --from mykey`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			createContent := func(title string, description string, args ...string) (govtypes.Content, error) {
				return types.NewSetProtoRevAdminAccountProposal(title, description, args[0]), nil
			}

			return ProposalExecute(cmd, args, createContent)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)

	return cmd
}

// CmdSetProtoRevEnabledProposal implements the command to submit a SetProtoRevEnabledProposal
func CmdSetProtoRevEnabledProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "set-enabled-proposal [boolean]",
		Args:    cobra.ExactArgs(1),
		Short:   "submit a set protorev enabled proposal to enable or disable the protocol",
		Example: fmt.Sprintf(`$ %s tx protorev set-protorev-enabled true --from mykey`, version.AppName),
		RunE: func(cmd *cobra.Command, args []string) error {
			createContent := func(title string, description string, args ...string) (govtypes.Content, error) {
				res, err := strconv.ParseBool(args[0])
				if err != nil {
					return nil, err
				}

				content := types.NewSetProtoRevEnabledProposal(title, description, res)
				return content, nil
			}

			return ProposalExecute(cmd, args, createContent)
		},
	}

	cmd.Flags().String(cli.FlagTitle, "", "title of proposal")
	cmd.Flags().String(cli.FlagDescription, "", "description of proposal")
	cmd.Flags().String(cli.FlagDeposit, "", "deposit of proposal")
	flags.AddTxFlagsToCmd(cmd)
	_ = cmd.MarkFlagRequired(cli.FlagTitle)
	_ = cmd.MarkFlagRequired(cli.FlagDescription)

	return cmd
}

// ProposalExecute is a helper function to execute a proposal command. It takes in a function to create the proposal content.
func ProposalExecute(cmd *cobra.Command, args []string, createContent func(title string, description string, args ...string) (govtypes.Content, error)) error {
	clientCtx, err := client.GetClientTxContext(cmd)
	if err != nil {
		return err
	}

	title, err := cmd.Flags().GetString(cli.FlagTitle)
	if err != nil {
		return err
	}

	description, err := cmd.Flags().GetString(cli.FlagDescription)
	if err != nil {
		return err
	}

	depositStr, err := cmd.Flags().GetString(cli.FlagDeposit)
	if err != nil {
		return err
	}

	deposit, err := sdk.ParseCoinsNormalized(depositStr)
	if err != nil {
		return err
	}

	from := clientCtx.GetFromAddress()

	content, err := createContent(title, description, args...)
	if err != nil {
		return err
	}

	msg, err := govtypes.NewMsgSubmitProposal(content, deposit, from)
	if err != nil {
		return err
	}

	if err = msg.ValidateBasic(); err != nil {
		return err
	}

	return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
}
