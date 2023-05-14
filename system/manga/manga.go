package manga

import (
	"fmt"
	"path/filepath"
	"time"

	"mangav4/system/app"
	"mangav4/system/file"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

// NewManga create new instance of Manga
func NewManga() *Manga {
	return &Manga{}
}

// GetManga get single manga by ID
func (m *Manga) GetManga(id uint) Manga {
	var manga Manga
	app.DB.Order("title").First(&manga, id)
	return manga
}

// GetMangaWithChapter get single manga with all chapter
func (m *Manga) GetMangaWithChapter(id uint) Manga {
	var manga Manga
	app.DB.Preload("Chapter.Language").Preload("Alter").Order("title").First(&manga, id)
	return manga
}

// GetMangas get list of manga []manga
func (m *Manga) GetMangas() []Manga {
	var manga []Manga
	app.DB.Order("title").Find(&manga)
	return manga
}

// GetMangaHome get list of mangaHome with paginate []manga
func (m *Manga) GetMangaHome(title *string, page int, limit int) MangaHome {
	mangaHomeApi := MangaHomeApi{}

	return mangaHomeApi.Paginate(title, page, limit)
}

/* GetPage for Fetching list of images chapter */
func (m *Manga) GetPage(id int) Page {
	var p PageApi
	app.DB.Table("chapters").
		Select("chapters.id", "chapter", "mangas.title", "manga_id").
		Joins("inner join mangas on chapters.manga_id = mangas.id").
		Where("chapters.id = ?", id).
		Take(&p)

	urlPath := filepath.Join(fixMangaTitle(p.Title), fmt.Sprintf("%g", p.Chapter))
	path := filepath.Join(file.MANGA_PATH, urlPath)
	files, err := GetFiles(path)
	if err != nil {
		runtime.LogError(*app.WailsContext, err.Error())
	}

	var page Page
	page.Pages = files
	page.Path = urlPath
	page.MangaId = p.MangaId

	// fetch all chapter
	var pnav []PageApiNav
	app.DB.Table("chapters").Select("id").
		Where("manga_id = ?", p.MangaId).
		Order("CAST(chapter as DECIMAL)").Find(&pnav)

	for _, n := range pnav {
		page.Chaps = append(page.Chaps, n.ID)
	}

	return page
}

// GetServer fetch all server data
func (f *Manga) GetServer() []Server {
	var servers []Server
	app.DB.Find(&servers)
	return servers
}

func (f *Manga) InsertManga(m Manga) (uint, error) {
	res := app.DB.FirstOrCreate(&m, m)
	if res.Error != nil {
		return 0, res.Error
	}
	return m.ID, nil
}

func (f *Manga) TestHomeApiQuery() interface{} {
	var m MangaHomeApi
	q := m.FetchQuery()
	v := []MangaHomeApi{}
	q.Find(&v)
	return v
}

func (f *Manga) UpdateChapJS(serverID uint, chapJS string) (bool, error) {
	var server Server
	server.ID = serverID
	res := app.DB.Model(&server).Update("chap_js_code", chapJS)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

func (f *Manga) UpdatePagesJS(serverID uint, pageJS string) (bool, error) {
	var server Server
	server.ID = serverID
	res := app.DB.Model(&server).Update("page_js_code", pageJS)
	if res.Error != nil {
		return false, res.Error
	}
	return true, nil
}

// PageApi for Fetching chapter and Manga title
type PageApi struct {
	ID      uint
	Chapter float32
	Title   string
	MangaId uint
}

type PageApiNav struct {
	ID uint
}

// Manga Struct for model manga
type Manga struct {
	ID           uint      `json:"id"`
	Title        string    `gorm:"not null" json:"title"`
	StatusEnding bool      `json:"status_end"`
	Mdex         string    `gorm:"size:36" orm:"null" json:"mdex"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	Chapter      []Chapter `json:"chapter"`
	Alter        []Alter   `json:"alter"`
}

type MangaHomeApi struct {
	ID           uint      `json:"id"`
	Title        string    `json:"title"`
	StatusEnding bool      `json:"status_end"`
	Mdex         string    `orm:"null" json:"mdex"`
	ChapterID    uint      `json:"chapter_id"`
	Chapter      float32   `json:"chapter"`
	DownloadTime time.Time `json:"download_time" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
}

func (m *MangaHomeApi) Paginate(title *string, page int, limit int) MangaHome {
	// declaration for total data is fetched
	var totalManga int64

	// getQuery
	query := m.FetchQuery()

	// if only
	if title != nil && *title != "" {
		ilike := fmt.Sprintf("%%%s%%", *title)
		query = app.DB.Table("(?) as mh", query).
			Select("mh.id", "mh.title", "mdex", "status_ending", "chapter_id", "chapter", "download_time").
			Joins("left join manga_alternatives on manga_alternatives.manga_id = mh.id").
			Where("mh.title LIKE ?", ilike).
			Or("manga_alternatives.title LIKE ?", ilike).
			Group("mh.id, mh.title,mdex,status_ending,chapter_id,chapter,download_time").
			Order("mh.title")

		// get totalManga with searhing manga title
		app.DB.Table("(?) as mhome", query).Count(&totalManga)

	} else {
		// get totalManga without searhing manga title
		query.Count(&totalManga)
	}

	// set pagination
	pagination := &Pagination{CurrentPage: page, Total: int(totalManga), PerPage: limit}
	pagination.Paginate()

	mhApi := []MangaHomeApi{}
	query.Offset(pagination.Offset).Limit(pagination.PerPage).Find(&mhApi)

	return MangaHome{
		Manga:      mhApi,
		Pagination: pagination,
	}

}

func (f MangaHomeApi) FetchQuery() *gorm.DB {
	latestSubQuery := app.DB.Table("chapters").Select("manga_id", "MAX(CAST(chapter AS decimal)) AS chapter").Group("manga_id").Order("manga_id")
	homeApiQuery := app.DB.Table("mangas").
		Select("mangas.id", "mangas.title", "mdex", "status_ending", "chapters.id as chapter_id", "latest.chapter", "chapters.created_at as download_time").
		Joins("inner join chapters on mangas.id = chapters.manga_id").
		Joins("inner join (?) latest on mangas.id = latest.manga_id", latestSubQuery).
		Where("latest.chapter = CAST(chapters.chapter as decimal)").
		Order("download_time desc, mangas.title")
	return homeApiQuery
}

type MangaHome struct {
	Manga      []MangaHomeApi `json:"manga"`
	Pagination *Pagination    `json:"pagination" ts_type:"Pagination"`
}

type Page struct {
	Pages   []string `json:"pages"`
	Path    string   `json:"path"`
	Chaps   []uint   `json:"chaps"`
	MangaId uint     `json:"manga_id"`
}

// type Nav struct {
// 	Next string `json:"next"`
// 	Prev string `json:"prev"`
// }
