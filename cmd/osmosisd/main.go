package main

import (
	"os"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"

	osmosis "github.com/arnac-io/osmosis/v17/app"
	"github.com/arnac-io/osmosis/v17/app/params"
	"github.com/arnac-io/osmosis/v17/cmd/osmosisd/cmd"
)

func main() {
	params.SetAddressPrefixes()
	rootCmd, _ := cmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, osmosis.DefaultNodeHome); err != nil {
		os.Exit(1)
	}
}
