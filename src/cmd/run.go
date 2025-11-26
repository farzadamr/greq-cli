package cmd

import (
	"github.com/farzadamr/greq-cli/internal/config"
	"github.com/farzadamr/greq-cli/internal/executor"
	"github.com/farzadamr/greq-cli/internal/ui"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run [file]",
	Short: "Run API test suite",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		file := args[0]

		// 1. Load File
		suite, err := config.LoadSuite(file)
		if err != nil {
			//handle ui error
		}
		testCount := len(suite.Tests)
		ui.RenderHeader(file, testCount, "dev")

		// 2. Run HTTP requests
		results := executor.RunSuite(suite)
		if results == nil {
			//handle ui error
		}

		//3. UI Render
		ui.RenderTestsResult(results)
		return nil
	},
}
