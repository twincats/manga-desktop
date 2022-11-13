package manga

import (
	"encoding/json"
	"path/filepath"

	"mangav4/system"

	"github.com/SamuelTissot/sqltime"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gorm.io/gorm"
)

// NewManga create new instance of Manga
func NewManga(db *gorm.DB) *Manga {
	DB = db
	return &Manga{}
}

// GetManga get single manga by ID
func (m *Manga) GetManga(id int) Manga {
	var manga Manga
	DB.Order("title").First(&manga, id)
	return manga
}

// GetMangas get list of manga []manga
func (m *Manga) GetMangas() []Manga {
	var manga []Manga
	DB.Order("title").Find(&manga)
	return manga
}

// GetMangaHome get list of mangaHome with paginate []manga
func (m *Manga) GetMangaHome(page int, limit int) MangaHome {
	mangaHomeApi := MangaHomeApi{}

	return mangaHomeApi.Paginate(page, limit)
}

/* GetPage for Fetching list of images chapter */
func (m *Manga) GetPage(id int) Page {
	var p PageApi
	DB.Table("chapters").
		Select("chapters.id", "chapter", "mangas.title").
		Joins("inner join mangas on chapters.manga_id = mangas.id").
		Where("chapters.id = ?", id).
		Take(&p)

	urlPath := filepath.Join(fixMangaTitle(p.Title), p.Chapter)
	path := filepath.Join(MangaPath, urlPath)
	files, err := GetFiles(path)
	if err != nil {
		runtime.LogError(*system.WailsContext, err.Error())
	}

	var page Page
	page.Pages = files
	page.Path = urlPath

	return page
}

// PageApi for Fetching chapter and Manga title
type PageApi struct {
	ID      int
	Chapter string
	Title   string
}

// Manga Struct for model manga
type Manga struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	StatusEnding bool         `json:"status_end"`
	Mdex         *int         `orm:"null" json:"mdex"`
	CreatedAt    sqltime.Time `orm:"auto_now_add;type(datetime)" json:"-"`
	UpdatedAt    sqltime.Time `orm:"auto_now;type(datetime)" json:"-"`
}

type MangaHomeApi struct {
	ID           int          `json:"id"`
	Title        string       `json:"title"`
	StatusEnding bool         `json:"status_end"`
	Mdex         *int         `orm:"null" json:"mdex"`
	ChapterID    int          `json:"chapter_id"`
	Chapter      json.Number  `json:"chapter"`
	DownloadTime sqltime.Time `json:"download_time"`
}

func (m *MangaHomeApi) Paginate(page int, limit int) MangaHome {
	// getQuery
	query := m.FetchQuery()

	// get totalManga
	var totalManga int64
	query.Count(&totalManga)

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
	latestSubQuery := DB.Table("chapters").Select("manga_id", "MAX(CAST(chapter AS decimal)) AS chapter").Group("manga_id").Order("manga_id")
	homeApiQuery := DB.Table("mangas").
		Select("mangas.id", "mangas.title", "mdex", "status_ending", "chapters.id as chapter_id", "latest.chapter", "TO_DATE(cast(chapters.created_at as TEXT) ,'YYYY-MM-DD') as download_time").
		Joins("inner join chapters on mangas.id = chapters.manga_id").
		Joins("inner join (?) latest on mangas.id = latest.manga_id", latestSubQuery).
		Where("latest.chapter = CAST(chapters.chapter as decimal)").
		Order("download_time desc, mangas.title")
	return homeApiQuery
}

type MangaHome struct {
	Manga      []MangaHomeApi `json:"manga"`
	Pagination *Pagination    `json:"pagination"`
}

type Page struct {
	Pages []string `json:"pages"`
	Path  string   `json:"path"`
	Nav   Nav      `json:"nav"`
}

type Nav struct {
	Next string `json:"next"`
	Prev string `json:"prev"`
}
