package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var dividecmd = &cobra.Command{
	Use:     "divide",
	Aliases: []string{"div"},
	Short:   "Divide one number by anorther",
	Long:    "Carry out division operation on 2 numbers",
	Args:    cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := Divide(args[0], args[1], roundup)
		if err != nil {
			return err
		}
		fmt.Printf("Divide %s by %s equals %s.\n\n", args[0], args[1], res)
		return nil
	},
}

func init() {
	dividecmd.Flags().BoolVarP(&roundup, "round", "r", false, "Round results up to 2 decimal places")
	rootcmd.AddCommand(dividecmd)
}
