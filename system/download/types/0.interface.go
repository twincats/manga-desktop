package types

import (
	"encoding/json"
	"mangav4/system/app"
	"reflect"
)

// CLASS is collection of manga server struct
var CLASS = map[string]interface{}{
	"Mangadex":      new(Mangadex),
	"Lhtranslation": new(Lhtranslation),
	"Ksgroupscans":  new(Ksgroupscans),
	"Gdegenscans":   new(Gdegenscans),
	"Westmanga":     new(Westmanga),
	"Mangasushi":    new(Mangasushi),
	"Flamescans":    new(Flamescans),
	"Readmanganato": new(Readmanganato),
	"KomikCast":     new(KomikCast),
	"Kiryuu":        new(Kiryuu),
	"PlatinumScans": new(PlatinumScans),
	"Komikindo":     new(Komikindo),
	"Mangatale":     new(Mangatale),
	"Mangkomik":     new(Mangkomik),
	"Inazumanga":    new(Inazumanga),
	"Maid":          new(Maid),
	"Mangaindo":     new(Mangaindo),
	"Masterkomik":   new(Masterkomik),
}

// Downloads is interface for every manga server struct
type Downloads interface {
	GetChapter(o Option) (*Chapter, error)
	GetPage(o Option) (*Page, error)
}

// Option is struct for downloading chapter or pages
type Option struct {
	URL        string `json:"url"`
	ServerName string `json:"server_name"`
	Offset     *int   `json:"offset"`
	Limit      *int   `json:"limit"`
	DataSaver  *bool  `json:"datasaver"`
}

type OptionPage struct {
	ServerName string        `json:"server_name"`
	MangaTitle string        `json:"manga_title"`
	DataSaver  bool          `json:"datasaver"`
	Chapters   []ChapterList `json:"chapters"`
}

// NewDownload return Downloads interface from manga struct with nil as not found
func NewDownload(className string) *Downloads {
	var d Downloads
	for c := range CLASS {
		if className == c {
			ref := reflect.ValueOf(CLASS[c])
			d = ref.Interface().(Downloads)
		}
	}
	return &d
}

// Chapter is main data output for GetChapter methods
type Chapter struct {
	Manga      string        `json:"manga"`
	Cover      string        `json:"cover_url"`
	Mdex       string        `json:"mdex"`
	MangaId    uint          `json:"manga_id"`
	ServerName string        `json:"server_name"`
	Chapter    []ChapterList `json:"chapter"`
	Limit      int           `json:"limit"`
	Total      int           `json:"total"`
	MdexData   interface{}   `json:"mdextest"`
}

// ChapterList is list of chapter for main data Chapter for GetChapter mehotds
type ChapterList struct {
	ID        string      `json:"id"`
	Chapter   json.Number `json:"chapter"`
	Title     string      `json:"title"`
	Volume    string      `json:"volume"`
	Timestamp int         `json:"timestamp"`
	Languange string      `json:"language"`
	GroupName string      `json:"group_name"`
	Status    bool        `json:"status"`
	Check     bool        `json:"check"`
}

// Page is main output for GetPage methods
type Page struct {
	IDChap int      `json:"id_chap"`
	Title  string   `json:"title"`
	Pages  []string `json:"pages"`
	Status bool     `json:"status"`
}

// ParallelFetchResult is output for ParallelFetch
type ParallelFetchResult struct {
	ID   int
	Url  string
	Body []byte
	Err  error
}

func ParallelFetch(urlList []string) []ParallelFetchResult {
	var result []ParallelFetchResult

	var client = app.Client
	var channelResult = make(chan ParallelFetchResult)

	for i, s := range urlList {
		go func(url string, i int) {
			//time.Sleep(time.Second)
			bodyByte, err := client.Get(url).Bytes()
			channelResult <- ParallelFetchResult{
				ID:   i,
				Url:  url,
				Body: bodyByte,
				Err:  err,
			}
		}(s, i)
	}

	for i := 0; i < len(urlList); i++ {
		result = append(result, <-channelResult)
	}

	return result
}
