package main

import (
    "github.com/spf13/cobra"
    "os"
    "fmt"
)

func main() {
    var bcliCmd = &cobra.Command{
        Use:   "bcli",
        Short: "The Blockchain 101 CLI",
        Run: func(cmd *cobra.Command, args []string) {
        },
    }

    bcliCmd.AddCommand(versionCmd)
    bcliCmd.AddCommand(balancesCmd())
    bcliCmd.AddCommand(txCmd())

    err := bcliCmd.Execute()
    if err != nil {
        fmt.Fprintln(os.Stderr, err)
        os.Exit(1)
    }
}

func incorrectUsageErr() error {
	return fmt.Errorf("incorrect usage")
}
