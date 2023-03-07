package cmd

import "github.com/docktermj/go-viper-bug/cmd3"

func init() {
	RootCmd.AddCommand(cmd3.RootCmd)
}
