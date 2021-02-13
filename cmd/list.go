package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/bah2830/badger-cli/pkg/badger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/dustin/go-humanize"
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
		fmt.Printf("Total space for shown keys %s and values %s, sum is %s\n", humanize.Bytes(uint64(totKeySize)), humanize.Bytes(uint64(totValSize)), humanize.Bytes(uint64(totValSize+int64(totKeySize))))

		var errGC error
		for errGC == nil {
			lsm, vlog := db.Size()
			total := uint64(lsm) + uint64(vlog)
			ratio := float64(totValSize)/float64(total)
			fmt.Printf("DB file size to size of values ratio %.2f\n", 1/ratio)
			fmt.Printf("DB sizes: lsm %s and vlog %s, total %s\n", humanize.Bytes(uint64(lsm)), humanize.Bytes(uint64(vlog)), humanize.Bytes(uint64(lsm+vlog)))
			errGC = db.RunValueLogGC(0.5)
			db.Close() // GC is claimed only on close :(
			db, err = badger.Open(cmd.Flag("dir").Value.String())
			// fmt.Errorf("See error %v", errGC)
			// lsm, vlog = db.Size()
			// fmt.Printf("After GC DB sizes: lsm %s and vlog %s, total %s\n", humanize.Bytes(uint64(lsm)), humanize.Bytes(uint64(vlog)), humanize.Bytes(uint64(lsm+vlog)))
		}
		lsm, vlog := db.Size()
		fmt.Printf("After all GC DB sizes: lsm %s and vlog %s, total %s\n", humanize.Bytes(uint64(lsm)), humanize.Bytes(uint64(vlog)), humanize.Bytes(uint64(lsm+vlog)))
		db.Close() // GC is claimed only on close :(
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
