package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var subtractcmd = &cobra.Command{
	Use:     "subtract",
	Aliases: []string{"sub"},
	Short:   "subtracta number from anorther",
	Long:    "Carry out subtraction operation on 2 integers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Subtraction of %s from %s equals %s.\n\n", args[1], args[0], Subtract(args[0], args[1]))
	},
}

func init() {
	rootcmd.AddCommand(subtractcmd)
}
