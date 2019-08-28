package cmd

import (
	"errors"
	"testing"

	"github.com/rushikesh2/GolangTraining/Gophercises/CLI/task/db"
)

var name = db.Dbname

func TestList(t *testing.T) {
	startDB()
	args := []string{}
	listCmd.Run(listCmd, args)
	defer func() {
		db.Dbname = name
	}()
	db.Dbname = "todo"
	listCmd.Run(listCmd, args)
	db.Db.Close()

}

var show = Tasklist

func TestListNegative(t *testing.T) {
	startDB()
	s := []string{"hi"}
	listCmd.Run(listCmd, s)

	defer func() {
		Tasklist = show
	}()
	Tasklist = func() ([]db.Task, error) {
		return nil, errors.New("error")
	}
	listCmd.Run(listCmd, s)
	db.Db.Close()
}
