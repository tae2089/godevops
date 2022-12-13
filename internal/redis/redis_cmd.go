package redis

import (
	"github.com/spf13/cobra"
)

var RedisCmd = &cobra.Command{
	Use:     "redis",
	Short:   "Redis Ping Check",
	Example: "args localhost:6379",
	Run: func(cmd *cobra.Command, args []string) {
		RedisPingCheck(args[0])

	},
	// RunE: func(cmd *cobra.Command, args []string) error {
	// 	err := RedisPingCheck(args[0])
	// 	return err
	// },
}
