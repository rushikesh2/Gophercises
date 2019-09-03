package main

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"testing"

	"github.com/alecthomas/chroma"
	"github.com/stretchr/testify/assert"
)

func TestMainFunc(t *testing.T) {

	tempListenAndServer := listenAndServerFunc
	defer func() {
		listenAndServerFunc = tempListenAndServer
	}()
	listenAndServerFunc = func(port string, handler http.Handler) error {
		panic("testing")
	}
	assert.PanicsWithValuef(t, "testing", main, "they should be equal")
}

func TestCreateLinks(t *testing.T) {
	stack := debug.Stack()
	link := createLinks(string(stack))
	if link == "" {
		t.Error("Expected link got", link)
	}
}

func TestLink(t *testing.T) {
	stack := debug.Stack()
	link := createLinks(string(stack))
	if link == "" {
		t.Error("Expected link got", link)
	}
}
func TestMiddleware(t *testing.T) {
	handler := http.HandlerFunc(panicDemo)
	executeRequest("Get", "/panic", middleware(handler))
}
func executeRequest(method string, url string, handler http.Handler) (*httptest.ResponseRecorder, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	resr := httptest.NewRecorder()
	handler.ServeHTTP(resr, req)
	return resr, err
}

func TestSourceHandler(t *testing.T) {
	testTable := []struct {
		testCaseName string
		url          string
		status       int
	}{
		{
			testCaseName: "TC1",
			url:          "line=24&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/Gophercises/RecoverMIddleware/main.go",
			status:       200,
		}, {
			testCaseName: "TC2",
			url:          "line=ewr&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/Gophercises/RecoverMIddleware/main.go",
			status:       500,
		}, {
			testCaseName: "TC3",
			url:          "line=24&path=/home/gslab/go/main.go",
			status:       500,
		},
	}
	for i := 0; i < len(testTable); i++ {
		req, err := http.NewRequest("GET", "http://localhost:5000/debug?"+testTable[i].url, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		resr := httptest.NewRecorder()
		sourceCodeHandler(resr, req)

		var cop = copyData
		defer func() {
			copyData = cop
		}()
		copyData = func(dst io.Writer, src io.Reader) (written int64, err error) {
			return 0, errors.New("error")
		}

		res := resr.Result()
		if res.StatusCode != testTable[i].status {
			t.Errorf("Test case Number: %v Expected %v , Actual status %v", testTable[i].testCaseName, testTable[i].status, res.StatusCode)
		}
	}
}

func TestSourceNegative(t *testing.T) {
	testTable := []struct {
		testCaseName string
		url          string
		status       int
	}{
		{
			testCaseName: "TC2",
			url:          "line=ewr&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/Gophercises/RecoverMIddleware/main.go",
			status:       200,
		}, {
			testCaseName: "TC1",
			url:          "line=24&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/Gophercises/RecoverMIddleware/main.go",
			status:       500,
		},
	}
	for i := 0; i < len(testTable); i++ {
		req, err := http.NewRequest("GET", "http://localhost:5000/debug?"+testTable[i].url, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		resr := httptest.NewRecorder()
		sourceCodeHandler(resr, req)
		var cop = copyData
		defer func() {
			copyData = cop
		}()
		copyData = func(dst io.Writer, src io.Reader) (written int64, err error) {
			//http.Error(w, err.Error(), http.StatusInternalServerError)
			return 0, errors.New("error")
		}
		res := resr.Result()
		if res.StatusCode != testTable[i].status {
			t.Errorf("Test case Number: %v Expected %v , Actual status %v", testTable[i].testCaseName, testTable[i].status, res.StatusCode)
		}
	}
}

func TestRemaining(t *testing.T) {

	var sty = styleH
	defer func() {
		styleH = sty
	}()
	styleH = func(name string) *chroma.Style {
		return nil
	}
	req, err := http.NewRequest("GET", "http://localhost:5000/debug?line=24&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/Gophercises/RecoverMIddleware/main.go", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	resr := httptest.NewRecorder()
	sourceCodeHandler(resr, req)
}
