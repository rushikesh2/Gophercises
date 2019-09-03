package encrypt

import (
	"crypto/cipher"
	"errors"
	"fmt"
	"io"
	"testing"
)

var key = "This is key"
var text = "This is my secret test"
var prevEnc = "2ef309bb9f2c2f06a46b109858872df4ef38a199a9e2b57d6e23f191cb315499ba1ac19857fa"

//

func TestEncDec(test *testing.T) {
	encText, err := Enc(key, text)
	if err != nil {
		test.Error(err)
	}
	fmt.Println(encText)
	var plaintext string
	plaintext, err = Dec(key, encText)
	fmt.Println(plaintext)
	if err != nil {
		test.Error(err)
	}
	//fmt.Println(plaintext)
	if plaintext != text {
		fmt.Printf("Failed.Expeced %s got %s", text, plaintext)
		//errors.New(fmt.printf("Failed.Expeced %s got %s", text, plaintext))
	}
}

func TestEn(t *testing.T) {
	tedef := Cip
	defer func() {
		Cip = tedef
	}()

	Cip = func(key string) (cipher.Block, error) {
		return nil, errors.New("Error")
	}
	Enc("hi", "wel")
}

func TestEnc(t *testing.T) {
	tedef := EncR
	defer func() {
		EncR = tedef
	}()

	EncR = func(r io.Reader, buf []byte) (n int, err error) {
		return 0, errors.New("Error")
	}
	Enc("hi", "wel")
}

func TestDn(t *testing.T) {
	tedef := Cip
	defer func() {
		Cip = tedef
	}()

	Cip = func(key string) (cipher.Block, error) {
		return nil, errors.New("Error")
	}
	Dec("hi", "wel")
}

func TestDec(t *testing.T) {
	tedef := decS
	defer func() {
		decS = tedef
	}()

	decS = func(s string) ([]byte, error) {
		return nil, errors.New("Error")
	}
	Dec("hi", "wel")
}

func TestDecc(t *testing.T) {
	tedef := decS
	defer func() {
		decS = tedef
	}()

	decS = func(s string) ([]byte, error) {
		return nil, errors.New("Error")
	}
	Dec("", "")
}
