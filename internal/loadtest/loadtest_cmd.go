package loadtest

import (
	"log"

	"github.com/spf13/cobra"
)

var LoadCmd = &cobra.Command{
	Use:   "load",
	Short: "this feature is simple load testing for server",
	RunE: func(cmd *cobra.Command, args []string) error {

		for {
			log.Print("1 ")
		}
	},
}

func init() {
	LoadCmd.Flags()
}
