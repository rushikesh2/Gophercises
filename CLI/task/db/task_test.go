package db

import (
	"fmt"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

func startDB() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	Init(dbPath)

}

func TestInit(t *testing.T) {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := Init(dbPath)
	if err != nil {
		t.Error("expected nit got", err)
	}
}

func TestCreateTask(t *testing.T) {
	_, err := CreateTask("Task create")
	if err != nil {
		t.Error("expected nil", err)
	}
}

func TestAllTask(t *testing.T) {
	//var task []string
	_, err1 := AllTasks()
	if err1 != nil {
		t.Error("Expected tasks got,", err1)
	}
}

func TestDeleteTask(t *testing.T) {
	err := DeleteTask(1)
	fmt.Println(err)
	if err != nil {
		t.Error("expected nil got", err)
	}
}

func TestIntToByte(t *testing.T) {
	bytes := itob(7)
	fmt.Println(bytes)
	if bytes == nil {
		t.Error("Expected got nil", nil)
	}
}

func TestBtoi(t *testing.T) {
	var result []byte
	result = make([]byte, 8, 10)
	result = []byte{0, 0, 0, 0, 0, 0, 0, 10}
	val := btoi(result)
	if val == 0 {
		t.Error("Expected int got", val)
	}
}
