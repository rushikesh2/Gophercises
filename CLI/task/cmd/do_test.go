package cmd

import (
	"errors"
	"testing"

	"github.com/rushikesh2/GolangTraining/Gophercises/CLI/task/db"
)

//var showAll = All

var Remove = RemoveTask
var TaskL = TaskList

func TestDoR(t *testing.T) {
	startDB()
	input1 := []string{"1", "2"}
	defer func() {
		RemoveTask = Remove
	}()

	RemoveTask = func(i int) error {
		return errors.New("Done")
	}
	doCmd.Run(doCmd, input1)
	db.Db.Close()
}
func TestDo(t *testing.T) {
	startDB()
	input1 := []string{"1", "2", "3"}
	input2 := []string{"1", "a"}
	doCmd.Run(doCmd, input1)
	doCmd.Run(doCmd, input2)
	db.Db.Close()
}

func TestDoNegative(t *testing.T) {
	startDB()
	input1 := []string{"1", "a"}
	//doCmd.Run(doCmd, input1)

	defer func() {
		TaskList = TaskL
	}()
	//doCmd.Run(doCmd, input1)

	TaskList = func() ([]db.Task, error) {
		return nil, errors.New("error")
	}
	doCmd.Run(doCmd, input1)
	db.Db.Close()
}
