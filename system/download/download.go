package download

import (
	"encoding/json"
	"errors"
	"fmt"
	"mangav4/system/app"
	"mangav4/system/app/helper"

	"mangav4/system/app/internet"
	"mangav4/system/download/types"
	"mangav4/system/file"
	"mangav4/system/manga"
	"os"
	"path/filepath"
	"reflect"

	vips "github.com/twincats/golibvips/libvips"
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
		var nilChap types.Chapter
		chap, err := d.GetChapter(o)
		if err != nil {

			return nilChap, err
		}
		chaps := f.CheckChapterDB(*chap)
		return chaps, nil
	}
	return types.Chapter{}, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetPage(o types.Chapter) (PageReport, error) {
	// option change to []types.ChapterList
	var pageReport PageReport
	if d := f.getDownloads(o.ServerName); d != nil {
		// setup download
		h := internet.NewFileDownloader()
		title := helper.FixMangaTitle(o.Manga)

		tx := app.DB.Begin()
		var mg manga.Manga
		if o.MangaId == 0 {
			mg.Title = o.Manga
			mg.Mdex = o.Mdex
			err := tx.Create(&mg).Error
			if err != nil {
				tx.Rollback()
				return pageReport, err
			}

			cvPath := filepath.Join(file.MANGA_PATH, title)
			h.SetupDirectory(cvPath)
			err = downloadCover(filepath.Join(cvPath, "cover.webp"), o.Cover)
			if err != nil {
				h.RemoveFolder(cvPath)
				tx.Rollback()
				return pageReport, err
			}
		}
		tx.SavePoint("manga")

		pageReport.Manga = o.Manga

		// batch download
		dl := internet.NewFileDownload(app.C)
		for ichap, chap := range o.Chapter {

			var failChap FiledChapReport
			failChap.ChapID = chap.ID
			failChap.Chapter = chap.Chapter

			chapRes, err := d.GetPage(types.Option{
				URL:        chap.ID,
				ServerName: o.ServerName,
				DataSaver:  &o.DataSaver,
			})

			if err != nil {
				failChap.Err = err
				pageReport.FailChap = append(pageReport.FailChap, failChap)
				continue
			}

			eventChap := EventChap{
				ListPage:  chapRes.Pages,
				IndexChap: ichap + 1,
				TotalChap: len(o.Chapter),
				TotalPage: len(chapRes.Pages),
				Chapter:   chap.Chapter,
			}
			runtime.EventsEmit(*app.WailsContext, "dl_eventchap", eventChap)

			dlPath := filepath.Join(file.MANGA_PATH, title, chap.Chapter.String())

			failed := dl.DownloadBatch(3, dlPath, chapRes.Pages, func(s internet.StatDownload) {
				if s.Err == nil {
					eventPage := EventPage{
						Images:       s.URL,
						IndexPage:    s.Index,
						StatError:    s.Err,
						Chapter:      chap.Chapter,
						StatusResume: s.StatusResume,
					}
					runtime.EventsEmit(*app.WailsContext, "dl_eventpage", eventPage)
				}
			})

			// auto retry 3 times
			failedRetry := dl.RetryDownloadsCounter(3, failed, func(s internet.StatDownload) {
				if s.Err == nil {
					eventPage := EventPage{
						Images:       s.URL,
						IndexPage:    s.Index,
						StatError:    s.Err,
						Chapter:      chap.Chapter,
						StatusResume: s.StatusResume,
					}
					runtime.EventsEmit(*app.WailsContext, "dl_eventpage", eventPage)
				}
			})

			// SAVE DB
			if len(failedRetry) < len(chapRes.Pages) {
				// save cover & manga
				var mcp manga.Chapter
				tx.Where("manga_id")
				mcpChap, _ := chap.Chapter.Float64()
				mcpLang := manga.LangByLang(chap.Languange)
				mcp.Chapter = float32(mcpChap)
				mcp.Volume = chap.Volume
				mcp.Group = chap.GroupName
				mcp.LanguageId = uint(mcpLang.ID)
				mcp.Language = mcpLang
				mcp.MangaId = o.MangaId

				if o.MangaId == 0 {
					mcp.MangaId = mg.ID
				}

				err := tx.Where(manga.Chapter{Chapter: mcp.Chapter, MangaId: mcp.MangaId}).FirstOrCreate(&mcp).Error
				if err == nil {
					failChap.StatusDB = true
				}

			} else {
				h.RemoveFolder(dlPath)
			}

			if len(failedRetry) > 0 || failChap.Err != nil {
				failChap.FailPage = failedRetry
				pageReport.FailChap = append(pageReport.FailChap, failChap)
			}

		} // END FOR

		// save cover for new manga
		err := tx.Commit().Error
		if err != nil {
			pageReport.Error = err.Error()
		}

		// return finish complete report
		return pageReport, nil
	}
	return pageReport, errors.New("error Server Name : " + o.ServerName + " Not Found or Implemented")
}

