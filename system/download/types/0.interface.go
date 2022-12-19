package types

import (
	"encoding/json"
	"reflect"
)

var CLASS = map[string]interface{}{
	"Mangadex":      new(Mangadex),
	"Lhtranslation": new(Lhtranslation),
}

type Downloads interface {
	GetChapter(o OptionDownload) (*Chapter, error)
	GetPage(o OptionDownload) (*Page, error)
}

type OptionDownload struct {
	Id string
}

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

type Chapter struct {
	Manga    string        `json:"manga"`
	Cover    string        `json:"cover_url"`
	Mdex     string        `json:"mdex"`
	MangaId  int           `json:"manga_id"`
	ServerId int           `json:"server_id"`
	Chapter  []ChapterList `json:"chapter"`
	Limit    int           `json:"limit"`
	Total    int           `json:"total"`
}

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

type Page struct {
	IDChap     int       `json:"id_chap"`
	Title      string    `json:"title"`
	Pages      []string  `json:"pages"`
	PagesSaver *[]string `json:"pagesSaver"`
	Status     bool      `json:"status"`
}
