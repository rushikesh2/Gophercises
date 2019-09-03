package main

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/rushikesh2/Gophercises/image/primitive"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	html := `<html><body>
			<form action="/upload" method="post" enctype="multipart/form-data">
				<input type="file" name="image">
				<button type="submit">Upload Image</button>
			</form>
			</body></html>`
	fmt.Fprint(w, html)
}

func modifyHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("file created")
	f, err := os.Open("./img/" + filepath.Base(r.URL.Path))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer f.Close()
	ext := filepath.Ext(f.Name())[1:]
	modeStr := r.FormValue("mode")
	if modeStr == "" {
		renderModeChoices(w, r, f, ext)
		return
	}
	mode, err := strconv.Atoi(modeStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	nStr := r.FormValue("n")
	if nStr == "" {
		renderNumShapeChoices(w, r, f, ext, primitive.Mode(mode))
		return
	}
	numShapes, err := strconv.Atoi(nStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_ = numShapes
	http.Redirect(w, r, "/img/"+filepath.Base(f.Name()), http.StatusFound)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	ext := filepath.Ext(header.Filename)[1:]
	onDisk, err := tempfile("", ext)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	defer onDisk.Close()
	_, err = io.Copy(onDisk, file)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	fmt.Println("file created")
	http.Redirect(w, r, "/modify/"+filepath.Base(onDisk.Name()), http.StatusFound)

}

func renderNumShapeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string, mode primitive.Mode) {
	opts := []genOpts{
		{N: 10, M: mode},
		{N: 20, M: mode},
		{N: 30, M: mode},
		{N: 40, M: mode},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}&n={{.NumShapes}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name      string
		Mode      primitive.Mode
		NumShapes int
	}
	var data []dataStruct
	for i, img := range imgs {
		data = append(data, dataStruct{
			Name:      filepath.Base(img),
			Mode:      opts[i].M,
			NumShapes: opts[i].N,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

func renderModeChoices(w http.ResponseWriter, r *http.Request, rs io.ReadSeeker, ext string) {
	opts := []genOpts{
		{N: 10, M: primitive.ModeCircle},
		{N: 10, M: primitive.ModeBeziers},
		{N: 10, M: primitive.ModePolygon},
		{N: 10, M: primitive.ModeCombo},
	}
	imgs, err := genImages(rs, ext, opts...)
	if err != nil {
		panic(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	html := `<html><body>
			{{range .}}
				<a href="/modify/{{.Name}}?mode={{.Mode}}">
					<img style="width: 20%;" src="/img/{{.Name}}">
				</a>
			{{end}}
			</body></html>`
	tpl := template.Must(template.New("").Parse(html))
	type dataStruct struct {
		Name string
		Mode primitive.Mode
	}
	var data []dataStruct
	for i, img := range imgs {
		data = append(data, dataStruct{
			Name: filepath.Base(img),
			Mode: opts[i].M,
		})
	}
	err = tpl.Execute(w, data)
	if err != nil {
		panic(err)
	}
}

type genOpts struct {
	N int
	M primitive.Mode
}

func genImages(rs io.ReadSeeker, ext string, opts ...genOpts) ([]string, error) {
	var ret []string
	for _, opt := range opts {
		rs.Seek(0, 0)
		f, err := genImage(rs, ext, opt.N, opt.M)
		if err != nil {
			return nil, err
		}
		ret = append(ret, f)
	}
	return ret, nil
}

func genImage(r io.Reader, ext string, numShapes int, mode primitive.Mode) (string, error) {
	out, err := primitive.Transform(r, ext, numShapes, primitive.WithMode(mode))
	if err != nil {
		return "", err
	}
	outFile, err := tempfile("", ext)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	io.Copy(outFile, out)
	return outFile.Name(), nil
}

var temp = ioutil.TempFile

func tempfile(prefix, ext string) (*os.File, error) {
	in, err := temp("./img/", prefix)
	if err != nil {
		return nil, errors.New("main: failed to create temporary file")
	}
	defer os.Remove(in.Name())
	return os.Create(fmt.Sprintf("%s.%s", in.Name(), ext))
}

func getHandlers() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/modify/", modifyHandler)
	mux.HandleFunc("/upload", uploadHandler)
	fs := http.FileServer(http.Dir("./img/"))
	mux.Handle("/img/", http.StripPrefix("/img", fs))
	return mux
}
func main() {
	http.ListenAndServe(":5000", getHandlers())
}
