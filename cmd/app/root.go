package app

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(Run)
}

var Cmd = &cobra.Command{
	Use:   "app",
	Short: "app commands",
}
