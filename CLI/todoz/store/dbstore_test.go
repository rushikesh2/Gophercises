package store

import (
	"fmt"
	"path/filepath"
	"testing"

	home "github.com/mitchellh/go-homedir"
)

//var testvar *testing.T

type teststruct struct {
	input    int
	expected []byte
}

// test boti() and itob()

var hdir, _ = home.Dir()

var path = filepath.Join(hdir, "todo.db")
var testval = teststruct{
	input:    5,
	expected: []byte{0, 0, 0, 0, 0, 0, 0, 5},
}

func TestItob(t *testing.T) {
	for index, value := range itob(5) {
		if value != testval.expected[index] {

			t.Error("Itob failed")

		}
	}
}

func TestBtoi(t *testing.T) {
	if btoi(testval.expected) != testval.input {
		t.Error("Btoi failed")
	}
}

func TestInsertTask(t *testing.T) {
	Init(path)
	_, err := InsertTask("Dummy task")
	if err != nil {
		t.Error("Insert Failed")
		fmt.Println(err)
	}
}

func TestRemoveTasks(t *testing.T) {
	err := RemoveTasks(5)
	if err != nil {
		t.Error("Delete failed")
	}
}

func TestGetAll(t *testing.T) {
	_, err := GetAll()
	if err != nil {
		t.Error("Get all failed")
	}
}

func TestInit(t *testing.T) {
	err := Init("/")
	if err != nil {
		t.Error("Failed init..")
	}
}
