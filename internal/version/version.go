package version

import (
	"log"

	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"rev"},
	Short:   "print the version",
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("godevops version - v0.0.1")
	},
}
