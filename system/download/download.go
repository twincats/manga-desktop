package download

import (
	"errors"
	"mangav4/system/download/types"
)

type Download struct {
	Dl types.Downloads `json:"downloads"`
}

// option for download params
type Option struct {
	URL        string `json:"url"`
	ServerName string `json:"server_name"`
	Page       *int   `json:"page"`
	Limit      *int   `json:"limit"`
}

// Get new Download instance
func NewDownload() *Download {
	return new(Download)
}

func (f *Download) DownloadChapter(o Option) (*types.Chapter, error) {
	class := *types.NewDownload(o.ServerName)
	if class == nil {
		return nil, errors.New("error: Server Name " + o.ServerName + " Not Found")
	}
	opt := types.OptionDownload{
		Id: o.URL,
	}

	return class.GetChapter(opt)
}

func (f *Download) DownloadPage(o Option) (*types.Page, error) {
	class := *types.NewDownload(o.ServerName)
	if class == nil {
		return nil, errors.New("error: Server Name " + o.ServerName + " Not Found")
	}
	opt := types.OptionDownload{
		Id: o.URL,
	}

	return class.GetPage(opt)
}
