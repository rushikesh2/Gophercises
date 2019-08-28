package cmd

import (
	"path/filepath"
	"testing"

	"github.com/rushikesh2/GolangTraining/Gophercises/CLI/task/db"

	homedir "github.com/mitchellh/go-homedir"
)

func startDB() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	db.Init(dbPath)
}

func TestAddCommand(t *testing.T) {
	startDB()
	// file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	// oldStdout := os.Stdout
	// os.Stdout = file
	// a := []string{"Watch Golang tutorial"}
	// addCmd.Run(addCmd, a)
	// file.Seek(0, 0)
	// content, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	t.Error("error occured while test case : ", err)
	// }
	// output := string(content)
	// val := strings.Contains(output, "Added")
	// assert.Equalf(t, true, val, "they should be equal")
	// file.Truncate(0)
	// file.Seek(0, 0)
	// os.Stdout = oldStdout
	// file.Close()
	//dbconnect.Close()

	args := []string{"Watch", "go"}
	a := []string{}
	addCmd.Run(addCmd, args)
	addCmd.Run(addCmd, args)
	addCmd.Run(addCmd, args)
	db.Db.Close()
	addCmd.Run(addCmd, a)
	//db.Close()
}

//var AddT = AddTask

// func TestAddNegative(t *testing.T) {
// 	//startDB()
// 	s := []string{"add", "New", "volume"}
// 	addCmd.Run(addCmd, s)

// 	defer func() {
// 		AddTask = AddT
// 	}()

// 	AddTask = func(task string) (int, error) {
// 		//return task, error.New(error)
// 		return 0, errors.New("error")
// 	}
// 	addCmd.Run(addCmd, s)
// }
