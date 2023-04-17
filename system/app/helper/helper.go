package helper

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"regexp"
	"sync"

	"github.com/climech/naturalsort"
	vips "github.com/twincats/golibvips/libvips"
)

// ParallelCode is multiple code in concurrency run as parrarel with slice data
func ParallelCode[T any](p int, val []T, f func(v T)) {
	var wg sync.WaitGroup
	wg.Add(p)

	channel := make(chan T)
	for i := 0; i < p; i++ {
		go func() {
			for {
				v, more := <-channel
				if !more {
					wg.Done()
					return
				}

				f(v)
			}
		}()
	}

	for _, v := range val {
		channel <- v
	}

	close(channel)
	wg.Wait()
}

// Contains return boolean if T is included in slice []T
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

// AutoCrop crop manga image to desired width with contstrait
func AutoCrop(image *vips.ImageRef, w int) bool {
	width := image.Width()
	height := image.Height()

	ratio := float32(width) / float32(height)

	if width > w {
		if ratio > 1 {
			if w < 1980 {
				w = 1980
			}
		} else {
			if w < 1000 {
				w = 1000
			}
		}
		err := image.ResizeWidthPixel(float64(w), vips.KernelMitchell)
		if err != nil {
			fmt.Println(err)
			return false
		}
		return true
	}
	return false
}

// ListFiles show all image with included extenstion in dir path
func ListFiles(dir string, exts []string) []string {
	var res []string
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if Contains(exts, filepath.Ext(d.Name())) && d.Name() != "cover.webp" {
			res = append(res, path)
		}
		return nil
	})
	naturalsort.Sort(res)
	return res
}

func StringReplace(regexString string, match string, replaceString string) string {
	reg := regexp.MustCompile(regexString)
	return reg.ReplaceAllString(match, replaceString)
}

func FixMangaTitle(title string) string {
	mTitle := StringReplace(`(?m)[^\w\-[\]()' ~.,!@&]|\.+$`, title, "")
	mTitle = StringReplace("(?m) $|^ ", mTitle, "")
	return mTitle
}

// incase needed round by precision
// func RoundFloat(val float64, precision uint) float64 {
// 	ratio := math.Pow(10, float64(precision))
// 	return math.Round(val*ratio) / ratio
// }
