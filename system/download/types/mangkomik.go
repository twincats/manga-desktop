package types

type Mangkomik struct {
	Downloads
}

func (d Mangkomik) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Mangkomik"
	return &c, nil
}

func (d Mangkomik) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Mangkomik"
	return &c, nil
}
