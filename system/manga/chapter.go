package manga

import (
	"mangav4/system/app"
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
	MangaId          uint      `json:"-"`
}

func (f *Chapter) InsertChapter(c Chapter) (uint, error) {
	res := app.DB.FirstOrCreate(&c, c)
	if res.Error != nil {
		return 0, res.Error
	}
	return c.ID, nil
}
