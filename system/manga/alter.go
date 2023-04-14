package manga

import "mangav4/system/app"

type Alter struct {
	ID      uint   `json:"id"`
	Title   string `json:"title"`
	MangaId int    `json:"manga_id"`
}

// TableName overrides the table name used by User to `profiles`
func (Alter) TableName() string {
	return "manga_alternatives"
}

// NewAlter create new instance of Alter
func NewAlter() *Alter {
	return &Alter{}
}

func (f *Alter) GetAlters() ([]Alter, error) {
	var alter []Alter
	res := app.DB.Find(&alter)
	if res.Error != nil {
		return alter, res.Error
	}
	return alter, nil
}

func (f *Alter) SetAlter(a Alter) error {
	res := app.DB.Save(&a)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
