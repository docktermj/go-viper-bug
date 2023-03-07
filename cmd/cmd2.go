package cmd

import "github.com/docktermj/go-viper-bug/cmd2"

func init() {
	RootCmd.AddCommand(cmd2.RootCmd)
}
