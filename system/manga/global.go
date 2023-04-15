package manga

import (
	"os"
	"regexp"

	"github.com/climech/naturalsort"
)

var (
	MangaPath string = `D:\DATA\Manga`
)

// Pagination struct for pagination need set Total
type Pagination struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	To          int `json:"to"`
	Total       int `json:"total"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"perpage"`
	Offset      int `json:"-"`
}

// set value for pagination struct
func (p *Pagination) Paginate() {
	//set From
	p.From = p.PerPage*p.CurrentPage - (p.PerPage - 1)

	//set To
	p.To = p.CurrentPage * p.PerPage
	if p.To > p.Total {
		p.To = p.Total
	}

	//set LastPage
	p.LastPage = p.Total / p.PerPage
	if p.LastPage > 0 {
		p.LastPage += 1
	}

	//limit currentpage to available page
	if p.CurrentPage < 1 {
		p.CurrentPage = 1
	}

	if p.CurrentPage > p.LastPage {
		p.CurrentPage = p.LastPage
	}

	//setting Offset for DB Query
	p.Offset = (p.CurrentPage - 1) * p.PerPage

}

// GetFiles fetch all files within directory
func GetFiles(path string) ([]string, error) {
	var files []string
	list, err := os.ReadDir(path)
	if err != nil {
		return files, err
	}

	for _, file := range list {
		if !file.IsDir() {
			files = append(files, file.Name())
		}
	}

	// sorting files as natural
	naturalsort.Sort(files)

	return files, nil
}

// StringReplace Regex string replace
func StringReplace(regexString string, match string, replaceString string) string {
	reg := regexp.MustCompile(regexString)
	return reg.ReplaceAllString(match, replaceString)
}

func fixMangaTitle(title string) string {
	mTitle := StringReplace(`(?m)[^\w\-[\]()' ~.,!@&]|\.+$`, title, "")
	mTitle = StringReplace("(?m) $|^ ", mTitle, "")
	return mTitle
}
