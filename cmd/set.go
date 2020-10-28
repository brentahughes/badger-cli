package cmd

import (
	"fmt"
	"log"

	"github.com/getcouragenow/badger-cli/pkg/badger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:   "set [key] [value]",
	Short: "Set a key and its value",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := badger.Open(cmd.Flag("dir").Value.String())
		if err != nil {
			log.Fatalln(err)
		}
		defer db.Close()

		ttl := viper.GetDuration("ttl")
		var opts *badger.EntryOptions
		if ttl.Seconds() > 0 {
			opts = &badger.EntryOptions{
				TTL: ttl,
			}
		}

		if err := db.Set(args[0], args[1], opts); err != nil {
			log.Fatalln(err)
		}

		fmt.Println(args[0])
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
	setCmd.PersistentFlags().Duration("ttl", 0, "Set ttl for the new key")
	viper.BindPFlag("ttl", setCmd.PersistentFlags().Lookup("ttl"))
}
