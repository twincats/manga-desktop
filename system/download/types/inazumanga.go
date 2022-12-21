package types

type Inazumanga struct {
	Downloads
}

func (d Inazumanga) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Inazumanga"
	return &c, nil
}

func (d Inazumanga) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Inazumanga"
	return &c, nil
}
