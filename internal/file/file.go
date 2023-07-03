package file

import "github.com/spf13/cobra"

var (
	// rootCmd represents the base command when called without any sub-commands
	FileCmd = &cobra.Command{
		Use:   "file",
		Short: `The file subcommand includes functionalities such as deletion.`,
		Long:  `The file subcommand includes functionalities such as deletion.`,
	}
)
