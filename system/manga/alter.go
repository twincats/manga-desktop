package manga

type Alter struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	MangaId int    `json:"manga_id"`
}

// TableName overrides the table name used by User to `profiles`
func (Alter) TableName() string {
	return "manga_alternatives"
}
