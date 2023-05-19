package tool

import (
	"io"
	"mangav4/system/app"
	"net/http"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Web struct {
	Url  string `json:"url"`
	Body string `json:"body"`
	Err  error  `json:"error"`
}

func NewWeb() *Web {
	return &Web{}
}

func (w *Web) Fetch(u string) (string, error) {
	res, err := app.C.R().Get(u)
	if err != nil {
		runtime.LogErrorf(*app.WailsContext, "Error Fetch : %v", err)
		return "", err
	}
	return res.String(), nil
}

func (f *Web) FetchPost(u string, data interface{}) (string, error) {
	res, err := app.C.R().SetBody(data).Post(u)
	if err != nil {
		runtime.LogErrorf(*app.WailsContext, "Error Fetch : %v", err)
		return "", err
	}
	return res.String(), nil
}

func (f *Web) WebBrowser(url string) (*Web, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	f.Url = url
	f.Body = string(body)
	f.Err = err
	return f, nil
}
