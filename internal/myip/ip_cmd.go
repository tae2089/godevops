package myip

import (
	"github.com/spf13/cobra"
)

var Myipcmd = &cobra.Command{
	Use:   "ip",
	Short: "what is myip",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := GetMyIp()
		return err
	},
}
