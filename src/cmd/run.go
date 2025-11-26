package cmd

import (
	"github.com/spf13/cobra"
)

var summery bool
var runCmd = &cobra.Command{
	Use:   "run [file]",
	Short: "Run API test suite",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		return nil
	},
}
