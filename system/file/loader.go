package file

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

/*
FileLoader struct is custom asset handler for image manga
with base as "file/{manga_path}"
*/
type FileLoader struct {
	http.Handler
}

/*
NewFileLoader create instance of FileLoader
*/
func NewFileLoader() *FileLoader {
	return &FileLoader{}
}

/*
ServeHTTP response to asset handle of manga
*/
func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {

	requestedFilename := strings.TrimPrefix(req.URL.Path, "/")
	if strings.Split(requestedFilename, "/")[0] == "file" {
		requestedFilename = MANGA_PATH + requestedFilename[4:]
		// println("Requesting file:", requestedFilename)
		fileData, err := os.ReadFile(requestedFilename)

		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			res.Write([]byte(fmt.Sprintf("Could not load file %s", requestedFilename)))
		}

		res.Write(fileData)
	}

}
