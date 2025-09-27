package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addcmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"add"},
	Short:   "Adds 2 numbers",
	Long:    "Carry out addition operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("addition of %s and %s equals %s.\n\n", args[0], args[1], Add(args[0], args[1]))
	},
}

func init() {
	rootcmd.AddCommand(addcmd)
}