func (f *Download) GetChapterMdexPagination(url string, limit, offset int) ([]types.ChapterList, error) {
	mdex := new(types.Mangadex)
	return mdex.GetChapterMdexPagination(url, limit, offset)
}

// download from frontend
func (f *Download) DownloadJS(p ParamJS) []internet.StatDownload {
	dlPath := filepath.Join(file.MANGA_PATH, p.Manga, string(p.Chapter.Chapter))
	fd := internet.NewFileDownload(app.C)
	failed := fd.DownloadBatch(3, dlPath, p.Pages, func(s internet.StatDownload) {
		if s.Err == nil {
			eventPage := EventPage{
				Images:       s.URL,
				IndexPage:    s.Index,
				StatError:    s.Err,
				Chapter:      p.Chapter.Chapter,
				StatusResume: s.StatusResume,
			}
			runtime.EventsEmit(*app.WailsContext, "dl_eventpage", eventPage)
		}
	})

	// auto retry 3 times
	failedRetry := fd.RetryDownloadsCounter(3, failed, func(s internet.StatDownload) {
		if s.Err == nil {
			eventPage := EventPage{
				Images:       s.URL,
				IndexPage:    s.Index,
				StatError:    s.Err,
				Chapter:      p.Chapter.Chapter,
				StatusResume: s.StatusResume,
			}
			runtime.EventsEmit(*app.WailsContext, "dl_eventpage", eventPage)
		}
	})

	// SAvae chapter
	if len(failed) < len(p.Pages) {
		//save chapter
		tx := app.DB.Begin()
		tx.Where("manga_id")
		var mcp manga.Chapter
		mcpChap, _ := p.Chapter.Chapter.Float64()
		mcpLang := manga.LangByLang(p.Chapter.Languange)
		mcp.Chapter = float32(mcpChap)
		mcp.Volume = p.Chapter.Volume
		mcp.Group = p.Chapter.GroupName
		mcp.LanguageId = uint(mcpLang.ID)
		mcp.Language = mcpLang
		mcp.MangaId = p.MangaID

		tx.Where(manga.Chapter{Chapter: mcp.Chapter, MangaId: mcp.MangaId}).FirstOrCreate(&mcp)

		err := tx.Commit().Error
		if err != nil {
			fmt.Println("Error save chapter", err)
		}
		fmt.Println("sukses save")
	}

	// return value
	if len(failedRetry) > 0 {
		fmt.Println("some download failed after retry")
		return failedRetry
	} else {
		fmt.Println("sukses download")
		return []internet.StatDownload{}
	}
}

func (f *Download) SaveMangaCover(c types.Chapter) (uint, error) {
	if c.MangaId > 0 {
		fmt.Println("error already have mangaid")
		return 0, errors.New("error: Already Have MangaID")
	} else {
		var mg manga.Manga
		mg.Title = c.Manga
		mg.Mdex = c.Mdex
		tx := app.DB.Begin()
		err := tx.Create(&mg).Error
		if err != nil {
			tx.Rollback()
			fmt.Println("error save manga trx")
			return 0, err
		}

		dl := internet.NewFileDownloader()
		title := helper.FixMangaTitle(c.Manga)
		cvPath := filepath.Join(file.MANGA_PATH, title)
		dl.SetupDirectory(cvPath)
		err = downloadCover(filepath.Join(cvPath, "cover.webp"), c.Cover)
		if err != nil {
			dl.RemoveFolder(cvPath)
			tx.Rollback()
			fmt.Println("error save cover")
			return 0, err
		}

		err = tx.Commit().Error
		if err != nil {
			fmt.Println("error commit trx")
			return 0, err
		}

		return mg.ID, nil
	}
}

