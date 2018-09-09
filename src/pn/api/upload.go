package api

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
)

func (a *API) UploadHandler(w http.ResponseWriter, r *http.Request) {
	a.Lock()
	if a.currentFilepath != "" {
		fmt.Printf("Removing current file in use: %v\n", a.currentFilepath)
		err := os.RemoveAll(a.currentFilepath)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		a.currentFilepath = ""
	}
	dir, err := ioutil.TempDir("", "")
	if err != nil {
		log.Fatal(err)
	}
	a.currentFilepath = dir
	a.Unlock()

	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	a.Lock()
	a.currentFilename = handler.Filename
	a.Unlock()

	fmt.Fprintf(w, "%v", handler.Header)

	f, err := os.OpenFile(path.Join(dir, handler.Filename), os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
