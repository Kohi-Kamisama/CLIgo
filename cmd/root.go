package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootcmd = &cobra.Command{
	Use:   "zero",
	Short: "zero is a cli tool for performing basic mathematical operations",
	Long:  "zero is a cli tool for preforming basic mathematical operations - additon, multiplication, division and subtraction",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	if err := rootcmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Oops. An eror while executing zero '%s'\n", err)
		os.Exit(1)

	}
}
