package main

import (
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	"github.com/rushikesh2/GolangTraining/Gophercises/CLI/task/cmd"
	"github.com/rushikesh2/GolangTraining/Gophercises/CLI/task/db"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	HandleError(db.Init(dbPath))
	HandleError(cmd.RootCmd.Execute())
}

//HandleError it will handle errors
func HandleError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		//os.Exit(1)
	}
}
