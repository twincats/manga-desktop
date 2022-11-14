package manga

import (
	"encoding/json"
	"mangav4/system/app"

	"github.com/SamuelTissot/sqltime"
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
	ID               uint         `json:"id"`
	Chapter          json.Number  `json:"chapter"`
	Title            string       `json:"title"`
	Volume           string       `json:"volume"`
	Group            string       `json:"group"`
	TimestampRelease int          `json:"timestamp_release"`
	StatusRead       bool         `json:"status_read"`
	LanguageId       int          `json:"language_id"`
	Language         Language     `gorm:"foreignKey:LanguageId" json:"language"`
	CreatedAt        sqltime.Time `json:"created_at"`
	UpdatedAt        sqltime.Time `json:"updated_at"`
	MangaId          int          `json:"-"`
}
