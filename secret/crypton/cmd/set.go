package cmd

import (
	"fmt"
	"strings"

	"github.com/rushikesh2/GolangTraining/Gophercises/secret/encrypt"
	"github.com/spf13/cobra"
)

var Setcmd = &cobra.Command{
	Use:   "set",
	Short: "set makes your system secure by providing API to enccypt and decrypt ",
	Run: func(cmd *cobra.Command, args []string) {
		vault := encrypt.NewVault(nkey, ".Vault")
		//code to include blank spaces
		key := args[0]
		value := strings.Join(args[1:], " ")
		err := vault.Set(key, value)
		if err != nil {
			panic(err)
		}
		fmt.Println("Data successfully encoded.")
	},
}

func init() {
	Rootcmd.AddCommand(Setcmd)
}
