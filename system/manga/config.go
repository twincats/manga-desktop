package manga

import (
	"fmt"
	"mangav4/system/app"
	"mangav4/system/file"
	"path/filepath"
	"strconv"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Config struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	MangaFolder string `json:"manga_folder"`
}

// NewConfig create new instance of Config
func NewConfig() *Config {
	return &Config{}
}

func (f *Config) GetConfig() (Config, error) {
	res := app.DB.First(f, 1)
	if res.Error != nil {
		return *f, res.Error
	}
	return *f, nil
}

func (f *Config) SetConfig(c Config) error {
	res := app.DB.Save(&c)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (f *Config) ClearAllConfig() error {
	res := app.DB.Exec("DELETE FROM configs")
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (f *Config) AutoScanDirs() ([]string, error) {
	cf, err := f.GetConfig()
	if err != nil {
		return nil, err
	}
	mangaList := file.ReadDir(cf.MangaFolder)
	for _, m := range mangaList {
		var manga Manga
		manga.Title = m
		mid, _ := manga.InsertManga(manga)
		mangaTitlePath := filepath.Join(cf.MangaFolder, m)

		runtime.EventsEmit(*app.WailsContext, "status_scans", m)

		chaps := file.ReadDir(mangaTitlePath)
		var chapters []Chapter
		for _, cp := range chaps {
			cfloat, _ := strconv.ParseFloat(cp, 32)
			chapters = append(chapters, Chapter{
				Chapter:    float32(cfloat),
				LanguageId: 1,
				Group:      "Unknown",
				MangaId:    mid,
			})
		}
		//save chapter
		res := app.DB.Create(&chapters)
		if res.Error != nil {
			fmt.Println(res.Error)
		}
		chapters = nil
	}
	file.MANGA_PATH = cf.MangaFolder
	return mangaList, nil
}
