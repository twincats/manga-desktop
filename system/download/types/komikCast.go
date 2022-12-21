package types

type KomikCast struct {
	Downloads
}

func (d KomikCast) GetChapter(o Option) (*Chapter, error) {
	var c Chapter
	c.Manga = "KomikCast"
	return &c, nil
}

func (d KomikCast) GetPage(o Option) (*Page, error) {
	var c Page
	c.Title = "KomikCast"
	return &c, nil
}
