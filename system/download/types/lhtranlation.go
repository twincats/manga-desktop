package types

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Lhtranslation struct {
	Downloads
}

func (d Lhtranslation) GetChapter(o Option) (*Chapter, error) {
	var mdex = getMangadexIdFromUrl(o.URL)
	urlManga := fmt.Sprintf("https://api.mangadex.org/manga/%s?includes[]=cover_art", mdex)
	resp, err := http.Get(urlManga)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var mdexmanga MdexManga
	err = json.Unmarshal(body, &mdexmanga)
	if err != nil {
		return nil, err
	}

	var c Chapter
	c.Manga = "Lhtranslation " + mdex
	c.MdexData = mdexmanga
	c.Mdex = mdex
	return &c, nil
}

func (d Lhtranslation) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Lhtranslation"
	return &c, nil
}
