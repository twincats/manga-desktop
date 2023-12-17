package internet

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cavaliergopher/grab/v3"
)

// NewFileDownloader return pointer FileDownloader
func NewFileDownloader() *FileDownloader {
	return &FileDownloader{}
}

// FileDownloader struct with http header
type FileDownloader struct {
	httpHeader *http.Header
}

// StatusBatch is ouput from StatusBatch reading chan from DownloadBatch
type StatusBatch struct {
	Err      error  `json:"err"`
	Index    int    `json:"index"`
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

// setupClient() is internal methods for creating grab client
func (f *FileDownloader) setupClient() *grab.Client {
	client := grab.NewClient()
	client.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:109.0) Gecko/20100101 Firefox/112.0"
	return client
}

// SetupDirectory() is internal methods for creating emty directory
func (f *FileDownloader) SetupDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// SetReferal is easy methods for Setting referal header return instance FileDownloader
func (f *FileDownloader) SetReferal(r string) *FileDownloader {
	if f.httpHeader == nil {
		f.httpHeader = &http.Header{}
	}
	f.httpHeader.Set("Referer", r)
	return f
}

// SetHeader is methods for Setting header return instance FileDownloader
func (f *FileDownloader) SetHeader(key, value string) *FileDownloader {
	if f.httpHeader == nil {
		f.httpHeader = &http.Header{}
	}
	f.httpHeader.Set(key, value)
	return f
}

// BatchDownload is methods for download in Batch with auto rename filename by index
func (f *FileDownloader) BatchDownload(worker int, dst string, remote_urls []string) <-chan *grab.Response {
	f.SetupDirectory(dst)

	clinet := f.setupClient()

	reqs := make([]*grab.Request, len(remote_urls))

	for i := 0; i < len(remote_urls); i++ {
		req, err := grab.NewRequest(dst, remote_urls[i])
		if err != nil {
			fmt.Println(err)
		}

		if f.httpHeader != nil {
			req.HTTPRequest.Header = *f.httpHeader
		}

		// auto rename filename by index
		ext := filepath.Ext(remote_urls[i])
		filename := fmt.Sprintf("%v%v", i+1, ext)
		req.Filename = filepath.Join(dst, filename)

		req.Tag = i

		reqs[i] = req
	}

	return clinet.DoBatch(worker, reqs...)

}

// StatusBatch read channael from BatchDownload return failed url and statusBatch
func (f *FileDownloader) StatusBatch(ch <-chan *grab.Response, emit func(f StatusBatch)) ([]string, []StatusBatch) {
	var statusBatch []StatusBatch
	var failed []string
	for resp := range ch {
		var status StatusBatch
		if resp.Err() != nil {
			status.Err = resp.Err()
			failed = append(failed, resp.Request.HTTPRequest.URL.String())
		}
		status.Filename = resp.Filename
		status.URL = resp.Request.HTTPRequest.URL.String()
		status.Index = resp.Request.Tag.(int)
		statusBatch = append(statusBatch, status)
		emit(status)
	}
	return failed, statusBatch
}

// Download single file without concurency
func (f *FileDownloader) Download(dst, url string) *grab.Response {
	client := f.setupClient()
	req, _ := grab.NewRequest(dst, url)
	return client.Do(req)
}

// FilterStatusErr is filter []StatusBatch to show Error data only
func (f *FileDownloader) FilterStatusErr(e []StatusBatch) []StatusBatch {
	var failed_only []StatusBatch
	for _, v := range e {
		if v.Err != nil {
			failed_only = append(failed_only, v)
		}
	}
	return failed_only
}

func (f *FileDownloader) RemoveFolder(p string) error {
	return os.RemoveAll(p)
}
