package myip

import (
	"github.com/spf13/cobra"
)

var Myipcmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"rev"},
	Short:   "print the version",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := GetMyIp()
		return err
	},
}
