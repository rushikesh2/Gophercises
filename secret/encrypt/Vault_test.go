package encrypt

import (
	"crypto/cipher"
	"errors"
	"io"
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

func TestLoadM(t *testing.T) {
	tedef := opFile
	defer func() {
		opFile = tedef
	}()

	opFile = func(name string) (*os.File, error) {
		return nil, errors.New("Error")
	}
	ekey := "text"
	v := NewVault(ekey, "/")
	v.LoadMapping()
}

func TestSaveM(t *testing.T) {
	tedef := EWrite
	defer func() {
		EWrite = tedef
	}()

	EWrite = func(key string, w io.Writer) (*cipher.StreamWriter, error) {
		return nil, errors.New("Error")
	}
	ekey := "text"
	fp := "abc"
	v := NewVault(ekey, fp)
	v.savemapping()
}
