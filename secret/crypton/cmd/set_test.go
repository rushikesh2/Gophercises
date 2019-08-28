package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing of the root command

func TestSet(t *testing.T) {
	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
	oldStdout := os.Stdout
	os.Stdout = file
	a := []string{"encd", "encr", "plain"}
	Setcmd.Run(Setcmd, a)
	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Error("error occured while test case : ", err)
	}
	output := string(content)
	val := strings.Contains(output, "Data successfully encoded.")
	assert.Equalf(t, true, val, "they should be equal")
	file.Truncate(0)
	file.Seek(0, 0)
	os.Stdout = oldStdout
	fmt.Println(string(content))
	file.Close()
}

// func TestSetNegative(t *testing.T) {
// 	file, _ := os.OpenFile("testing.txt", os.O_CREATE|os.O_RDWR, 0666)
// 	oldStdout := os.Stdout
// 	os.Stdout = file
// 	//encodingKey = "1234"
// 	a := []string{" ", " ", " "}
// 	Setcmd.Run(Setcmd, a)
// 	file.Seek(0, 0)
// 	content, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		t.Error("error occured while testing : ", err)
// 	}
// 	output := string(content)
// 	val := strings.Contains(output, "")
// 	assert.Equalf(t, true, val, "should be equal")
// 	file.Truncate(0)
// 	file.Seek(0, 0)
// 	os.Stdout = oldStdout
// 	fmt.Println(string(content))
// 	file.Close()
// }