func (f *Download) RetryDownloads(count int, retryData []internet.StatDownload) ([]internet.StatDownload, error) {

	if count == 0 {
		return retryData, errors.New("retryDownload : count must be more than 0")
	}

	dl := internet.NewFileDownload(app.C)

	moreFailed := dl.RetryDownloadsCounter(count, retryData, func(s internet.StatDownload) {
		eventPage := EventPage{
			Images:       s.URL,
			IndexPage:    s.Index,
			StatError:    s.Err,
			StatusResume: s.StatusResume,
		}
		runtime.EventsEmit(*app.WailsContext, "dl_eventpage", eventPage)
	})

	return moreFailed, nil
}

type ParamJS struct {
	MangaID uint              `json:"manga_id"`
	Manga   string            `json:"manga"`
	Chapter types.ChapterList `json:"chapter" ts_type:"types.ChapterList"`
	Pages   []string          `json:"pages"`
}

type EventChap struct {
	ListPage  []string    `json:"list_page"`
	IndexChap int         `json:"index_chap"`
	TotalChap int         `json:"total_chap"`
	TotalPage int         `json:"total_page"`
	Chapter   json.Number `json:"chapter"`
}

type EventPage struct {
	Images       string      `json:"images"`
	IndexPage    int         `json:"index_page"`
	StatError    error       `json:"stat_error"`
	Chapter      json.Number `json:"chapter"`
	StatusResume bool        `json:"status_resume"`
}

// CheckTitleDB is id used for update status CheckChapterDB
type CheckTitleDB struct {
	Id    uint
	Title string
}

// CheckChapDB chapter  used for update status CheckChapterDB
type CheckChapDB struct {
	ID      uint    `json:"id"`
	Chapter float32 `json:"chapter"`
	Title   string  `json:"title"`
}

// CheckChapterDB changing status based on DB
func (f *Download) CheckChapterDB(c types.Chapter) types.Chapter {
	var checkTitleDB CheckTitleDB
	ilike := fmt.Sprintf("%%%s%%", c.Manga)
	app.DB.Table("mangas").Select("mangas.id", "mangas.title").
		Joins("left join manga_alternatives on manga_alternatives.manga_id = mangas.id").
		Where("mangas.title LIKE ?", ilike).
		Or("manga_alternatives.title LIKE ?", ilike).Scan(&checkTitleDB)

	c.MangaId = checkTitleDB.Id
	if checkTitleDB.Title != "" {
		c.Manga = checkTitleDB.Title
	}

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

type PageReport struct {
	Manga    string            `json:"manga"`
	StatusDL bool              `json:"status_dl"`
	Error    string            `json:"error"`
	FailChap []FiledChapReport `json:"fail_chap"`
}

type FiledChapReport struct {
	ChapID   string                  `json:"chap_id"`
	Chapter  json.Number             `json:"chapter"`
	StatusDB bool                    `json:"status_db"`
	Err      error                   `json:"error"`
	FailPage []internet.StatDownload `json:"fail_page"`
}

// Download Cover
func downloadCover(dst, url string) error {
	bytes, err := app.Client.Get(url).Bytes()
	if err != nil {
		return err
	}

	imgRef, err := vips.NewImageFromBuffer(bytes)
	if err != nil {
		return err
	}

	// autocrop to 300
	helper.AutoCrop(imgRef, 300)

	ep := vips.NewWebpExportParams()
	ep.StripMetadata = true
	ep.Quality = 60

	webpByte, _, err := imgRef.ExportWebp(ep)
	if err != nil {
		return err
	}
	// save exported byte to file
	err = os.WriteFile(dst, webpByte, 0644)
	if err != nil {
		return err
	}
	return nil
}
