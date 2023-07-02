package file

import (
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var FileDeleteCmd = &cobra.Command{
	Use:   "load",
	Short: "this feature is simple load testing for server",
	RunE: func(cmd *cobra.Command, args []string) error {
		dir := strings.TrimSpace(viper.GetString("dir"))
		day := viper.GetInt("day")
		err := deleteOldFiles(dir, day)
		return err
	},
}
