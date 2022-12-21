package types

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

type Mangadex struct {
	Downloads
}

func (d Mangadex) GetChapter(o Option) (*Chapter, error) {
	var mdex = getMangadexIdFromUrl(o.URL)
	urlManga := fmt.Sprintf("https://api.mangadex.org/manga/%s?includes[]=cover_art", mdex)

	_, body, err := fasthttp.Get(nil, urlManga)
	if err != nil {
		return nil, err
	}

	var mdexmanga MdexManga
	err = json.Unmarshal(body, &mdexmanga)
	if err != nil {
		return nil, err
	}

	var c Chapter
	c.Manga = "Mangadex " + mdex
	c.MdexData = mdexmanga
	c.Mdex = mdex
	return &c, nil
}

func (d Mangadex) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Mangadex"
	return &c, nil
}

func getMangadexIdFromUrl(mdex string) string {
	if _, err := uuid.Parse(mdex); err == nil {
		return mdex
	}
	mdexL := strings.Split(mdex, "/")
	for _, s := range mdexL {
		if len(s) == 36 {
			return s
		}
	}
	return ""
}

type MdexManga struct {
	Data struct {
		Id         string
		Attributes struct {
			Title       MdexLang
			AltTitles   []MdexLang
			Description MdexLang
		}
		Relationships []Relationship
	}
}

type MdexLang struct {
	En   string
	Ja   string
	JaRo string `json:"ja-ro"`
}

type Relationship struct {
	Id         string
	Type       string
	Attributes struct {
		Name     string
		FileName string
	}
}
