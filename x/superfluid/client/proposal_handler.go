package client

import (
	"github.com/arnac-io/osmosis/x/superfluid/client/cli"
	"github.com/arnac-io/osmosis/x/superfluid/client/rest"

	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"
)

var (
	SetSuperfluidAssetsProposalHandler    = govclient.NewProposalHandler(cli.NewCmdSubmitSetSuperfluidAssetsProposal, rest.ProposalSetSuperfluidAssetsRESTHandler)
	RemoveSuperfluidAssetsProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitRemoveSuperfluidAssetsProposal, rest.ProposalRemoveSuperfluidAssetsRESTHandler)
	UpdateUnpoolWhitelistProposalHandler  = govclient.NewProposalHandler(cli.NewCmdUpdateUnpoolWhitelistProposal, rest.ProposalUpdateUnpoolWhitelistProposal)
)
