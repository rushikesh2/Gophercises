package main

import (
	"fmt"
	"path/filepath"

	home "github.com/mitchellh/go-homedir"
	"github.com/rushikesh2/Gophercises/CLI/todoz/cmd"
	"github.com/rushikesh2/Gophercises/CLI/todoz/store"
)

func main() {
	hdir, _ := home.Dir()

	path := filepath.Join(hdir, "todo.db")

	HandleError(store.Init(path))

	HandleError(cmd.RootCmd.Execute())
}

func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		return //os.Exit(1)
	}
}
