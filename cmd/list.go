package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/bah2830/badger-cli/pkg/badger"
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

		var totValSize int64 = 0
		var totKeySize int = 0
		fmt.Printf("% 10s % 10s % 5s % -30s \n", "SIZE", "VERSION", "META", "KEY")
		fmt.Println(strings.Repeat("=", 60))
		for _, k := range keys {
			fmt.Println(k)
			totValSize += k.Size
			totKeySize += len(k.Key)
		}

		fmt.Printf("\n\nReturned keys:   %d\n", len(keys))
		fmt.Printf("Matched keys:    %d\n", total)
		fmt.Printf("Total space for shown keys %d and values %d, sum is %d\n", totKeySize, totValSize, totValSize+int64(totKeySize))

		lsm, vlog := db.Size()
		fmt.Printf("DB sizes: lsm %d and vlog %d, total %d\n", lsm, vlog, lsm+vlog)
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
