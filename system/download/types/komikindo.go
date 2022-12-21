package types

type Komikindo struct {
	Downloads
}

func (d Komikindo) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Komikindo"
	return &c, nil
}

func (d Komikindo) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Komikindo"
	return &c, nil
}
