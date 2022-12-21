package types

import (
	"encoding/json"
	"reflect"
)

// CLASS is collection of manga server struct
var CLASS = map[string]interface{}{
	"Mangadex":      new(Mangadex),
	"Lhtranslation": new(Lhtranslation),
	"Ksgroupscans":  new(Ksgroupscans),
	"Gdegenscans":   new(Gdegenscans),
	"Westmanga":     new(Westmanga),
	"Mangasushi":    new(Mangasushi),
	"Flamescans":    new(Flamescans),
	"Readmanganato": new(Readmanganato),
	"KomikCast":     new(KomikCast),
	"Kiryuu":        new(Kiryuu),
	"PlatinumScans": new(PlatinumScans),
	"Komikindo":     new(Komikindo),
	"Mangatale":     new(Mangatale),
	"Mangkomik":     new(Mangkomik),
	"Inazumanga":    new(Inazumanga),
	"Maid":          new(Maid),
	"Mangaindo":     new(Mangaindo),
	"Masterkomik":   new(Masterkomik),
}

// Downloads is interface for every manga server struct
type Downloads interface {
	GetChapter(o Option) (*Chapter, error)
	GetPage(o Option) (*Page, error)
}

// Option is struct for downloading chapter or pages
type Option struct {
	URL        string `json:"url"`
	ServerName string `json:"server_name"`
	Page       *int   `json:"page"`
	Limit      *int   `json:"limit"`
}

// NewDownload return Downloads interface from manga struct with nil as not found
func NewDownload(className string) *Downloads {
	var d Downloads
	for c := range CLASS {
		if className == c {
			ref := reflect.ValueOf(CLASS[c])
			d = ref.Interface().(Downloads)
		}
	}
	return &d
}

// Chapter is main data output for GetChapter methods
type Chapter struct {
	Manga      string        `json:"manga"`
	Cover      string        `json:"cover_url"`
	Mdex       string        `json:"mdex"`
	MangaId    *int          `json:"manga_id"`
	ServerName int           `json:"server_name"`
	Chapter    []ChapterList `json:"chapter"`
	Limit      int           `json:"limit"`
	Total      int           `json:"total"`
	MdexData   MdexManga     `json:"mdextest"`
}

// ChapterList is list of chapter for main data Chapter for GetChapter mehotds
type ChapterList struct {
	ID        string      `json:"id"`
	Chapter   json.Number `json:"chapter"`
	Title     string      `json:"title"`
	Volume    string      `json:"volume"`
	Timestamp int         `json:"timestamp"`
	Languange string      `json:"language"`
	GroupName string      `json:"group_name"`
	Status    bool        `json:"status"`
	Check     bool        `json:"check"`
}

// Page is main output for GetPage methods
type Page struct {
	IDChap     int       `json:"id_chap"`
	Title      string    `json:"title"`
	Pages      []string  `json:"pages"`
	PagesSaver *[]string `json:"pagesSaver"`
	Status     bool      `json:"status"`
}