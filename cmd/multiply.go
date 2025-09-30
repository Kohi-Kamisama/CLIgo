package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var roundup bool
var multiplycmd = &cobra.Command{
	Use:     "multiply",
	Aliases: []string{"mul"},
	Short:   "Multiply 2 numbers",
	Long:    "Carry out multiplication operations on 2 numbers",
	Args:    cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Multiplication of %s by %s equals %s\n\n", args[0], args[1], Multiply(args[0], args[1], roundup))
	},
}

func init() {
	multiplycmd.Flags().BoolVarP(&roundup, "round", "r", false, "Round results up to 2 decimal places")
	rootcmd.AddCommand(multiplycmd)
}
