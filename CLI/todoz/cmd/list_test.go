package cmd

import (
	"errors"
	"path/filepath"
	"testing"

	home "github.com/mitchellh/go-homedir"
	"github.com/rushikesh2/Gophercises/CLI/todoz/store"
)

var Temp = ShowAll
var Temp2 = TempRemove

func TestDoneCmd(t *testing.T) {
	hdir, _ := home.Dir()

	path := filepath.Join(hdir, "todo.db")
	store.Init(path)
	valid_args := []string{"99", "2", "3"}
	invalid_args := []string{"1", "h"}
	doneCmd.Run(doneCmd, valid_args)
	doneCmd.Run(doneCmd, invalid_args)
	defer func() {
		ShowAll = Temp
		TempRemove = Temp2
	}()

	TempRemove = func(i int) error {
		return errors.New("Done")
	}
	doneCmd.Run(doneCmd, valid_args)

	ShowAll = func() ([]store.Todoz, error) {
		return nil, errors.New("error")
	}
	doneCmd.Run(doneCmd, valid_args)
}
func TestList(t *testing.T) {
	//store.Init(path)
	arr := []string{"Hello", "hi"}
	listCmd.Run(listCmd, arr)
	store.DbCon.Close()
	store.Init("dummy")
	listCmd.Run(listCmd, arr)

}
func TestListNegative(t *testing.T) {

	defer func() {
		ShowAll = Temp

	}()

	ShowAll = func() ([]store.Todoz, error) {
		return nil, errors.New("error")
	}
	s := []string{}
	listCmd.Run(listCmd, s)
}

// func TestDone(t *testing.T) {
// 	var tempda = Temp2
// 	defer func() {
// 		Temp2 = tempda
// 	}()
// 	Temp2 = func(id int) error {
// 		return errors.New("error")
// 	}
// 	s := []string{}
// 	doneCmd.Run(doneCmd, s)
// }
