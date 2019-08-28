package encrypt

import (
	"os"
	"path/filepath"
	"testing"

	homedir "github.com/mitchellh/go-homedir"
)

type testStruct struct {
	key, val string
}

var testcase = []testStruct{
	{"key1", "Dummy string1"},
	{"key2", "dummy string 2"},
	{"key3", "Tesitng vaoflasmfms"},
}
var dummyvault = NewVault("It says this is not supposed to be smallaaaaaaaaa", "testing.txt")

func secretpath() string {
	home, _ := homedir.Dir()
	return filepath.Join(home, "test.txt")
}

// func TestSetAndGet(test *testing.T) {
// 	for _, values := range testcase {
// 		err := dummyvault.Set(values.key, values.val)
// 		if err != nil {
// 			test.Error(err)
// 		}

// 	}
// 	for _, values := range testcase {
// 		val, err := dummyvault.Get(values.key)
// 		if err != nil {
// 			test.Error(err)
// 			break
// 		}
// 		if val != values.val {
// 			fmt.Printf("Expected %s and got %s \n", values.val, val)
// 			test.Error("Test failed ")
// 		}
// 	}

// }

func TestSet(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	err := vault.Set("hello", "testing")
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestSetNegative(t *testing.T) {
	file := secretpath()
	vault := NewVault("", file)
	err := vault.Set("demo", "testing")
	if err == nil {
		t.Error("Expected  Error but got nil")
	}
}

func TestGet(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	_, err := vault.Get("hello")
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestGetNegative(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	_, err := vault.Get("abc")
	if err == nil {
		t.Error("Expected Error but got nil")
	}
	vault = NewVault("", file)
	_, err = vault.Get("abc")
	if err == nil {
		t.Error("Expected Error but got nil ")
	}
}

func TestLoad(t *testing.T) {
	file := secretpath()
	vault := NewVault("demo", file)
	err := vault.LoadMapping()
	if err != nil {
		t.Error("Expected nil but got", err)
	}
}

func TestLoadNegative(t *testing.T) {
	home, _ := homedir.Dir()
	file := filepath.Join(home, "")
	vault := NewVault("abc", file)
	err := vault.LoadMapping()
	if err == nil {
		t.Error("Expected error but got nil", err)
	}
	os.Remove(file)
}
func TestSave(t *testing.T) {
	var v Vault
	err := v.savemapping()
	if err == nil {
		t.Error("Expected error but got nil ")
	}
	deleteFile()
}

func deleteFile() {
	file := secretpath()
	os.Remove(file)
}
