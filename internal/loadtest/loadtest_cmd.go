package loadtest

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

var LoadCmd = &cobra.Command{
	Use:   "load",
	Short: "this feature is simple load testing for server",
	RunE: func(cmd *cobra.Command, args []string) error {
		for {
			resp, err := http.Get(args[0])
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			fmt.Print("OK")
		}
	},
}
