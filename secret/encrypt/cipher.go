package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
)

// returns cipher block based on key provided.
func getCipherBlock(key string) (cipher.Block, error) {

	hash := md5.New()
	fmt.Fprint(hash, key)
	cipherk := hash.Sum(nil)
	return aes.NewCipher(cipherk) //function requies fixed sized

}

func GetEncryptStream(key string, iv []byte) (cipher.Stream, error) {
	var cipherKey cipher.Stream
	block, err := getCipherBlock(key)
	if err == nil {
		cipherKey = cipher.NewCFBEncrypter(block, iv)
	}
	return cipherKey, err
}

// EncWriter encodes and writes to given stream writer.
// func EncWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
// 	iv := make([]byte, aes.BlockSize)                       //acts as salt and is to be read/written first by StreamReader/Writer
// 	if _, err := io.ReadFull(rand.Reader, iv); err != nil { //reads random values into byte size
// 		return nil, err
// 	}
// 	stream, err := GetEncryptStream(key, iv)
// 	if err != nil {
// 		return nil, err
// 	}
// 	bytesWritten, err := w.Write(iv)
// 	if bytesWritten != len(iv) || err != nil {
// 		fmt.Println(err)
// 		return nil, errors.New("Iv cannot be written")
// 	}
// 	return &cipher.StreamWriter{
// 		S: stream, W: w,
// 	}, nil
// }
// EncWriter encodes and writes to given stream writer.
func EncWriter(key string, w io.Writer) (*cipher.StreamWriter, error) {
	var streamWriter *cipher.StreamWriter
	iv := make([]byte, aes.BlockSize) //acts as salt and is to be read/written first by StreamReader/Writer
	io.ReadFull(rand.Reader, iv)      //reads random values into byte size

	stream, _ := GetEncryptStream(key, iv)
	_, err := w.Write(iv)
	if err == nil {
		streamWriter = &cipher.StreamWriter{S: stream, W: w}
	}
	return streamWriter, err

}

func GetDecryptStream(key string, iv []byte) (cipher.Stream, error) {
	var cipherKey cipher.Stream
	block, err := getCipherBlock(key)
	if err == nil {

		cipherKey = cipher.NewCFBDecrypter(block, iv)
	}
	return cipherKey, err
}

func DecReader(key string, r io.Reader) (*cipher.StreamReader, error) {
	iv := make([]byte, aes.BlockSize) //acts as salt and is to be read/written first by StreamReader/Writer
	n, err := r.Read(iv)
	if n != len(iv) || err != nil {
		return nil, errors.New(" EncReader:unable to read iv")
	}
	stream, err := GetDecryptStream(key, iv)
	// if err != nil {
	// 	return nil, err
	// }
	return &cipher.StreamReader{S: stream, R: r}, err

}

func Enc(key, text string) (string, error) {

	block, err := getCipherBlock(key)
	if err != nil {
		return "", err
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))

	// It's important to remember that ciphertexts must be authenticated
	// (i.e. by using crypto/hmac) as well as being encrypted in order to
	// be secure.
	encvalue := fmt.Sprintf("%x", ciphertext)

	return encvalue, nil
}

func Dec(key, encodedval string) (string, error) {

	block, err := getCipherBlock(key)
	if err != nil {
		//fmt.Println(err.Error() + "3")
		return "", err
	}
	ciphertext, err := hex.DecodeString(encodedval)
	if err != nil {
		//fmt.Println(err.Error() + "4")
		return "", err
	}
	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}
