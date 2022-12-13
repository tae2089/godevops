package internal

import (
	"os"

	"github.com/tae2089/godevops/internal/base64"
	"github.com/tae2089/godevops/internal/redis"
	"github.com/tae2089/godevops/internal/loadtest"
	"github.com/tae2089/godevops/internal/version"


	"github.com/spf13/cobra"
	"github.com/tae2089/godevops/internal/myip"
)

var rootCmd = &cobra.Command{
	Use:   "godevops",
	Short: "Godevops is a Useful command-line interface for Devops.",
	Long: `A Fast and Flexible Static Site Generator built with
				  love by spf13 and friends in Go.
				  Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(myip.Myipcmd)
	rootCmd.AddCommand(base64.Base64Cmd)
	rootCmd.AddCommand(redis.RedisCmd)
	rootCmd.AddCommand(loadtest.LoadCmd)
}
