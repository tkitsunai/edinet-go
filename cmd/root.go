package cmd

import (
	"github.com/tkitsunai/edinet-go/conf"
	"github.com/tkitsunai/edinet-go/datastore"
	"github.com/tkitsunai/edinet-go/di"
	"github.com/tkitsunai/edinet-go/logger"
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
		var storage datastore.Engine
		persistent := conf.LoadServerConfig().Persistent
		persistence := persistent.IsPersistence()
		if persistence {
			storage = datastore.GetEngineByName(persistent.Engine)
		} else {
			storage = datastore.DefaultEngine
		}

		err := storage.Open()
		if err != nil {
			logger.Logger.Error().Msgf("storage open failed %s", err)
			return err
		}
		defer storage.Close()

		injector := di.SetUpContainer(storage.GetDriver())

		s := server.NewServer(injector, server.Config{
			Mode: server.OfMode(conf.LoadServerConfig().Mode),
		})

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
