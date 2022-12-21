package types

type Mangaindo struct {
	Downloads
}

func (d Mangaindo) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Mangaindo"
	return &c, nil
}

func (d Mangaindo) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Mangaindo"
	return &c, nil
}
