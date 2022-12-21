package types

type Readmanganato struct {
	Downloads
}

func (d Readmanganato) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Readmanganato"
	return &c, nil
}

func (d Readmanganato) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Readmanganato"
	return &c, nil
}
