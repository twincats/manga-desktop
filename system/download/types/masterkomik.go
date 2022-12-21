package types

type Masterkomik struct {
	Downloads
}

func (d Masterkomik) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Masterkomik"
	return &c, nil
}

func (d Masterkomik) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Masterkomik"
	return &c, nil
}
