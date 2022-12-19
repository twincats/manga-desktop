package types

type Lhtranslation struct {
	Downloads
}

func (d Lhtranslation) GetChapter(o OptionDownload) (*Chapter, error) {
	var c Chapter
	c.Manga = "Lhtranslation"
	return &c, nil
}

func (d Lhtranslation) GetPage(o OptionDownload) (*Page, error) {
	var c Page
	c.Title = "Lhtranslation"
	return &c, nil
}
