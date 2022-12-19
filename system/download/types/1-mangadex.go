package types

type Mangadex struct {
	Downloads
}

func (d Mangadex) GetChapter(o OptionDownload) (*Chapter, error) {
	var c Chapter
	c.Manga = "Mangadex"
	return &c, nil
}

func (d Mangadex) GetPage(o OptionDownload) (*Page, error) {
	var c Page
	c.Title = "Mangadex"
	return &c, nil
}
