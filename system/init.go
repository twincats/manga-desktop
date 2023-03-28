package system

import (
	"fmt"
	"mangav4/system/app"
	"mangav4/system/download"
	"mangav4/system/file"
	"mangav4/system/manga"
	"mangav4/system/tool"
)

// initialize bind instance
func InitializeBinding() []interface{} {

	binding := []interface{}{
		manga.NewManga(),
		manga.NewChapter(),
		manga.NewConfig(),
		tool.NewDialog(),
		tool.NewWeb(),
		download.NewDownload(),
	}

	return binding
}

func DatabaseStartUp() {
	// connecting databse sqlite
	app.ConnectDatabaseSqlite("./mangav4-desktop.db3")

	// checking databse manga not empty / new created db
	var mangaEmpty int
	app.DB.Raw("SELECT EXISTS (SELECT 1 FROM mangas)").Scan(&mangaEmpty)
	if mangaEmpty == 0 {
		fmt.Println("empty data")
		app.DB.AutoMigrate(
			&manga.Config{},
			&manga.Manga{},
			&manga.Chapter{},
			&manga.Alter{},
			&manga.Server{},
			&manga.Language{},
		)

		var serverEmpty int
		app.DB.Raw("SELECT EXISTS (SELECT 1 FROM servers)").Scan(&serverEmpty)
		if serverEmpty == 0 {
			manga.SeedingServer()

			var languageEmpty int
			app.DB.Raw("SELECT EXISTS (SELECT 1 FROM languages)").Scan(&languageEmpty)
			if languageEmpty == 0 {
				manga.SeedingLanguages()
			}
		}

	} else {
		var configs manga.Config
		config, err := configs.GetConfig()
		if err != nil {
			fmt.Println(err)
		}

		file.MANGA_PATH = config.MangaFolder
	}

}
