package main

import (
	"errors"
	"testing"
)

func TestMainFunc(t *testing.T) {
	HandleError(errors.New("dummy error"))
	main()

}
