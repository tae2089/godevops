package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"rev"},
	Short:   "print the version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("godevops version - v0.0.1")
	},
}
