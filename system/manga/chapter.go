package manga

import (
	"fmt"
	"mangav4/system/app"
	"mangav4/system/file"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// NewChapter create new instance of NewChapter
func NewChapter() *Chapter {
	return &Chapter{}
}

func (f *Chapter) GetChapter(id uint) Chapter {
	app.DB.Preload("Language").Find(f)
	return *f
}

type Chapter struct {
	ID               uint      `json:"id"`
	Chapter          float32   `gorm:"not null" json:"chapter"`
	Title            string    `json:"title"`
	Volume           string    `json:"volume"`
	Group            string    `json:"group"`
	TimestampRelease int64     `json:"timestamp_release"`
	StatusRead       bool      `json:"status_read"`
	LanguageId       uint      `json:"language_id"`
	Language         Language  `gorm:"foreignKey:LanguageId" json:"language" ts_type:"Language"`
	CreatedAt        time.Time `json:"created_at" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	UpdatedAt        time.Time `json:"updated_at" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	MangaId          uint      `json:"manga_id"`
}

func (f *Chapter) InsertChapter(c Chapter) (uint, error) {
	res := app.DB.FirstOrCreate(&c, c)
	if res.Error != nil {
		return 0, res.Error
	}
	return c.ID, nil
}

func (f *Chapter) InsertChapterBatch(c []Chapter) (*[]Chapter, error) {
	result := app.DB.Create(&c)
	if result.Error != nil {
		return nil, result.Error
	}
	return &c, nil
}

func (f *Chapter) DeleteChapterBatch(c []Chapter) (int64, error) {
	result := app.DB.Delete(&c)
	if result.Error != nil {
		return 0, result.Error
	}
	if len(c) == int(result.RowsAffected) {
		fmt.Println("BASE MANGA PATH" + file.MANGA_PATH)
		m := NewManga()
		title := m.GetManga(c[0].MangaId).Title
		for _, v := range c {
			err := os.RemoveAll(filepath.Join(file.MANGA_PATH, title, strconv.FormatFloat(float64(v.Chapter), 'g', -1, 32)))
			if err != nil {
				return 0, err
			}
		}
	}
	return result.RowsAffected, nil
}

func (f *Chapter) UpdateRead(ID uint, status_read bool) error {
	result := app.DB.Model(&Chapter{ID: ID}).Update("status_read", status_read)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
