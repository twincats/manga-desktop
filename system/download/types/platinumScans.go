package types

type PlatinumScans struct {
	Downloads
}

func (d PlatinumScans) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "PlatinumScans"
	return &c, nil
}

func (d PlatinumScans) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "PlatinumScans"
	return &c, nil
}
