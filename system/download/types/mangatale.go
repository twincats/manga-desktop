package types

type Mangatale struct {
	Downloads
}

func (d Mangatale) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Mangatale"
	return &c, nil
}

func (d Mangatale) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Mangatale"
	return &c, nil
}
