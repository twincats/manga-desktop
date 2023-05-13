package manga

import (
	"fmt"
	"mangav4/system/app"
)

type Server struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Url          string `json:"url"`
	Search       bool   `json:"search"`
	StatusActive bool   `json:"status_active"`
	ChapJSCode   string `json:"chap_jscode"`
	PageJSCode   string `json:"page_jscode"`
}

func SeedingServer() {
	var serversSeed = []Server{
		{ID: 1, Name: "Mangadex", Url: "mangadex.org", Search: true, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 2, Name: "Mangakakalot", Url: "mangakakalot.com", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 3, Name: "Lhtranslation", Url: "lhtranslation.net", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 4, Name: "Ksgroupscans", Url: "ksgroupscans.com", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 5, Name: "Gdegenscans", Url: "gdscans.com", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 6, Name: "Westmanga", Url: "westmanga.info", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 7, Name: "Mangasushi", Url: "mangasushi.org", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 8, Name: "Flamescans", Url: "flamescans.org", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 9, Name: "Readmanganato", Url: "readmanganato.com", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 10, Name: "Readkomik", Url: "readkomik.com", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 11, Name: "KomikCast", Url: "komikcast.site", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 12, Name: "Kiryuu", Url: "kiryuu.id", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 13, Name: "PlatinumScans", Url: "platinumscans.com", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 14, Name: "Komikindo", Url: "komikindo.co", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 15, Name: "Mangatale", Url: "mangatale.co", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 16, Name: "Mangkomik", Url: "mangkomik.net", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 17, Name: "Tritinia", Url: "tritinia.org", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 18, Name: "Maid", Url: "www.maid.my.id", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 19, Name: "Mangaindo", Url: "mangaindo.me", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
		{ID: 20, Name: "Masterkomik", Url: " masterkomik.com", Search: false, StatusActive: true, ChapJSCode: "", PageJSCode: ""},
	}

	result := app.DB.Create(&serversSeed)
	if result.Error != nil {
		fmt.Println(result.Error)
	}
}
