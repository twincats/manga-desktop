package file

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"mangav4/system/app"
	"mangav4/system/app/helper"

	getSize "github.com/markthree/go-get-folder-size/src"
	vips "github.com/twincats/golibvips/libvips"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// NewConvert create new instance of NewChapter
func NewConvert() *Convert {
	return &Convert{}
}

// Convert object is used for main convert function
type Convert struct {
	Title       string   `json:"title"`
	Quality     int      `json:"quality"`
	Parallel    int      `json:"parallel"`
	ResizeWidth int      `json:"resize_width"`
	Resize      bool     `json:"resize"`
	Delete      bool     `json:"delete"`
	Engine      bool     `json:"engine"`
	Ext         []string `json:"ext"`
}

type ConvertResult struct {
	Manga       string  `json:"manga"`
	SizeBefore  int64   `json:"size_before"`
	SizeAfter   int64   `json:"size_after"`
	SizePercent float64 `json:"size_percent"`
	TotalOK     int     `json:"total_ok"`
	TotalFailed int     `json:"total_failed"`
	TotalResize int     `json:"total_resize"`
	TotalDelete int     `json:"total_delete"`
}

type StartConvertEventData struct {
	SizeBefore   int `json:"size_before"`
	TotalConvert int `json:"total_convert"`
}

type RunConvertEventData struct {
	Error    string `json:"error"`
	FilePath string `json:"filepath"`
	Filename string `json:"filename"`
	Chap     string `json:"chap"`
	Resized  bool   `json:"resized"`
	Deleted  bool   `json:"deleted"`
}

func (f *Convert) DoConvert(c Convert) (ConvertResult, error) {

	path := filepath.Join(MANGA_PATH, c.Title)
	ListImages := helper.ListFiles(path, c.Ext)

	if len(ListImages) < 1 {
		return ConvertResult{}, errors.New("no image to Convert")
	}

	var res ConvertResult
	res.Manga = c.Title
	bfSize, _ := getSize.Parallel(path)

	// emit Event once
	runtime.EventsEmit(*app.WailsContext, "start_convert",
		&StartConvertEventData{
			SizeBefore:   int(bfSize),
			TotalConvert: len(ListImages),
		})

	// setting parallel with default 5 / if not set
	parallel := c.Parallel
	if c.Parallel == 0 {
		parallel = 5
	}

	// runnung code in parrarel
	helper.ParallelCode(parallel, ListImages, func(v string) {

		fList := strings.Split(v, "\\")
		ext := filepath.Ext(v)
		var rc RunConvertEventData
		rc.FilePath = v

		if len(fList) > 0 {
			rc.Filename = fList[len(fList)-1]
			rc.Chap = fList[len(fList)-2]
		}

		image, err := vips.NewImageFromFile(v)
		// show failed error
		if err != nil {
			rc.Error += "error read image file. "
			fmt.Println(err, rc.Filename)
		}

		if c.Resize {
			if helper.AutoCropManga(image, c.ResizeWidth) {
				res.TotalResize += 1
				rc.Resized = true
			}
		}

		// export image to webp
		ep := vips.NewWebpExportParams()
		ep.StripMetadata = true
		ep.Quality = c.Quality

		webpByte, _, err := image.ExportWebp(ep)
		if err != nil {
			rc.Error += "error export image. "
			fmt.Println(err, rc.Filename)
		}

		// save exported byte to file
		err = os.WriteFile(v[:len(v)-len(ext)]+".webp", webpByte, 0644)
		if err != nil {
			rc.Error += "error save converted image. "
			fmt.Println(err, rc.Filename)
		}

		// delete original
		if c.Delete {
			// only delete if not converted to webp
			if ext != ".webp" {
				err = os.Remove(v)
				if err != nil {
					rc.Error += "error delete image. "
					fmt.Println(err, rc.Filename)
				} else {
					res.TotalDelete += 1
					rc.Deleted = true
				}
			}
		}

		// calculate failed convert
		if rc.Error != "" {
			res.TotalFailed += 1
		}

		// calculate OK convert
		fmt.Println("sukses : " + rc.Filename)
		res.TotalOK += 1

		//send RunConvertEventData w/ Event runtime
		runtime.EventsEmit(*app.WailsContext, "run_convert", rc)

	})

	afSize, _ := getSize.Parallel(path)

	res.SizeBefore = bfSize
	res.SizeAfter = afSize
	res.SizePercent = (float64(afSize) - float64(bfSize)) / float64(bfSize) * 100

	return res, nil
}
