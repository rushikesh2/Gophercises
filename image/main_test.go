package main

import (
	"errors"
	"io"

	"os"
	"testing"

	"github.com/rushikesh2/Gophercises/image/primitive"
)

// func TestUpload(t *testing.T) {
// 	testServer := httptest.NewServer(getHandlers())
// 	defer testServer.Close()

// 	newreq := func(method, url string) *http.Request {
// 		file, err := os.Open("./input.png")
// 		if err != nil {
// 			t.Error("Expected nil got", err)
// 		}
// 		body := &bytes.Buffer{}
// 		writer := multipart.NewWriter(body)
// 		part, err := writer.CreateFormFile("image", file.Name())
// 		if err != nil {
// 			t.Error("Expected nil got", err)
// 		}
// 		_, err = io.Copy(part, file)
// 		if err != nil {
// 			t.Error("Expected nil got", err)
// 		}
// 		err = writer.Close()
// 		if err != nil {
// 			t.Error("Expected nil got", err)
// 		}
// 		req, err := http.NewRequest(method, url, body)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		req.Header.Set("Content-Type", writer.FormDataContentType())
// 		return req
// 	}

// 	testCases := []struct {
// 		name   string
// 		req    *http.Request
// 		status int
// 	}{
// 		{name: "TC2", req: newreq("POST", testServer.URL+"/upload"), status: 200},
// 	}
// 	for _, tests := range testCases {
// 		t.Run(tests.name, func(t *testing.T) {
// 			resp, err := http.DefaultClient.Do(tests.req)
// 			if err != nil {
// 				t.Fatal(err)
// 			}
// 			defer resp.Body.Close()
// 			if resp.StatusCode != tests.status {
// 				t.Error("Expected 200 got", tests.status)
// 			}
// 		})
// 	}
// }

// func TestModify(t *testing.T) {
// 	testServer := httptest.NewServer(getHandlers())
// 	defer testServer.Close()

// 	newreq := func(method, url string, body io.Reader) *http.Request {
// 		req, err := http.NewRequest(method, url, body)
// 		if err != nil {
// 			t.Fatal(err)
// 		}
// 		return req
// 	}

// 	testCases := []struct {
// 		name   string
// 		req    *http.Request
// 		status int
// 	}{
// 		{name: "TC3", req: newreq("GET", testServer.URL+"/modify/img/690011768.png?mode=2", nil), status: 200},
// 		{name: "TC4", req: newreq("GET", testServer.URL+"/modify/img/011501652.png?mode=2&number=10", nil), status: 200},
// 	}
// 	for _, tests := range testCases {
// 		t.Run(tests.name, func(t *testing.T) {
// 			resp, err := http.DefaultClient.Do(tests.req)
// 			if err != nil {
// 				t.Error("Expected nil got", err)
// 			}
// 			defer resp.Body.Close()
// 			if resp.StatusCode != tests.status {
// 				t.Error("Expected 200 got", resp.StatusCode)
// 			}
// 		})
// 	}
// }
func TestGenimage(t *testing.T) {
	var image io.Reader
	image, _ = os.Open("input.png")

	_, err := genImage(image, ".png", 2, primitive.ModeTriangle)
	if err != nil {
		t.Error("Error to transform", err)
	}
}

func TestTempfile(t *testing.T) {
	_, err := tempfile("", "")
	if err != nil {
		t.Error("primitive: failed to create temporary file")
	}
}

func TestTemp(t *testing.T) {
	tedef := temp
	defer func() {
		temp = tedef
	}()

	temp = func(dir, pattern string) (f *os.File, err error) {
		return nil, errors.New("Error")
	}
	tempfile("", "")
}

func TestMain(m *testing.M) {
	main()
}
