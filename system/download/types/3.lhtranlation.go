package types

type Lhtranslation struct {
	Downloads
}

func (d Lhtranslation) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Lhtranslation"
	return &c, nil
}

func (d Lhtranslation) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Lhtranslation"
	return &c, nil
}
