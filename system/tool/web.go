package tool

import (
	"io"
	"net/http"
)

type Web struct {
	Url  string `json:"url"`
	Body string `json:"body"`
	Err  error  `json:"error"`
}

func NewWeb() *Web {
	return &Web{}
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
