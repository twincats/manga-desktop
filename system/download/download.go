package download

import (
	"errors"
	"mangav4/system/app"
	"mangav4/system/download/types"
	"reflect"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Download struct {
	types.Downloads
}

// Get new Download instance
func NewDownload() *Download {
	return new(Download)
}

func (f *Download) getDownloads(className string) types.Downloads {
	var d types.Downloads
	for c := range types.CLASS {
		if className == c {
			ref := reflect.ValueOf(types.CLASS[c])
			d = ref.Interface().(types.Downloads)
		}
	}
	return d
}

func (f *Download) GetChapter(o types.Option) (*types.Chapter, error) {
	if d := f.getDownloads(o.ServerName); d != nil {
		return d.GetChapter(o)
	}
	return nil, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetPage(o types.Option) (*types.Page, error) {
	if d := f.getDownloads(o.ServerName); d != nil {
		//simulation download
		pages_data, err := d.GetPage(o)
		if err != nil {
			return nil, err
		}
		//
		for _, p := range pages_data.Pages {
			time.Sleep(time.Second)
			runtime.EventsEmit(*app.WailsContext, "dlProgress", p)
		}
		return pages_data, err
	}
	return nil, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetChapterMdexPagination(url string, limit, offset int) ([]types.ChapterList, error) {
	mdex := new(types.Mangadex)
	return mdex.GetChapterMdexPagination(url, limit, offset)
}
