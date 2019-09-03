package primitive

import (
	"errors"
	"io"
	"os"
	"testing"
)

func TestTranform(t *testing.T) {
	var image io.Reader
	image, _ = os.Open("input.png")
	image, err := Transform(image, ".png", 20, WithMode(ModeCircle))
	if err != nil {
		t.Error("Error to transform", err)
	}
}

func TestPrimitive(t *testing.T) {
	_, err := primitive("in.png", "out.png", 20, "abc.png")
	if err == nil {
		t.Error("Expected image got")
	}

}

func TestPri(t *testing.T) {
	tedef := tempD
	defer func() {
		tempD = tedef
	}()

	tempD = func(prefix, ext string) (*os.File, error) {
		return nil, errors.New("Error")
	}
	var image io.Reader
	image, _ = os.Open("input.png")
	Transform(image, ".png", 20, WithMode(ModeCircle))
}

func TestPrimit(t *testing.T) {
	tedef := tempDa
	defer func() {
		tempDa = tedef
	}()

	tempDa = func(prefix, ext string) (*os.File, error) {
		return nil, errors.New("Error")
	}
	var image io.Reader
	image, _ = os.Open("input.png")
	Transform(image, ".png", 20, WithMode(ModeCircle))
}

func TestPriCop(t *testing.T) {
	tedef := cop
	defer func() {
		cop = tedef
	}()

	cop = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, errors.New("Error")
	}
	var image io.Reader
	image, _ = os.Open("input.png")
	Transform(image, ".png", 20, WithMode(ModeCircle))
}

func TestPriCops(t *testing.T) {
	tedef := cops
	defer func() {
		cops = tedef
	}()

	cops = func(dst io.Writer, src io.Reader) (written int64, err error) {
		return 0, errors.New("Error")
	}
	var image io.Reader
	image, _ = os.Open("input.png")
	Transform(image, ".png", 20, WithMode(ModeCircle))
}

func TestPrim(t *testing.T) {
	tedef := prim
	defer func() {
		prim = tedef
	}()
	prim = func(inputFile, outputFile string, numShapes int, args ...string) (string, error) {
		return "img", errors.New("Error")
	}
	var image io.Reader
	image, _ = os.Open("input.png")
	Transform(image, ".png", 20, WithMode(ModeCircle))
}

func TestTempFile(t *testing.T) {
	_, err := tempfile("", "")
	if err != nil {
		t.Error("primitive: failed to create temporary file")
	}
}

func TestTem(t *testing.T) {
	tedef := temp
	defer func() {
		temp = tedef
	}()

	temp = func(dir, pattern string) (f *os.File, err error) {
		return nil, errors.New("Error")
	}
	tempfile("", "")
}
