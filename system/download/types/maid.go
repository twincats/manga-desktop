package types

type Maid struct {
	Downloads
}

func (d Maid) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Maid"
	return &c, nil
}

func (d Maid) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Maid"
	return &c, nil
}
