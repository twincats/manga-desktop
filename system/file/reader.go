package file

import (
	"io/fs"
	"os"
	"path/filepath"

	"github.com/climech/naturalsort"
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

/*
	FindFileByExt get list of files by extention
	first root is path to search, second ext is
	array of string extention
*/
func (f *FilesDataReader) FindFileByExt(root string, ext []string) []string {
	var a []string
	filepath.WalkDir(root, func(s string, d fs.DirEntry, e error) error {
		if e != nil {
			return e
		}
		for _, i := range ext {
			if filepath.Ext(d.Name()) == i {
				a = append(a, s)
			}
		}
		return nil
	})
	naturalsort.Sort(a)
	return a
}
