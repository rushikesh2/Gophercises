package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			t.Errorf("error occured while calling function")
		}
	}()
	main()
}
