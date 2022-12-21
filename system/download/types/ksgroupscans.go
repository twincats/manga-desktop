package types

type Ksgroupscans struct {
	Downloads
}

func (d Ksgroupscans) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "Ksgroupscans"
	return &c, nil
}

func (d Ksgroupscans) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "Ksgroupscans"
	return &c, nil
}
