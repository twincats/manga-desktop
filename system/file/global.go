package file

import (
	"os"
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
