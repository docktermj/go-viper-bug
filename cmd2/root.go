/*
 */
package cmd2

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cmd2",
	Short: "cmd2",
	Long:  `cmd2`,
	PreRun: func(cmd *cobra.Command, args []string) {
		viper.BindPFlag("database-url", cmd.Flags().Lookup("database-url"))
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error = nil
		fmt.Printf("cmd2 database-url: %s\n", viper.GetString("database-url"))
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
	fmt.Printf(">>>>>> cmd2.init()\n")
	defaultDatabaseUrl := ""
	RootCmd.Flags().String("database-url", defaultDatabaseUrl, "URL of database to initialize [TEST_DATABASE_URL]")

	// Integrate with Viper.

	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("TEST")

	// Define flags in Viper.

	fmt.Printf(">>>>>> cmd2.init() %v\n", RootCmd.Flags().Lookup("database-url"))

	viper.SetDefault("database-url", defaultDatabaseUrl)
	viper.BindPFlag("database-url", RootCmd.Flags().Lookup("database-url"))

}

func initConfig() {
	viper.AutomaticEnv()
}
