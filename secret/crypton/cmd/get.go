package cmd

import (
	"fmt"

	"github.com/rushikesh2/GolangTraining/Gophercises/secret/encrypt"
	"github.com/spf13/cobra"
)

var Getcmd = &cobra.Command{
	Use:   "get",
	Short: "get makes your system secure by providing API to enccypt and decrypt data ",
	Run: func(cmd *cobra.Command, args []string) {
		vault := encrypt.NewVault(nkey, ".Vault")
		key := args[0]
		value, err := vault.Get(key)
		if err != nil {
			fmt.Println("no values set")
			return
		}
		fmt.Printf("%s(key) = %s(value)", key, value)
	},
}

func init() {
	Rootcmd.AddCommand(Getcmd)
}
