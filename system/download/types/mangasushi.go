package types

type Mangasushi struct {
	Downloads
}

func (d Mangasushi) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Mangasushi"
	return &c, nil
}

func (d Mangasushi) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Mangasushi"
	return &c, nil
}
