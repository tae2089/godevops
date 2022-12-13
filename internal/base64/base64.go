package base64

import (
	b64 "encoding/base64"
	"log"

	"github.com/spf13/cobra"
)

var Base64Cmd = &cobra.Command{
	Use:     "base64",
	Aliases: []string{"b64"},
	Short:   "print the version",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var base64EncodeCmd = &cobra.Command{
	Use:   "encode",
	Short: "encode a base64 string",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := EncodingData(args[0])
		return err
	},
}

var base64DecodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "decode a base64 string",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := DecodingData(args[0])
		return err
	},
}

func EncodingData(data string) error {
	encodedData := b64.RawStdEncoding.EncodeToString([]byte(data))
	log.Print(encodedData)
	return nil
}

func DecodingData(data string) error {
	decodedData, err := b64.RawStdEncoding.DecodeString(data)
	if err != nil {
		return err
	}
	log.Print(string(decodedData))
	return nil
}

func init() {
	Base64Cmd.AddCommand(base64EncodeCmd, base64DecodeCmd)
}
