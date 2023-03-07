package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var RootCmd = &cobra.Command{
	Use:   "viper-test",
	Short: "Show a bug in viper",
	Long:  `Welcome to go-viper-bug!`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error = nil
		fmt.Printf("cmd database-url: %s\n", viper.GetString("database-url"))
		return err
	},
}

func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
