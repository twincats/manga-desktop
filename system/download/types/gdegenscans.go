package types

type Gdegenscans struct {
	Downloads
}

func (d Gdegenscans) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Gdegenscans"
	return &c, nil
}

func (d Gdegenscans) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Gdegenscans"
	return &c, nil
}
