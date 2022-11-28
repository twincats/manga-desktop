package system

import (
	"mangav4/system/manga"
	"mangav4/system/tool"
)

// initialize bind instance
func InitializeBinding() []interface{} {

	binding := []interface{}{
		manga.NewManga(),
		manga.NewChapter(),
		tool.NewDialog(),
	}

	return binding
}
