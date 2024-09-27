package manga

import (
	"fmt"
	"mangav4/system/app"
)

type Language struct {
	ID       int    `json:"id"`
	Lang     string `json:"lang"`
	LangCode string `json:"lang_code"`
}

func LangByLang(s string) Language {
	var l Language
	app.DB.Where("lang = ?", s).First(&l)
	return l
}

func SeedingLanguages() {
	var languagesSeed = []Language{
		{ID: 1, Lang: "English", LangCode: "en"},
		{ID: 2, Lang: "Indonesia", LangCode: "id"},
	}

	result := app.DB.Create(&languagesSeed)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
