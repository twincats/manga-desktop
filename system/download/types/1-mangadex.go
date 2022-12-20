package types

type Mangadex struct {
	Downloads
}

func (d Mangadex) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Mangadex"
	return &c, nil
}

func (d Mangadex) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Mangadex"
	return &c, nil
}
