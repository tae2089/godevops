package file

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var fileDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "this feature is to delete files that you insert day that you want to delete older days",
	RunE: func(cmd *cobra.Command, args []string) error {

		dirPath := strings.TrimSpace(viper.GetString("path"))
		if dirPath == "" {
			return errors.New("path is not set")
		}
		day := viper.GetInt("day")
		if day == 0 {
			return errors.New("day is not set")
		}
		err := deleteOldFiles(dirPath, day)
		return err
	},
}

func init() {
	fileDeleteCmd.PersistentFlags().StringP("path", "p", "", `[required] you can enter the path of the directory that you want to delete files`)
	fileDeleteCmd.PersistentFlags().StringP("day", "d", "", `[required] you want to delete files that you insert day that you want to delete older days and day is bigger than 0`)
	// mapping viper
	viper.BindPFlag("path", fileDeleteCmd.PersistentFlags().Lookup("path"))
	viper.BindPFlag("day", fileDeleteCmd.PersistentFlags().Lookup("day"))
	FileCmd.AddCommand(fileDeleteCmd)
}
