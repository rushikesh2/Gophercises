package main
import (
	"testing"
	"errors"
)


func TestMain(t *testing.T){
	HandleError(errors.New("dummy error"))
	main()
	
}