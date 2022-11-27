package system

import (
	"mangav4/system/manga"
)

// initialize bind instance
func InitializeBinding() []interface{} {

	binding := []interface{}{
		manga.NewManga(),
		manga.NewChapter(),
	}

	return binding
}
