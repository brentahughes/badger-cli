package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/getcouragenow/badger-cli/pkg/badger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List keys in the database",
	Run: func(cmd *cobra.Command, args []string) {
		db, err := badger.Open(cmd.Flag("dir").Value.String())
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Close()

		keys, total, err := db.List(cmd.Flag("prefix").Value.String(), viper.GetInt("limit"), viper.GetInt("offset"))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Printf("% -30s % 10s % 10s % 5s\n", "KEY", "VERSION", "SIZE", "META")
		fmt.Println(strings.Repeat("=", 60))
		for _, k := range keys {
			fmt.Println(k)
		}

		fmt.Printf("\n\nReturned keys:   %d\n", len(keys))
		fmt.Printf("Matched keys:    %d\n", total)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringP("prefix", "p", "", "Key prefix for the search")
	listCmd.PersistentFlags().IntP("limit", "l", 200, "Number of results to return")
	listCmd.PersistentFlags().IntP("offset", "o", 0, "Offset to start at")
	viper.BindPFlag("limit", listCmd.PersistentFlags().Lookup("limit"))
	viper.BindPFlag("offset", listCmd.PersistentFlags().Lookup("offset"))
}
