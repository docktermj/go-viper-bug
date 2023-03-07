/*
 */
package cmd4

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cmd4",
	Short: "cmd4",
	Long:  `cmd4`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error = nil
		fmt.Printf("cmd4 database-url: %s\n", viper.GetString("database-url"))
		return err
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	defaultDatabaseUrl := ""
	RootCmd.Flags().String("database-url", defaultDatabaseUrl, "URL of database to initialize [TEST_DATABASE_URL]")

	// Integrate with Viper.

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("TEST")

	// Define flags in Viper.

	viper.SetDefault("database-url", defaultDatabaseUrl)
	viper.BindPFlag("database-url", RootCmd.Flags().Lookup("database-url"))

}

func initConfig() {
	viper.AutomaticEnv()
}
