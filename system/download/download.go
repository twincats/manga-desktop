package download

import (
	"errors"
	"mangav4/system/app"
	"mangav4/system/download/types"
	"reflect"
	"sync"
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
		//need database check
		//change title base on db manga table
		//change status chapter base on db chapter table
		return d.GetChapter(o)
	}
	return nil, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetPage(o types.Option) (*types.Page, error) {
	// option change to []types.ChapterList
	if d := f.getDownloads(o.ServerName); d != nil {
		//simulation download
		pages_data, err := d.GetPage(o)
		if err != nil {
			return nil, err
		}
		//
		var parallel = 2
		var wg sync.WaitGroup
		wg.Add(parallel)

		channel := make(chan DownloadChannel)
		//loop parrarel
		for ii := 0; ii < parallel; ii++ {
			go func() {
				for {
					v, more := <-channel
					if !more {
						wg.Done()
						return
					}
					//testing long download time 1 second
					//place to download image one by one
					time.Sleep(time.Second)
					runtime.EventsEmit(*app.WailsContext, "dlProgress", v)
					//ouput ok needed :index, chapter, totalpages, mangaTitle
					//ouput err global failed url
				}
			}()
		}

		for i, p := range pages_data.Pages {
			channel <- DownloadChannel{Index: i + 1, Url: p}
		}

		close(channel)
		wg.Wait()
		// return data of all finished download
		// save to db if no err
		// global failed url return failed
		return pages_data, err
	}
	return nil, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetChapterMdexPagination(url string, limit, offset int) ([]types.ChapterList, error) {
	mdex := new(types.Mangadex)
	return mdex.GetChapterMdexPagination(url, limit, offset)
}

type DownloadChannel struct {
	Index int    `json:"index"`
	Url   string `json:"Url"`
}
