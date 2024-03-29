package file

import (
	"os"
)

/*
	FileReader is struct for accessing file data related
*/
type FilesDataReader struct {
	MangaPath string
}

/*
	NewFileReader create new instance of FileReader
*/
func NewFilesDataReader() *FilesDataReader {
	return &FilesDataReader{}
}

/*
	GetFileManga get list of directory where MANGA_PATH is located
	currently is set with global file package variable MANGA_PATH
*/
func (f *FilesDataReader) GetFileManga() []string {
	dirs, _ := os.ReadDir(MANGA_PATH)
	var manga []string
	for _, i := range dirs {
		manga = append(manga, i.Name())
	}
	return manga
}
