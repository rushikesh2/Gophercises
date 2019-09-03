package main

import (
	"testing"
)

func TestMainFunc(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("error occured while calling function")
		}
	}()
	main()
}
