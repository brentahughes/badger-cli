package cmd

import (
	"fmt"
	"log"

	"github.com/bah2830/badger-cli/pkg/badger"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get [key]",
	Short: "Get content of a specific key",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := badger.Open(cmd.Flag("dir").Value.String())
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Close()

		value, err := db.Get(args[0])
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
