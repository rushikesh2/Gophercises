package cmd

import (
	"path/filepath"
	"testing"

	home "github.com/mitchellh/go-homedir"
	"github.com/rushikesh2/Gophercises/CLI/todoz/store"
)

var hdir, _ = home.Dir()

var path = filepath.Join(hdir, "todo.db")

func TestAdd(t *testing.T) {
	store.Init(path)
	args := []string{"Add", "New", "value"}
	a := []string{}
	Addnewtask.Run(Addnewtask, args)
	store.DbCon.Close()
	//store.Init("/")
	Addnewtask.Run(Addnewtask, a)
}
