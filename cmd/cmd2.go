package cmd

import cmd "github.com/docktermj/go-viper-bug/cmd2"

func init() {
	RootCmd.AddCommand(cmd.RootCmd)
}
