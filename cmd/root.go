package cmd

import (
	"github.com/tkitsunai/edinet-go/server"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "edinet-go",
	Short: "edinet-go",
	Long:  `edinet-go`,
	RunE: func(cmd *cobra.Command, args []string) error {
		s := server.NewServer()
		return s.Run()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
