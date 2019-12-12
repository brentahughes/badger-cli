package cmd

import (
	"log"

	"github.com/bah2830/badger-cli/pkg/badger"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a key and its contents",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := badger.Open(cmd.Flag("dir").Value.String())
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Close()

		if err := db.Delete(args...); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
