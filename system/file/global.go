package file

import (
	"io"
	"io/fs"
	"mangav4/system/app/helper"
	"os"
	"path/filepath"
)

// global path for base manga_path
var MANGA_PATH string

func ReadDir(path string) []string {
	dirs, _ := os.ReadDir(path)
	var listdir []string
	for _, i := range dirs {
		if i.IsDir() {
			listdir = append(listdir, i.Name())
		}

	}
	return listdir
}

func FirstImageinDir(searchpath string) string {
	var imgpath string
	filepath.WalkDir(searchpath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if helper.Contains([]string{".png", ".jpg", ".webp"}, filepath.Ext(d.Name())) && !d.IsDir() {
			imgpath = path
			return io.EOF
		}
		return nil
	})
	return imgpath
}
