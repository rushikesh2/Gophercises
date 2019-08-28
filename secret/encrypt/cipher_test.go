package encrypt

import (
	"fmt"
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
