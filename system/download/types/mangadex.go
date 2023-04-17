package types

import (
	"encoding/json"
	"fmt"
	"mangav4/system/app"
	"mangav4/system/app/internet"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Mangadex struct {
	Downloads
}

func (d Mangadex) GetChapter(o Option) (*Chapter, error) {
	var mdex = getMangadexIdFromUrl(o.URL)
	limit, offset := getDefaultLimitOffset(o)

	urlManga := fmt.Sprintf("https://api.mangadex.org/manga/%s?includes[]=cover_art", mdex)
	urlChapter := fmt.Sprintf("https://api.mangadex.org/manga/%s/feed?translatedLanguage[]=en&translatedLanguage[]=id&includes[]=scanlation_group&limit=%d&offset=%d&order[chapter]=desc", mdex, limit, offset)

	var UrlList = []string{urlManga, urlChapter}
	result := ParallelFetch(UrlList)

	var mdexmanga MdexManga
	var mdexChapter MdexChapter

	for _, fetchResult := range result {
		if fetchResult.Err != nil {
			return nil, fetchResult.Err
		}

		switch fetchResult.ID {
		case 0:
			err := internet.BodyParser(fetchResult.Body, &mdexmanga)
			if err != nil {
				return nil, err
			}
		case 1:
			err := internet.BodyParser(fetchResult.Body, &mdexChapter)
			if err != nil {
				return nil, err
			}
		}
	}

	var c Chapter
	c.Mdex = mdex
	c.ServerName = o.ServerName
	c.Limit = mdexChapter.Limit
	c.Total = mdexChapter.Total
	setMangadexTitle(mdexmanga, &c)
	setMangadexCover(mdexmanga, &c)
	c.Chapter = setMangadexChapter(mdexChapter)

	return &c, nil
}

func (d Mangadex) GetPage(o Option) (*Page, error) {
	mdex_page_url := fmt.Sprintf("https://api.mangadex.org/at-home/server/%s", o.URL)
	pageByte, err := app.Client.Get(mdex_page_url).Bytes()
	if err != nil {
		return nil, err
	}

	var mdex_page MdexPages
	err = internet.BodyParser(pageByte, &mdex_page)
	if err != nil {
		return nil, err
	}

	var c Page
	c.Title = "Mangadex"
	var pages []string
	if *o.DataSaver {
		for _, v := range mdex_page.Chapter.DataSaver {
			url := fmt.Sprintf("%s/data-saver/%s/%s", mdex_page.BaseUrl, mdex_page.Chapter.Hash, v)
			pages = append(pages, url)
		}
	} else {
		for _, v := range mdex_page.Chapter.Data {
			url := fmt.Sprintf("%s/data/%s/%s", mdex_page.BaseUrl, mdex_page.Chapter.Hash, v)
			pages = append(pages, url)
		}
	}
	c.Pages = pages

	return &c, nil
}

func (f *Mangadex) GetChapterMdexPagination(url string, limit, offset int) ([]ChapterList, error) {
	mdex := getMangadexIdFromUrl(url)
	urlChapter := fmt.Sprintf("https://api.mangadex.org/manga/%s/feed?translatedLanguage[]=en&translatedLanguage[]=id&includes[]=scanlation_group&limit=%d&offset=%d&order[chapter]=desc", mdex, limit, offset)
	bodyByte, err := app.Client.Get(urlChapter).Bytes()
	if err != nil {
		return nil, err
	}

	var mdexChap MdexChapter
	err = internet.BodyParser(bodyByte, &mdexChap)
	if err != nil {
		return nil, err
	}

	return setMangadexChapter(mdexChap), nil

}

// getMangadexIdFromUrl is get mdex UUID from url / string
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

// getDefaultLimitOffset return (limit,offset)
func getDefaultLimitOffset(o Option) (int, int) {
	var limit = 30
	var offset = 0

	if o.Limit != nil {
		limit = *o.Limit
	}

	if o.Offset != nil {
		offset = *o.Offset
	}
	return limit, offset
}

// setMangadexTitle is fill up mangadex Title
func setMangadexTitle(m MdexManga, c *Chapter) {
	if m.Data.Attributes.Title.En != "" {
		c.Manga = m.Data.Attributes.Title.En
	} else if m.Data.Attributes.Title.JaRo != "" {
		c.Manga = m.Data.Attributes.Title.JaRo
	} else if m.Data.Attributes.Title.Ja != "" {
		c.Manga = m.Data.Attributes.Title.Ja
	} else {
		c.Manga = ""
	}
}

// setMangadexCover is fill up mangadex cover
func setMangadexCover(m MdexManga, c *Chapter) {
	for _, v := range m.Data.Relationships {
		if v.Type == "cover_art" {
			c.Cover = fmt.Sprintf("https://uploads.mangadex.org/covers/%s/%s.256.jpg", m.Data.Id, v.Attributes.FileName)
		}
	}
}

// setMangadexChapter is fill up mangadex chapter
func setMangadexChapter(m MdexChapter) []ChapterList {
	var chapterList []ChapterList
	for _, v := range m.Data {
		chapterList = append(chapterList, ChapterList{
			ID:        v.Id,
			Chapter:   v.Attributes.Chapter,
			Volume:    v.Attributes.Volume,
			Title:     v.Attributes.Title,
			Timestamp: int(v.Attributes.PublishAt.Unix()),
			Languange: getLang(v.Attributes.TranslatedLanguage),
			GroupName: getGroupName(v.Relationships),
		})
	}
	return chapterList
}

// getLang  2 digit string full langauage string (English | Indonesia)
func getLang(s string) string {
	if s == "en" {
		return "English"
	} else {
		return "Indonesia"
	}
}

// getGroupName turn relationshio group as string
func getGroupName(R []MdexRelationship) string {
	var groups []string
	for _, r := range R {
		if r.Type == "scanlation_group" {
			groups = append(groups, r.Attributes.Name)
		}
	}
	if len(groups) > 0 {
		return joinString(groups)
	} else {
		return "No Groups"
	}
}

// JoinString join array string with &
func joinString(ss []string) string {
	var v string
	if len(ss) > 1 {
		for i, s := range ss {
			if i+1 < len(ss) {
				v += s + " & "
			} else {
				v += s
			}
		}
	} else if len(ss) == 1 {
		v = ss[0]
	} else {
		v = ""
	}
	return v
}

type MdexManga struct {
	Data struct {
		Id         string
		Attributes struct {
			Title       MdexLang
			AltTitles   []MdexLang
			Description MdexLang
		}
		Relationships []MdexRelationship
	}
}

type MdexLang struct {
	En   string
	Ja   string
	JaRo string `json:"ja-ro"`
}

type MdexRelationship struct {
	Id         string
	Type       string
	Attributes struct {
		Name     string
		FileName string
	}
}

type MdexChapter struct {
	Data []struct {
		Id         string
		Attributes struct {
			Chapter            json.Number
			Volume             string
			Title              string
			TranslatedLanguage string
			PublishAt          time.Time
		}
		Relationships []MdexRelationship
	}
	Limit  int
	Offset int
	Total  int
}

type MdexPages struct {
	BaseUrl string
	Chapter struct {
		Hash      string
		Data      []string
		DataSaver []string
	}
}
