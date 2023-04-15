package download

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (f *Download) GetChapter(o types.Option) (types.Chapter, error) {
	if d := f.getDownloads(o.ServerName); d != nil {
		//need database check
		//change title base on db manga table
		//change status chapter base on db chapter table
		chap, err := d.GetChapter(o)
		chaps := CheckChapterDB(*chap)
		return chaps, err
	}
	return types.Chapter{}, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetPage(o types.Option) (types.Page, error) {
	// option change to []types.ChapterList
	var nilpage types.Page
	if d := f.getDownloads(o.ServerName); d != nil {
		//simulation download
		pages_data, err := d.GetPage(o)
		if err != nil {
			return nilpage, err
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
		return *pages_data, err
	}
	return nilpage, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetChapterMdexPagination(url string, limit, offset int) ([]types.ChapterList, error) {
	mdex := new(types.Mangadex)
	return mdex.GetChapterMdexPagination(url, limit, offset)
}

type DownloadChannel struct {
	Index int    `json:"index"`
	Url   string `json:"Url"`
}

// check dbChapter
// 1. check mangaId
// 2. chapter is already downloaded
type CheckTitleDB struct {
	Id    uint
	Title string
}

type CheckChapDB struct {
	ID      uint    `json:"id"`
	Chapter float32 `json:"chapter"`
	Title   string  `json:"title"`
}

func CheckChapterDB(c types.Chapter) types.Chapter {
	var checkTitleDB CheckTitleDB
	ilike := fmt.Sprintf("%%%s%%", c.Manga)
	app.DB.Table("mangas").Select("mangas.id", "mangas.title").
		Joins("left join manga_alternatives on manga_alternatives.manga_id = mangas.id").
		Where("mangas.title LIKE ?", ilike).
		Or("manga_alternatives.title LIKE ?", ilike).Scan(&checkTitleDB)

	c.MangaId = checkTitleDB.Id

	if checkTitleDB.Id > 0 {
		var checkChapDb []CheckChapDB
		app.DB.Table("chapters").
			Select("id", "chapter", "title").
			Where("manga_id = ?", checkTitleDB.Id).Scan(&checkChapDb)
		for i, cp := range c.Chapter {
			for _, v := range checkChapDb {
				if cp.Chapter == json.Number(fmt.Sprintf("%v", v.Chapter)) {
					c.Chapter[i].Status = true
				}
			}
		}
	}
	return c
}
