package types

type Kiryuu struct {
	Downloads
}

func (d Kiryuu) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Kiryuu"
	return &c, nil
}

func (d Kiryuu) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Kiryuu"
	return &c, nil
}
