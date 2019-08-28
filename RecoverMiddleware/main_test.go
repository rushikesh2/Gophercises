package main

import (
	"net/http"
	"net/http/httptest"
	"runtime/debug"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

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
			url:          "line=24&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/GolangTraining/Gophercises/RecoverMIddleware/main.go",
			status:       200,
		}, {
			testCaseName: "TC2",
			url:          "line=ewr&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/GolangTraining/Gophercises/RecoverMIddleware/main.go",
			status:       200,
		}, {
			testCaseName: "TC3",
			url:          "line=24&path=C:/Users/gs-2019/gocode/src/github.com/rushikesh2/GolangTraining/Gophercises/RecoverMIddleware/main_test.go",
			status:       200,
		},
	}
	for i := 0; i < len(testTable); i++ {
		req, err := http.NewRequest("GET", "http://localhost:5000/debug?"+testTable[i].url, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		resr := httptest.NewRecorder()
		sourceCodeHandler(resr, req)
		res := resr.Result()
		if res.StatusCode != testTable[i].status {
			t.Errorf("Test case Number: %v Expected %v , Actual status %v", testTable[i].testCaseName, testTable[i].status, res.StatusCode)
		}
	}
}

func TestSource(t *testing.T) {

}
