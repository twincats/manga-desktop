package types

type Westmanga struct {
	Downloads
}

func (d Westmanga) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Westmanga"
	return &c, nil
}

func (d Westmanga) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Westmanga"
	return &c, nil
}
