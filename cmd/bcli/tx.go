package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"blockchain101.crypto/blockchain101/database"
	"os"
)

const flagFrom = "from"
const flagTo = "to"
const flagValue = "value"
const flagData = "data"

func txCmd() *cobra.Command {
    var txsCmd = &cobra.Command{
        Use:   "tx",
        Short: "Interact with txs (add...).",
        PreRunE: func(cmd *cobra.Command, args []string) error {
            return incorrectUsageErr()
        },
        Run: func(cmd *cobra.Command, args []string) {
        },
    }

    txsCmd.AddCommand(txAddCmd())

    return txsCmd
}

func txAddCmd() *cobra.Command {
    var cmd = &cobra.Command{
        Use:   "add",
        Short: "Adds new TX to database.",
        Run: func(cmd *cobra.Command, args []string) {
            from, _ := cmd.Flags().GetString(flagFrom)
            to, _ := cmd.Flags().GetString(flagTo)
            value, _ := cmd.Flags().GetUint(flagValue)
						data, _ := cmd.Flags().GetString(flagData)

            fromAcc := database.NewAccount(from)
            toAcc := database.NewAccount(to)

            tx := database.NewTx(fromAcc, toAcc, value, data)

            state, err := database.NewStateFromDisk()
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }

            // defer means, at the end of this function execution,
            // execute the following statement (close DB file with all TXs)
            defer state.Close()

            // Add the TX to an in-memory array (pool)
            err = state.Add(tx)
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }

            // Flush the mempool TXs to disk
            _, err = state.Persist()
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
                os.Exit(1)
            }

            fmt.Println("TX successfully added to the ledger.")
        },
    }

    cmd.Flags().String(flagFrom, "", "From what account to send tokens")
    cmd.MarkFlagRequired(flagFrom)

    cmd.Flags().String(flagTo, "", "To what account to send tokens")
    cmd.MarkFlagRequired(flagTo)

    cmd.Flags().Uint(flagValue, 0, "How many tokens to send")
    cmd.MarkFlagRequired(flagValue)

		cmd.Flags().String(flagData, "", "Possible values: 'reward'")

    return cmd
  }
