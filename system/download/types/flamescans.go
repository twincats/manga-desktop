package types

type Flamescans struct {
	Downloads
}

func (d Flamescans) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Flamescans"
	return &c, nil
}

func (d Flamescans) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Flamescans"
	return &c, nil
}
